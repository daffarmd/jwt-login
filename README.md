## ⚙️ Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/daffarmd/jwt-login.git
cd jwt-login

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


