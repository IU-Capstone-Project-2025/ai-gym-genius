# GymGenius 🏋️‍♂️

*Offline-first workout tracker with AI-powered feedback and a full web admin console.*

---

## Table of Contents
1. [Monorepo Layout](#monorepo-layout)
2. [Key Features](#key-features)
3. [Tech Stack](#tech-stack)
4. [Getting Started](#getting-started)
5. [License](#license)

---

## Monorepo Layout

```text
.
├─ mobile/               # Flutter app (Android & iOS)
│  └─ lib/
├─ api/                  # Go services (gin + gRPC)
│  ├─ cmd/
│  └─ internal/
├─ admin/                # Vue-TS web console
│  └─ src/
├─ scripts/              # helper bash/PowerShell helpers
└─ docs/                 # ADRs, diagrams, onboarding
```

---

## Key Features

| Mobile                                                                                                                                               | Backend                                                                                                                                         | Admin                                                                                                                                           |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| -  Start/stop workouts, add sets/reps/weight-  Heat-map statistics and graphs-  ❄️ **Offline-first** (local DB → async sync)-  AI "Get Feedback" button | -  gRPC + REST API-  Delta-sync algorithm-  PostgreSQL + SQLC-  Stripe web-hooks for subscriptions-  OpenAI-powered workout analysis | -  Metrics dashboard (PromQL via Grafana)-  Role-based access control |

---

## Tech Stack

| Layer  | Tech                                          |
| ------ | --------------------------------------------- |
| Mobile | Flutter 3.22 / Dart 3 · Bloc · Sqlite     |
| API    | Go 1.22 · gin · gRPC-Gateway · PostgreSQL >14 |
| Admin  | React 18 (Vite) · TypeScript 5 · shadcn/ui · TanStack Query |
| AI     | OpenAI API · Golang concurrency workers |
| DevOps | Docker-Compose · GitHub Actions CI · Terraform (Linodes) |

---

## Getting Started

### Mobile App

#### Prerequisites
- [Docker](https://docs.docker.com/get-docker/) installed on your machine

#### Build and Run
```bash
# Build the Flutter web app
docker build -t gym-genius-mobile -f mobile/Dockerfile .

# Run the container
docker run -p 8080:80 gym-genius-mobile

# Access the app at http://localhost:8080
```

#### Local Development (Optional)
```bash
cd mobile/
flutter pub get
flutter run -d chrome --web-port 8080
```

## License

Distributed under the **MIT license** – see [`LICENSE`](LICENSE) for details.

```
Happy gains 💪!
```
