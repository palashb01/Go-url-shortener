# 🚀 URL Shortener App

A full-featured URL Shortener service built with **Golang**, **gRPC**, **gRPC-Gateway**, **Redis**, **Docker**, and a clean **HTML frontend UI**. You can shorten long URLs, resolve short codes back to original URLs, and interact via REST, gRPC, or Swagger UI.

Hosted live on [Render](https://render.com) with Upstash Redis 🚀

---

## 🧱 Tech Stack

| Layer        | Tech                          |
|--------------|-------------------------------|
| Language     | Go 1.23+                      |
| API          | gRPC + gRPC-Gateway (REST)    |
| Storage      | Upstash Redis (TLS-secured)   |
| Frontend     | HTML + CSS + JS               |
| API Docs     | Swagger/OpenAPI + Swagger UI  |
| Packaging    | Docker + Multi-stage builds   |
| Deployment   | Render (Docker image)         |

---

## 📁 Project Structure

```bash
.
├── cmd/                    # Main server entrypoint
│   └── server/main.go
├── internal/
│   └── service/            # gRPC service logic
├── gen/                   # Generated protobuf & gateway files
│   ├── go/
│   └── openapiv2/
├── proto/                 # Protobuf definitions
│   └── shortener/v1/
│       └── shortener.proto
├── frontend/              # Static frontend UI served via Go
│   └── index.html
├── swagger-ui/            # Swagger UI assets (optional)
├── docker-compose.yml     # Local dev with Docker
├── Dockerfile             # Multi-stage container image
├── buf.yaml               # Buf module config
├── buf.gen.yaml           # Buf codegen plugins
└── README.md              # This file
```

---

## 🧪 Features

- 🔗 **Shorten URLs** — POST a long URL and get a short code
- 🧭 **Resolve Short Codes** — Resolve any code to the original URL
- 💬 **REST + gRPC APIs** (via gRPC-Gateway)
- 🧼 Clean static UI with live link shortening
- 📜 **Swagger UI** auto-generated from proto
- 🐳 Production-ready Docker build & deploy

---

## 💻 Local Development

### Prerequisites
- Go 1.23+
- Docker + Docker Compose
- [Buf CLI](https://buf.build/docs/installation)

### Generate Protobuf & Gateway Files
```bash
buf generate
```

### Run Locally (All Services)
```bash
docker-compose up --build
```

> Exposed Ports:
> - Frontend UI → http://localhost:3000
> - REST API    → http://localhost:8080
> - gRPC        → localhost:9090
> - Swagger UI  → http://localhost:8081

---

## 🔗 API Endpoints

### `POST /v1/shorten`
**Request:**
```json
{
  "originalUrl": "https://palash.com"
}
```
**Response:**
```json
{
  "shortUrl": "https://your-app.onrender.com/v1/resolve/abc123"
}
```

### `GET /v1/resolve/{shortCode}`
**Response:**
```json
{
  "originalUrl": "https://palash.com"
}
```

---

## 🖥️ Frontend UI

- Auto-served from `/`
- Type a long URL → shorten → copy
- Paste a short code → resolve → get original URL
- Link to Swagger docs auto-generated via `window.location.origin`

---

## 📜 Swagger Docs

- Live at `/docs` on the deployed app
- Interactive docs from `shortener.swagger.json`
- Generated via `buf + openapiv2`

---

## 🧑‍💻 Deployment

Deployed to [Render](https://render.com) with:
- Dockerfile (multi-stage)
- Upstash Redis (external cloud DB)
- Auto-detected port binding using `$PORT`
- Environment Variables:
  - `REDIS_ADDR`
  - `REDIS_PASSWORD`
  - `API_HOST`

You can fork and deploy this to your own Render project!

---

## 🤝 Contributing

Want to contribute, improve the UI, or extend features? PRs welcome! 🫶

---

> This is a fun and self-learning project only.

