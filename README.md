# Information notice / Bilgilendirme notu

- [Türkçe (TR)](#türkçe-tr)
- [English (ENG)](#english-eng)

---

## Türkçe (TR)

Bu depodaki yazılımlar, yetkisiz sistemlere erişim, hizmet kesintisi veya kötüye kullanıma yönlendirilebilecek bileşenler içerebilir. Bu belge **kurulum, yapılandırma veya çalıştırma talimatı vermez**; yalnızca bağlam ve savunma odaklı genel bilgi sunar.

### Yasal ve etik uyarı

- Bilgisayar sistemlerine izinsiz müdahale birçok ülkede **suçtur** (yetkisiz erişim, veri bozma, hizmet dışı bırakma vb.).
- Yalnızca **yazılı izin** (sözleşme, penetration test kapsamı, kurumsal politika) ve **kendi kontrolünüzdeki ortamlarda** güvenlik testi yapın.
- Bu depoyu veya benzer araçları üretim ağlarında, üçüncü taraflara ait sistemlerde veya izinsiz bağlantılarla kullanmayın.

### Botnet ve C&C kavramları (genel)

**Botnet:** Çok sayıda ele geçirilmiş veya kullanıcı rızası dışında çalıştırılmış istemcinin tek merkez veya kontrol katmanıyla koordine edildiği yapılara verilen genel addır.

**Komuta ve kontrol (C&C / C2):** Saldırganların botlara güncelleme, yapılandırma veya ek yük iletmesi için kullandığı iletişim kanalları. Tespit için ağ izleme, DNS/HTTPS anomalileri ve davranış analizi sık kullanılır.

Savunma tarafında tipik yaklaşımlar: uç birim koruması, ağ segmentasyonu, sıfıra yakın güven, günlük izleme, tehdit avcılığı ve olay müdahale süreçleri.

### Güvenlik araştırması ve öğrenme

Aşağıdaki konular yasal ve kontrollü laboratuvar ortamlarında daha güvenli şekilde çalışılır:

- Ağ güvenliği ve olay müdahale (DFIR)
- Güvenlik açığı yönetimi ve yamalama
- Etik hackerlık ve sertifikalı eğitim programları

Resmî dokümantasyon, üniversite/şirket lab’leri ve CTF yarışmaları, bu alanda tercih edilen yollar arasındadır.

### Bu depo hakkında

Bu README, **bilgilendirme amaçlıdır**. Depo içeriğinin incelenmesi veya kullanımı size ve bulunduğunuz yargı bölgesine göre hukuki sonuçlar doğurabilir. Şüphe durumunda bir hukuk veya güvenlik danışmanına başvurun.

---

## English (ENG)

Software in this repository may include components that could be abused for unauthorized access, service disruption, or other harmful purposes. This document **does not provide installation, configuration, or runtime instructions**; it only offers general context and a defensive perspective.

### Legal and ethical notice

- Unauthorized interference with computer systems is a **crime** in many jurisdictions (unauthorized access, data damage, denial of service, etc.).
- Perform security testing only with **explicit written authorization** (contract, penetration test scope, organizational policy) and in **environments under your control**.
- Do not use this repository or similar tools against production networks, third-party systems, or without authorization.

### Botnets and C&C (general concepts)

**Botnet:** A general term for architectures where many compromised or unsolicited clients are coordinated through a central or hierarchical control plane.

**Command and control (C&C / C2):** Communication channels attackers use to deliver updates, configuration, or additional payloads to bots. Detection commonly relies on network monitoring, DNS/HTTPS anomalies, and behavioral analysis.

Typical defensive practices include endpoint protection, network segmentation, zero trust principles, logging, threat hunting, and incident response processes.

### Security research and learning

The topics below are safer to study in legal, controlled lab environments:

- Network security and digital forensics / incident response (DFIR)
- Vulnerability management and patching
- Ethical hacking and accredited training programs

Official documentation, university or corporate labs, and CTF competitions are common, legitimate learning paths.

### About this repository

This README is **for informational purposes only**. Reviewing or using repository content may have legal consequences depending on your situation and jurisdiction. When in doubt, consult legal counsel or a qualified security advisor.
