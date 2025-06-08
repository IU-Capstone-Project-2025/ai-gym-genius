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
````

---

## Key Features

| Mobile                                                                                                                                               | Backend                                                                                                                                         | Admin                                                                                                                                           |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| • Start/stop workouts, add sets/reps/weight<br>• Heat-map statistics and graphs<br>• ❄️ **Offline-first** (local DB → async sync)<br>• AI “Get Feedback” button | • gRPC + REST API<br>• Delta-sync algorithm<br>• PostgreSQL + SQLC<br>• Stripe web-hooks for subscriptions<br>• OpenAI-powered workout analysis | • Metrics dashboard (PromQL via Grafana)<br>• Role-based access control |

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

Will be added soon. 

## License

Distributed under the **MIT license** – see [`LICENSE`](LICENSE) for details.

```
Happy gains 💪!
```
