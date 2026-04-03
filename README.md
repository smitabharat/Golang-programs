
# 🚀 Go Practice Projects

A collection of beginner-to-intermediate Go (Golang) programs demonstrating core concepts such as HTTP servers, REST APIs, constants, and algorithms.

---

## 🧩 Programs Included

### 1. 📡 Employee HTTP Server
- Built using Go’s standard `net/http` package  
- Returns a list of employees in JSON format  
- Endpoint: `GET /employee`  
- Runs on: `http://localhost:8080`

---

### 2. ⚡ Fiber Addition API
- Built using the Fiber web framework  
- Accepts JSON input and returns the sum  
- Endpoint: `POST /add`  
- Runs on: `http://localhost:3000`

**Sample Request:**
```json
{
  "num1": 10,
  "num2": 20
}
````

**Response:**

```json
{
  "num1": 10,
  "num2": 20,
  "result": 30
}
```

---

### 3. 🔢 Two Sum Algorithm

* Solves the classic "Two Sum" problem
* Uses a hashmap for efficient lookup
* Time Complexity: `O(n)`

---

### 4. 🧮 Constants Example

* Demonstrates Go constant behavior
* Shows how implicit constant values work

---

## 🛠️ Prerequisites

* Go installed (version 1.18+ recommended)

Check version:

```bash
go version
```

---

## ▶️ How to Run

Run each program individually:

```bash
go run employee_server.go
go run fiber_add_api.go
go run two_sum.go
go run constants_example.go
```

---

## 📦 Install Dependencies (Fiber)

Before running Fiber API:

```bash
go mod init go-practice
go get github.com/gofiber/fiber/v2
```

---

## 🌱 Learning Goals

This repository helps you understand:

* Building REST APIs in Go
* Using standard vs third-party frameworks
* JSON handling
* Basic algorithms and data structures
* Project organization in Go

---

## 🤝 Contributing

Contributions are welcome! Feel free to:

* Improve code
* Add new examples
* Optimize existing solutions

---

## 📜 License

This project is open-source and available under the MIT License.

---

## 👨‍💻 Author

Smita Bharat

GitHub: [https://github.com/smitabharat](https://github.com/smitabharat)

---

⭐ If you find this helpful, consider giving it a star!

```
