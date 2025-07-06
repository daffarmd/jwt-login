# Yang sudah done adalah jwt auth nya!! (JWT AUTH)
# Yang belum adalah CRUD saja untuk TASK - MANAGER
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

## ⚙️ Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/daffarmd/task-manager.git
cd task-manager

### 2. Clone the repository
```bash
cp .env.example .env
```
Fill in your PostgreSQL credentials:
```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=taskmanager
PORT=8080
JWT_SECRET=supersecretkey
```

### 3. Run database migration
Run the SQL in **migrations/init.sql** manually using **psql**, **DBeaver**, or **pgAdmin**.


