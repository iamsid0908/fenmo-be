# Fenmo — Personal Expense Tracker (Backend)

A RESTful backend API for tracking personal and shared expenses, built with **Go**, **Echo**, and **PostgreSQL**.

---

## Table of Contents
- [Tech Stack & Rationale](#tech-stack--rationale)
- [System Design](#system-design)
- [Key Design Decisions](#key-design-decisions)
- [API Overview](#api-overview)
- [Project Structure](#project-structure)
- [Running Locally](#running-locally)
- [Trade-offs & Known Limitations](#trade-offs--known-limitations)

---

## Tech Stack & Rationale

### Go (Golang)
- **Required by the JD** — primary backend language specified in the job description.
- **Performance** — Go is statically typed and compiled, giving near-C-level throughput with minimal resource usage. A single small instance can handle thousands of concurrent requests thanks to goroutines.
- **Concurrency model** — goroutines + channels make it easy to handle concurrent HTTP requests, background jobs (email sending, OTP expiry), and future event-driven features without heavyweight threads.
- **Simplicity** — small standard library surface, fast compile times, and straightforward error handling make the codebase easy to review and maintain.
- **Ecosystem** — mature HTTP frameworks (Echo), ORMs (GORM), and JWT libraries exist and are production-grade.

### PostgreSQL
- **Required by the JD** — PostgreSQL is explicitly mentioned in the role.
- **ACID compliance** — every expense write is atomic. This is critical when we add shared-expense splitting: debiting one user and crediting another must succeed or fail together. PostgreSQL transactions will handle this cleanly without any application-level hacks.
- **Scalability** — PostgreSQL scales vertically very well and supports read replicas for horizontal read scaling. Features like table partitioning (partition `expenses` by `user_id` or `date`) make it suitable for millions of records.
- **Rich query support** — window functions, CTEs, and JSON columns let us run analytics (monthly summaries, category breakdowns) in the database layer rather than in application code.
- **pgvector extension** — already integrated (`pgvector-go`) for potential AI-powered features like semantic expense categorization or smart search.
- **Future transactions** — PostgreSQL's `BEGIN / COMMIT / ROLLBACK` support maps directly to GORM's `.Transaction()` helper. When we implement shared-expense settlement, the entire flow (create expense → split → update balances) can run in a single DB transaction.

### Echo (HTTP Framework)
- Minimal and fast; middleware (JWT, CORS, logging) plugs in cleanly.
- Built-in request binding and validation hooks.

### GORM (ORM)
- Reduces boilerplate for CRUD while still allowing raw SQL for complex queries.
- Auto-migration support speeds up schema iteration during development.

---

## System Design

```
┌─────────────────────────────────────────────────────────┐
│                        Client                           │
│              (Mobile App / Web Frontend)                │
└───────────────────────┬─────────────────────────────────┘
                        │ HTTPS
                        ▼
┌─────────────────────────────────────────────────────────┐
│                   Echo HTTP Server                      │
│                                                         │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐              │
│  │   Auth   │  │ Expense  │  │ Category │  ...         │
│  │ Handler  │  │ Handler  │  │ Handler  │              │
│  └────┬─────┘  └────┬─────┘  └────┬─────┘              │
│       │              │              │                   │
│  ┌────▼─────────────▼──────────────▼─────┐             │
│  │           Service Layer                │             │
│  │  (business logic, validation, mapping) │             │
│  └────────────────────┬───────────────────┘             │
│                        │                               │
│  ┌─────────────────────▼───────────────────┐           │
│  │           Domain Layer (GORM)            │           │
│  │      (DB queries, model structs)         │           │
│  └─────────────────────┬───────────────────┘           │
└────────────────────────┼────────────────────────────────┘
                         │
            ┌────────────▼────────────┐
            │      PostgreSQL DB      │
            │                        │
            │  users                 │
            │  expenses              │
            │  categories            │
            │  user_lists            │
            │  auth_tokens           │
            └────────────────────────┘
                         │
            ┌────────────▼────────────┐
            │   External Services     │
            │  • Google OAuth         │
            │  • GitHub OAuth         │
            │  • SMTP (Email/OTP)     │
            │  • Azure OpenAI (AI)    │
            └─────────────────────────┘
```

### Layer Responsibilities

| Layer | Responsibility |
|---|---|
| **Handler** | Parse request, input validation, bind to models, return HTTP response |
| **Service** | Business logic — expense rules, OTP generation, token management |
| **Domain** | Database access via GORM — queries, preloads, pagination |
| **Models** | Shared struct definitions (DB models + request/response DTOs) |
| **Middleware** | JWT authentication, request logging |
| **Config** | DB connection pool, environment config, webhook setup |

---

## Key Design Decisions

### 1. Layered Architecture (Handler → Service → Domain)
Each layer has a single responsibility and communicates through interfaces. This makes unit testing and future changes (e.g., swapping GORM for sqlc) straightforward without touching business logic.

### 2. JWT Authentication
Stateless JWT tokens mean the server does not need a session store. Every protected route validates the token via the `JWTVerify()` middleware — horizontally scalable with zero shared state.

### 3. OTP-based Email Verification
Users verify their email before login is permitted (`user_active` flag). This prevents spam accounts and is a prerequisite for any future payment or financial feature.

### 4. OAuth Support (Google + GitHub)
Reduces friction for users who prefer social login, and avoids us storing passwords for those accounts.

### 5. User Lists for Shared Expenses
Expenses belong to a `UserList` (a group). This is the foundation for splitting bills — the DB schema is already normalized to support multi-user expense groups without breaking changes.

### 6. Pagination on Expense Listing
All list endpoints return `page`, `page_size`, `total_pages`, and `total_records`. This keeps payloads small and API responses fast regardless of how many records a user accumulates.

### 7. Soft-delete Ready
GORM's `DeletedAt` pattern can be added to any model without a migration rewrite, preserving audit history.

---

## API Overview

| Method | Endpoint | Auth | Description |
|---|---|---|---|
| GET | `/health` | ❌ | Health check |
| POST | `/v1/auth/register` | ❌ | Register new user |
| POST | `/v1/auth/resend-otp` | ❌ | Resend OTP |
| POST | `/v1/auth/verify-otp` | ❌ | Verify email OTP |
| POST | `/v1/auth/login` | ❌ | Login |
| GET | `/v1/auth/validate` | ✅ | Validate JWT session |
| GET | `/v1/auth/logout` | ✅ | Logout |
| GET | `/v1/auth/google` | ❌ | Google OAuth URL |
| GET | `/v1/auth/github` | ❌ | GitHub OAuth URL |
| GET | `/v1/user/get-user` | ✅ | Get user profile |
| POST | `/v1/user/update-profile` | ✅ | Update profile |
| GET | `/v1/user-list/get` | ✅ | Get user lists |
| POST | `/v1/user-list/create` | ✅ | Create user list |
| GET | `/v1/user-list/get_expenses` | ✅ | Expenses in a list |
| POST | `/v1/category/create` | ✅ | Create category |
| GET | `/v1/category/list` | ✅ | List categories |
| POST | `/v1/expense` | ✅ | Create expense |
| GET | `/v1/expense/list` | ✅ | List expenses (paginated, sorted newest first) |

---

## Project Structure

```
fenmo/
├── main.go                  # Entry point
├── config/                  # DB connection, env config, webhooks
├── domain/                  # GORM queries (data access layer)
├── service/                 # Business logic layer
├── handler/                 # HTTP handlers + request validation
│   └── validation/          # Input validators per feature
├── middleware/              # JWT auth middleware
├── models/                  # Shared structs (DB + DTOs)
├── route/                   # Route registration
├── utils/                   # Constants, error vars, helpers
├── template/                # Email HTML templates
├── Dockerfile
└── ecosystem.config.json    # PM2 config (process manager)
```

---

## Running Locally

### Prerequisites
- Go 1.21+
- PostgreSQL 15+

### Setup

```bash
# Clone the repository
git clone <repo-url>
cd fenmo

# Copy environment variables
cp .env.example .env
# Fill in DB_URL, JWT_SECRET, SMTP credentials, OAuth keys

# Run database migrations (auto-migration via GORM on startup)
go run main.go
```

### Environment Variables

| Variable | Description |
|---|---|
| `DB_URL` | PostgreSQL connection string |
| `JWT_SECRET` | Secret key for signing JWT tokens |
| `SMTP_HOST` | Email server host |
| `SMTP_PORT` | Email server port |
| `GOOGLE_CLIENT_ID` | Google OAuth client ID |
| `GOOGLE_CLIENT_SECRET` | Google OAuth client secret |
| `GITHUB_CLIENT_ID` | GitHub OAuth client ID |
| `GITHUB_CLIENT_SECRET` | GitHub OAuth client secret |

---

## Trade-offs & Known Limitations

### Deployment
Due to time constraints, the service is **not yet deployed** to a cloud environment. The application is fully functional locally. A `Dockerfile` is included for containerization and the `ecosystem.config.json` supports deployment via PM2 on a VPS or cloud VM.

**Planned deployment approach:**
- Containerize with Docker
- Host on a cloud provider (AWS EC2 / Railway / Render)
- Use a managed PostgreSQL instance (AWS RDS / Supabase)
- Nginx as a reverse proxy with SSL termination

### What I Would Add With More Time

| Feature | Reason |
|---|---|
| **DB Transactions** | Wrap shared-expense creation + balance updates in a single `BEGIN/COMMIT` block using GORM's `.Transaction()` to guarantee consistency |
| **Unit & Integration Tests** | Service layer logic is testable via interface mocking — tests are the next priority |
| **Rate Limiting** | Prevent OTP abuse and brute-force login attempts |
| **Role-based Access Control** | Fine-grained permissions per user list |
| **Expense Settlement** | Calculate net balances within a user list and record settlements |
| **WebSocket / SSE** | Real-time notifications when a group member adds an expense |
| **CI/CD Pipeline** | GitHub Actions → Docker build → deploy on merge to main |
| **Observability** | Structured logging (already partially in place), metrics, and distributed tracing |

### Current Assumptions
- Single currency per expense (currency field exists but multi-currency conversion is not yet implemented).
- No soft-delete — deleted records are permanently removed.
- Email delivery is synchronous — in production this should be moved to a background worker queue (e.g., Redis + Go worker).

---

> Built as part of a job application assignment. The focus was on clean architecture, correct use of Go idioms, and a PostgreSQL schema that is ready to scale — rather than a fully deployed production system within the time limit.
