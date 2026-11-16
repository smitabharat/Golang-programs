# 🚀 Go Practice Programs

A collection of beginner to intermediate Go (Golang) programs covering language basics, algorithms, concurrency and REST APIs.

Each program lives in its own folder, so every one of them can be run on its own.

---

## 📁 Project Structure

```
Golang-programs/
├── basics/          Go language fundamentals
├── algorithms/      Common coding interview problems
├── concurrency/     Goroutines, channels, WaitGroup, worker pool
└── web/             REST APIs with net/http, Fiber and Gin
```

---

## 🧩 Programs Included

### 🔤 Basics

| Program | What it shows |
|---|---|
| `basics/constants` | Implicit repetition of values in a `const` block |
| `basics/slice_len_cap` | How `len` and `cap` change when re-slicing and appending |
| `basics/interface_shapes` | Interfaces with `Rectangle` and `Circle` shapes |
| `basics/custom_error` | Creating a custom error type and checking it with `errors.As` |
| `basics/panic_recover` | How `panic` stops execution and `recover` catches it |
| `basics/department_salary` | Total salary per department using a struct slice |
| `basics/sort_and_slice` | Sort random integers and print from index 2 to the end |
| `basics/sort_strings` | Split, sort and capitalise a comma separated list |

### 🔢 Algorithms

| Program | What it shows |
|---|---|
| `algorithms/two_sum` | Two Sum using a hash map, O(n) |
| `algorithms/valid_parentheses` | Bracket matching with a stack |
| `algorithms/longest_substring` | Longest substring without repeating characters (sliding window) |
| `algorithms/next_greater_element` | Next greater element with a monotonic stack |
| `algorithms/next_greater_circular` | Next greater element in a circular array |
| `algorithms/prime_numbers` | Prime numbers from 1 to 100 |
| `algorithms/square_recursion` | Square of a number using recursion only |
| `algorithms/count_characters` | Frequency of each character in a string |
| `algorithms/count_integers` | Frequency of each integer in a slice |
| `algorithms/first_non_repeating_char` | First character that appears only once |
| `algorithms/reverse_words` | Reverse every word but keep word order |
| `algorithms/reverse_slice` | Reversed copy of a slice |
| `algorithms/reverse_string_vowels` | Vowel positions plus the reversed string |

### ⚙️ Concurrency

| Program | What it shows |
|---|---|
| `concurrency/channel_send` | A goroutine sending numbers over a channel |
| `concurrency/even_odd_alternate` | Two goroutines printing odd and even numbers in turn |
| `concurrency/sum_even_squares` | Sum of squares of even numbers with goroutines |
| `concurrency/worker_pool` | Worker pool pattern with jobs and results channels |
| `concurrency/url_checker` | Check several websites at the same time |

### 🌐 Web / REST APIs

| Program | Framework | Endpoint | Port |
|---|---|---|---|
| `web/employee_server` | net/http | `GET /employee` | 8080 |
| `web/odd_even_server` | net/http | `GET /numbers` | 8080 |
| `web/fiber_add_api` | Fiber | `POST /add` | 3000 |
| `web/fiber_employee_post` | Fiber | `POST /employee` | 3000 |
| `web/meeting_scheduler` | Gin | `POST /schedule`, `GET /users/:userId/calendar` | 8080 |

**Fiber Addition API sample request** (`POST /add`):

```json
{
  "num1": 10,
  "num2": 20
}
```

**Response:**

```json
{
  "num1": 10,
  "num2": 20,
  "result": 30
}
```

**Meeting scheduler sample request** (`POST /schedule`):

```json
{
  "participantIds": ["u1", "u2"],
  "durationMinutes": 60,
  "timeRange": {
    "start": "2025-08-02T09:00:00Z",
    "end": "2025-08-02T18:00:00Z"
  }
}
```

---

## 🛠️ Prerequisites

* Go 1.22 or newer

```bash
go version
```

---

## ▶️ How to Run

Clone the repository and download the dependencies (only needed for the Fiber and Gin programs):

```bash
git clone https://github.com/smitabharat/Golang-programs.git
cd Golang-programs
go mod tidy
```

Run any program by its folder:

```bash
go run ./algorithms/two_sum
go run ./basics/interface_shapes
go run ./concurrency/worker_pool
go run ./web/employee_server
go run ./web/fiber_add_api
```

---

## 🌱 Learning Goals

This repository helps you understand:

* Building REST APIs in Go
* Using the standard library vs third-party frameworks
* JSON handling
* Goroutines, channels and synchronisation
* Basic algorithms and data structures
* Project organisation in Go

---

## 🤝 Contributing

Contributions are welcome! Feel free to:

* Improve code
* Add new examples
* Optimise existing solutions

---

## 📜 License

This project is open-source and available under the MIT License.

---

## 👨‍💻 Author

Smita Bharat

GitHub: [https://github.com/smitabharat](https://github.com/smitabharat)

---

⭐ If you find this helpful, consider giving it a star!
