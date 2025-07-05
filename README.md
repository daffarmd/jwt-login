# task-manager
# 📝 Task Manager API

A simple, clean, and scalable RESTful API for managing tasks — built with **Go**, **Fiber**, and **PostgreSQL** using a layered architecture.

---

## 🚀 Features

- ✅ User registration and login with hashed password
- ✅ JWT-based authentication
- ✅ Create, read, update, and delete (CRUD) tasks
- ✅ Clean architecture with separation of concerns
- ✅ Fully tested with unit tests for repository and service layers

---

## 🧰 Tech Stack

- **Golang** – main backend language
- **Fiber** – high performance web framework
- **PostgreSQL** – relational database
- **bcrypt** – for password hashing
- **JWT** – for secure user authentication
- **sql.DB** – database driver
- **Testify** – for unit testing

---

## 📁 Project Structure
task-manager/
│
├── cmd/ # App entry point
├── config/ # Config loader (.env)
├── internal/
│ ├── model/ # Domain models
│ ├── repository/ # Database access layer
│ ├── service/ # Business logic
│ └── handler/ # HTTP handlers (controllers)
├── migrations/ # SQL schema files
├── pkg/ # Shared packages (DB, etc.)
│
├── .env.example # Sample environment file
├── go.mod / go.sum # Go modules
└── README.md # This file


---

## ⚙️ Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/daffarmd/task-manager.git
cd task-manager
