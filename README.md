# Utopiq-CaseStudy

Structure Project

monitoring-demo/
├── frontend/
│   ├── Dockerfile
│   ├── package.json
│   ├── server.js
│   └── nginx.conf
├── backend/
│   ├── Dockerfile
│   ├── main.go
│   └── go.mod
├── docker-compose.yml
├── prometheus/
│   └── prometheus.yml
├── loki/
│   └── loki-config.yml
├── alertmanager/
│   └── alertmanager.yml
├── grafana/
│   └── provisioning/
│       ├── datasources/
│       └── dashboards/
├── .github/workflows/ci-cd.yml
└── .env (secret management)



Phase 1
# Buat file .env dari template
cp .env.example .env
# Edit .env dengan credentials asli

# Jalankan semua services
docker compose up -d

# Cek logs
docker compose logs -f

# Akses:
# - Frontend: http://localhost:3000
# - Grafana: http://localhost:3001 (user: admin, pass: dari .env)
# - Prometheus: http://localhost:9090
# - AlertManager: http://localhost:9093

Testing Alert ke Telegram:
===========================
# Simulate high error rate
for i in {1..100}; do curl http://localhost:3000/api/proxy; done

# Atau matikan backend
docker compose stop backend
# Alert akan trigger dalam 1-2 menit


Integrasi Logging dengan Loki di Grafana:
Login Grafana (localhost:3001)

Add datasource → Pilih Loki

URL: http://loki:3100

Explore → Pilih {service="backend"}


📊 Monitoring Metrics yang Ditampilkan
Dashboard Grafana akan show:
- Request rate (RPS) per service
- Error rate (5xx status codes)
- Latency P50, P95, P99 (histogram quantiles)
- Service health (up/down)
- Log aggregation dari semua container via Loki
- Alert history dari AlertManager

Alert Rules (via Telegram):
- High error rate > 10% dalam 5 menit → Critical
- P95 latency > 1 second → Warning
- Service down → Critical (wake someone up)

🔐 Secret Management Strategy
Untuk Docker Desktop local:
- .env file (jangan commit ke git!)
- Docker secrets (production)

GitHub Secrets untuk CI/CD:
- DB_PASSWORD
- GRAFANA_PASSWORD
- TELEGRAM_BOT_TOKEN
- TELEGRAM_CHAT_ID
- GH_PAT (untuk approval gate)