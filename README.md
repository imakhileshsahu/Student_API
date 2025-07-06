# 🎓 Student Management API

A lightweight RESTful API built with **Go** and **MySQL** to manage student records.

---

## 🚀 Features
✅ Add a student  
✅ Get student by ID  
✅ List all students  
✅ MySQL-backed storage  
✅ Graceful shutdown & logging  

---

## 📦 Tech Stack
- [Go](https://golang.org/) — Backend language
- [MySQL](https://www.mysql.com/) — Database
- [Fiber](https://gofiber.io/) or `net/http` — HTTP router
- [Postman](https://www.postman.com/) — API testing

---

## 🏗️ Project Structure
Student_API_Project/
├── cmd/student-api/main.go # Application entrypoint
├── internal/
│ ├── config/config.go # Config loader (YAML)
│ ├── mysql/mysql.go # MySQL implementation
│ ├── storage/storage.go # Storage interface
│ ├── type/types.go # Data types
│ └── http/handlers/student.go # HTTP handlers
├── config/local.yaml # Local config file
├── go.mod, go.sum # Dependencies
├── README.md # Project documentation

### 1️⃣ Clone the repository
```bash
git clone https://github.com/YOUR_USERNAME/student-api.git
cd student-api

2️⃣ Setup MySQL database
Run these commands in your MySQL:


CREATE DATABASE student_api;

CREATE USER IF NOT EXISTS 'studentuser'@'localhost' IDENTIFIED BY 'Student@123';

GRANT ALL PRIVILEGES ON student_api.* TO 'studentuser'@'localhost';

FLUSH PRIVILEGES;


3️⃣ Update config
Edit config/local.yaml:


env: "dev"
http_server:
  address: "localhost:8082"
db_user: "studentuser"
db_password: "Student@123"
db_host: "127.0.0.1"
db_port: "3306"
db_name: "student_api"


4️⃣ Run locally

go run ./cmd/student-api/main.go --config ./config/local.yaml

🌐 API Endpoints


| Method | Endpoint             | Description          |
| ------ | -------------------- | -------------------- |
| POST   | `/api/students`      | Create a new student |
| GET    | `/api/students`      | Get all students     |
| GET    | `/api/students/{id}` | Get student by ID    |


🧪 Example Requests
📋 Create a student

curl -X POST http://localhost:8082/api/students \
-H "Content-Type: application/json" \
-d '{"name":"Alice", "email":"alice@example.com", "age":20}'


📋 Get all students

curl http://localhost:8082/api/students
📋 Get a student by ID

curl http://localhost:8082/api/students/1





















