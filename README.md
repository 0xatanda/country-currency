```markdown
# 🌍 Country Currency API

A simple RESTful API built with **Go (Golang)**, **Gin**, and **GORM** that fetches and stores country data (including currencies, region, and capital) into a **PostgreSQL** database.

---

## 🚀 Features

- Fetches live country data from the [REST Countries API](https://restcountries.com/v3.1/all)
- Stores countries and currencies into PostgreSQL
- Provides endpoints to:
  - Refresh countries from external API
  - Retrieve all countries from the database
- Built using a modular structure (`cmd/`, `internal/` folders)
- Uses GORM ORM for database operations

---

## 🧩 Tech Stack

| Layer | Technology |
|-------|-------------|
| Language | Go 1.21+ |
| Framework | Gin Web Framework |
| ORM | GORM |
| Database | PostgreSQL |
| Deployment | Localhost / Docker-ready |

---

## 📁 Project Structure

```

country-currency-api/
├── cmd/
│   └── main.go              # App entrypoint
├── internal/
│   ├── database/
│   │   └── db.go            # Database connection logic
│   ├── handlers/
│   │   └── country_handler.go # HTTP request handlers
│   ├── models/
│   │   └── country.go       # Country model definition
│   └── services/
│       └── country_service.go # Business logic for fetching/storing data
├── go.mod
├── go.sum
└── README.md

````

---

## ⚙️ Setup Instructions

### 1️⃣ Clone the repository

```bash
git clone https://github.com/<your-username>/country-currency-api.git
cd country-currency-api
````

### 2️⃣ Install dependencies

```bash
go mod tidy
```

### 3️⃣ Setup PostgreSQL

Ensure PostgreSQL is running and create a new database:

```bash
createdb country_currency_api
```

### 4️⃣ Configure environment variables

Create a `.env` file in the project root:

```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=country_currency_api
```

---

## ▶️ Run the API

```bash
go run ./cmd/main.go
```

You should see:

```
✅ Database connected successfully
✅ Server running on port 8080
```

---

## 🧪 Test the API

### 1️⃣ Refresh Countries

Fetches all countries from external API and saves to PostgreSQL.

```bash
curl -X POST http://localhost:8080/countries/refresh
```

Expected response:

```json
{"message": "Countries refreshed successfully"}
```

### 2️⃣ Get All Countries

```bash
curl http://localhost:8080/countries
```

Expected response:

```json
[
  {
    "id": 1,
    "name": "Nigeria",
    "capital": "Abuja",
    "region": "Africa",
    "currency": "NGN"
  },
  ...
]
```

---

## 🧾 Example API Routes

| Method | Endpoint             | Description                          |
| ------ | -------------------- | ------------------------------------ |
| `POST` | `/countries/refresh` | Refresh countries from external API  |
| `GET`  | `/countries`         | Get all countries stored in database |

---

## 🧰 Developer Commands

### Run tests

```bash
go test ./...
```

### Format code

```bash
go fmt ./...
```

### Run linter (optional)

```bash
golangci-lint run
```

---

## 🗃️ Git Workflow

```bash
git add .
git commit -m "Fix panic in RefreshCountries and improve error handling"
git push origin main
```

---

## 🧑‍💻 Author

**Atanda Nafiu**
DevOps & Backend Developer
📧 [Your Email]
🔗 [Your GitHub](https://github.com/<your-username>)

---

## 🧠 Notes

* Make sure PostgreSQL service is active before running.
* Handle `.env` secrets securely in production.
* The app uses safe type assertions to avoid nil panic errors when API data fields are missing.

---

## 🏁 License

MIT License © 2025 Atanda Nafiu


