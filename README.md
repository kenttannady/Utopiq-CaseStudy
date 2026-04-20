# 🚀 Utopiq Case Study

Project ini adalah demo sistem monitoring berbasis container menggunakan:
- Prometheus (metrics)
- Grafana (visualisasi)
- Loki (logging)
- AlertManager (alerting)
- Docker Compose (orchestration)

---

## 📁 Project Structure


monitoring-demo/
├── frontend/
│ ├── Dockerfile
│ ├── package.json
│ ├── server.js
│ └── nginx.conf
├── backend/
│ ├── Dockerfile
│ ├── main.go
│ └── go.mod
├── docker-compose.yml
├── prometheus/
│ └── prometheus.yml
├── loki/
│ └── loki-config.yml
├── alertmanager/
│ └── alertmanager.yml
├── grafana/
│ └── provisioning/
│ ├── datasources/
│ └── dashboards/
├── .github/workflows/ci-cd.yml
└── .env (secret management)


---

## ⚙️ Setup & Running (Phase 1)

### 1. Setup Environment Variables

```bash


1. Jalankan Semua Services
docker compose up -d
2. Cek Logs
docker compose logs -f
🌐 Service Access
Service	URL	Notes
Frontend	http://localhost:3000
	UI aplikasi
Grafana	http://localhost:3001
	user: admin
Prometheus	http://localhost:9090
	metrics
AlertManager	http://localhost:9093
	alert monitoring

🔑 Password Grafana diambil dari .env

🚨 Testing Alert (Telegram)
Simulasi High Error Rate
for i in {1..100}; do curl http://localhost:3000/api/proxy; done
Simulasi Service Down
docker compose stop backend

⏱️ Alert akan ter-trigger dalam 1–2 menit.

📦 Logging dengan Loki di Grafana
Login ke Grafana → http://localhost:3001
Masuk ke Data Sources
Tambahkan datasource:
Type: Loki
URL: http://loki:3100
Masuk ke Explore
Gunakan query:
{service="backend"}
📊 Monitoring Metrics

Dashboard Grafana menampilkan:

📈 Request rate (RPS) per service
❌ Error rate (HTTP 5xx)
⏱️ Latency (P50, P95, P99)
💚 Service health (up/down)
📜 Log aggregation dari semua container (via Loki)
🚨 Alert history dari AlertManager
🔔 Alert Rules (Telegram)
Condition	Severity
Error rate > 10% (5 menit)	Critical
P95 latency > 1 detik	Warning
Service down	Critical
🔐 Secret Management Strategy
Local Development (Docker Desktop)
Gunakan .env file (JANGAN commit ke Git!)
Gunakan Docker secrets untuk production
CI/CD (GitHub Secrets)

Tambahkan secrets berikut di GitHub:

DB_PASSWORD
GRAFANA_PASSWORD
TELEGRAM_BOT_TOKEN
TELEGRAM_CHAT_ID
GH_PAT (untuk approval gate)
📝 Notes
Pastikan Docker & Docker Compose sudah terinstall
Jangan pernah commit file .env
Gunakan .env.example sebagai template
📌 TODO (Optional Improvement)
 Tambah authentication di frontend
 Setup HTTPS (reverse proxy / Traefik / Nginx)
 Integrasi dengan cloud monitoring
 Auto scaling services
