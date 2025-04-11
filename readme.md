# 🚀 URL Shortener App

A full-featured URL Shortener service built with **Golang**, **gRPC**, **gRPC-Gateway**, **Redis**, **Docker**, and a minimal **HTML frontend UI**. You can shorten long URLs, resolve short codes back to original URLs, and interact via REST, gRPC, or Swagger UI.

---

## 🧱 Tech Stack

| Layer        | Tech                         |
|--------------|-------------------------------|
| Language     | Go 1.23+                      |
| API          | gRPC + gRPC-Gateway (REST)    |
| Storage      | Redis                         |
| Frontend     | HTML + CSS + JS (no framework)|
| API Docs     | Swagger/OpenAPI + Swagger UI  |
| Packaging    | Docker & docker-compose       |
| Deployment   | Cloud-ready (Railway/Fly.io)  |

---

## 📁 Project Structure

```bash
.
├── cmd/                    # Main server entrypoint
│   └── server/main.go
├── internal/
│   ├── service/            # gRPC service logic
├── gen/                   # Generated protobuf & gateway files
│   ├── go/
│   └── openapiv2/
├── proto/                 # Protobuf definitions
│   └── shortener/v1/
│       └── shortener.proto
├── frontend/              # Static UI (served via NGINX)
│   └── index.html
├── docker-compose.yml     # Full Docker setup
├── Dockerfile             # Docker image for Go server
├── buf.yaml               # Buf module config
├── buf.gen.yaml           # Buf codegen plugins
└── README.md              # This file
```

---

## 🧪 Features

- 🔗 **Shorten URLs** — POST a long URL and get a short code
- 🧭 **Resolve Short Codes** — Resolve any code to the original URL
- 💬 **REST + gRPC APIs**
- 🧼 Clean, pretty HTML UI to interact visually
- 🧪 Swagger/OpenAPI documentation with live testing
- 🐳 Docker + Compose for unified local/dev/prod environments

---

## 💻 Local Development

### 🧰 Prerequisites
- Go 1.23+
- Docker + Docker Compose
- [Buf CLI](https://buf.build/docs/installation) (for generating proto files)

### 🔧 Generate Protobuf Code

```bash
buf generate
```

### 🐳 Run the Full App with Docker Compose

```bash
docker-compose up --build
```

> Ports:
> - Frontend UI → http://localhost:3000
> - REST API → http://localhost:8080
> - gRPC Server → localhost:9090
> - Swagger UI → http://localhost:8081

---

## 🔗 REST API Endpoints

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
  "shortUrl": "http://localhost:8080/v1/resolve/abc123"
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

Available at: `http://localhost:3000`

- Input a long URL → Click “Shorten” → Get a short URL
- Enter a short code → Click “Resolve” → Get original URL
- Includes link to Swagger docs

---

## 📜 Swagger Docs

Available at: `http://localhost:8081`

Swagger UI loads `shortener.swagger.json` generated from your protobuf definitions. You can:
- Test endpoints live
- See expected request/response formats
- Try `POST /v1/shorten`, `GET /v1/resolve/{shortCode}`

---

## 🤝 Contributing

Want to contribute, improve the UI, or extend features? PRs welcome! 🫶

---

## This is a fun and self-learning project only.