# 🌐 Godplan System Backend

> **Scalable Multi-Tenant Business Engine** - The core high-performance backend powering the Godplan ecosystem.

[![Go Version](https://img.shields.io/badge/Go-1.24-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Status](https://img.shields.io/badge/Status-Production-brightgreen?style=flat)](https://github.com/nepskuy/be-godplan)
[![Framework](https://img.shields.io/badge/Framework-Gin-0081CB?style=flat)](https://gin-gonic.com/)

Godplan System Backend is a high-concurrency microservice designed for multi-tenant business operations. It provides a unified API for resource management, complex business logic processing, and multi-tenant data sharding.

---

## 🔥 Key Technical Capabilities

### 🏢 Database-Level Isolation
- **Advanced Multi-Tenancy**: Implements strict data separation architecture at the schema and query level to ensure 100% tenant privacy.
- **Dynamic Sharding Support**: Designed for horizontal scalability across multiple database clusters.

### 💼 Core Business Modules
- **Employee Lifecycle Management**: End-to-end management of personnel, contracts, and hierarchies.
- **Attendance & Operations**: Real-time operational tracking with geofencing and administrative controls.
- **Payroll & Finance**: Automated calculation engines for complex salary structures and tax compliance.
- **CRM Pipeline**: Integrated customer relationship management with deal-to-project conversions.

### 🛡️ Hardened Security
- **JWT Authentication**: Secure, stateless authentication for all API requests.
- **Bcrypt Encryption**: Industry-standard password hashing.
- **Global Audit Trail**: Systematic logging of all administrative actions (includes IP, User Agent, and timestamp).

### 🤖 Intelligent Strategic Advisor
- **Scenario Simulation**: Advanced AI logic for business impact assessment and growth projection (Powered by Step 3.5).

---

## 🛠️ Technology Stack

| Layer | Technology |
| :--- | :--- |
| **Language** | Go (Golang) 1.24 |
| **Framework** | [Gin Gonic](https://gin-gonic.com/) |
| **ORM** | [GORM](https://gorm.io/) |
| **Database** | PostgreSQL |
| **Containerization** | Docker & Docker Compose |

---

## 🚀 Getting Started

### 1. Prerequisites
- Go 1.24+
- Docker & Docker Compose
- PostgreSQL 14+

### 2. Environment Configuration
Copy the production environment example:
```bash
cp .env.prod .env.local
```
Update your credentials in `.env.local`:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=godplan
JWT_SECRET=your_secret_key
```

### 3. Docker Deployment (Recommended)
Run the entire stack using Docker:
```bash
# Start all services
docker-compose up -d

# Check service logs
docker-compose logs -f app
```

### 4. Manual Running
```bash
# Download dependencies
go mod download

# Run service
go run cmd/main.go
```

---

## 📂 Directory Structure

- `cmd/`: Service entry points (CLI, HTTP server).
- `pkg/handlers/`: RESTful API endpoint handlers.
- `pkg/services/`: Modularized business logic layer.
- `pkg/repositories/`: Data access objects (DAO) and GORM models.
- `pkg/middleware/`: Global request processing (Auth, CORS, Security).
- `scripts/`: Maintenance and migration scripts.

---

## 🛡️ Security Best Practices
Always ensure that the `DATABASE_URL` uses `sslmode=require` in production and that sensitive environment variables are managed through a secure vault or encrypted CI/CD secrets.

---

## 📄 License
Copyright © 2026. All rights reserved. Built with ❤️ by **Godplan Engineering Team**.
