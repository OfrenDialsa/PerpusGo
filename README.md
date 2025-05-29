# 📚 Perpus Go

Perpus Go is a backend project developed in **Golang** using the **Fiber** web framework and **PostgreSQL** as the database. It provides a simple RESTful API with features like **JWT-based authentication** and **book management**. This project is designed as a foundational backend for a library management system or similar use cases.

## 🚀 Features

- 🔐 **JWT Authentication**
  - User login and register with secure token generation
  - Middleware-protected routes
- 📘 **Book Management**
  - CRUD operations for books (Create, Read, Update, Delete)
  - Book listing and search functionality
  - Managing Book stocks and Journals
  - Applies charges for books returned past the due date.
- 🛠️ Built with:
  - Go (Golang)
  - Fiber (Web framework)
  - PostgreSQL (Database)

## 🧱 Project Structure

```
📁 PerpusGo
├── 📁 .github
├── 📁 domain
│   ├── auth.go
│   ├── book.go
│   ├── book_stock.go
│   ├── charge.go
│   ├── customer.go
│   ├── journal.go
│   ├── media.go
│   └── user.go
├── 📁 dto
│   ├── auth_data.go
│   ├── book_data.go
│   ├── book_stock_data.go
│   ├── customer_data.go
│   ├── journal_data.go
│   ├── media_data.go
│   └── response.go
├── 📁 internal
│   ├── 📁 api
│   │   ├── auth.go
│   │   ├── book_stock.go
│   │   ├── book.go
│   │   ├── customer.go
│   │   ├── journal.go
│   │   └── media.go
│   ├── 📁 config
│   │   ├── loader.go
│   │   └── model.go
│   ├── 📁 connection
│   │   └── connection.go
│   ├── 📁 repository
│   │   ├── book_stock.go
│   │   ├── book.go
│   │   ├── charge.go
│   │   ├── customer.go
│   │   ├── journal.go
│   │   ├── media.go
│   │   └── user.go
│   ├── 📁 service
│   │   ├── auth.go
│   │   ├── book_stock.go
│   │   ├── book.go
│   │   ├── customer.go
│   │   ├── journal.go
│   │   └── media.go
│   └── 📁 util
├── 📁 sql
│   ├── books-book-stocks.sql
│   └── customer-user.sql
├── 📁 storage
├── .env
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## 🛠️ Setup & Installation

### Prerequisites

- Go 1.18+
- PostgreSQL installed and running
- Git

### Steps

1. **Clone the repository**

```bash
git clone https://github.com/OfrenDialsa/PerpusGo.git
cd PerpusGo
```

2. **Create `.env` file**

```env
SERVER_HOST=your_host
SERVER_PORT=9000
SERVER_ASSET_URL=http://localhost:9000/media

DB_HOST=your_host
DB_PORT=5432
DB_NAME=your_dbname
DB_USER=your_user
DB_PASS=your_pass
DB_TZ=Asia/Jakarta

JWT_KEY=your_secret
JWT_EXP=your_exp

STORAGE_PATH = yourpath
```

3. **Install dependencies**

```bash
go mod tidy
```

4. **Run the application**

```bash
go run main.go
```

The server will start on `http://localhost:9000`.

## 🧪 API Endpoints

### Auth

| Method | Endpoint       | Description        |
|--------|----------------|--------------------|
| POST   | `/login`       | User login         |
| POST   | `/register`    | User register      |

### Customers (Protected by JWT)

| Method | Endpoint           | Description            |
|--------|--------------------|------------------------|
| GET    | `/customers`       | List all customers     |
| GET    | `/customers/:id`   | Get customers by ID    |
| POST   | `/customers`       | Add new customer       |
| PUT    | `/customers/:id`   | Update customer data   |
| DELETE | `/customers/:id`   | Delete customer        |

### Books (Protected by JWT)

| Method | Endpoint       | Description            |
|--------|----------------|------------------------|
| GET    | `/books`       | List all books         |
| GET    | `/books/:id`   | Get book by ID         |
| POST   | `/books`       | Add a new book         |
| PUT    | `/books/:id`   | Update book details    |
| DELETE | `/books/:id`   | Delete a book          |

### Book-stocks (Protected by JWT)

| Method | Endpoint        | Description            |
|--------|-----------------|------------------------|
| POST   | `/book-stocks`  | Add a new book stock   |
| DELETE | `/book-stocks`  | Delete a book stock    |

### Journals (Protected by JWT)

| Method | Endpoint          | Description            |
|--------|-------------------|------------------------|
| GET    | `/journals`       | List all journals      |
| POST   | `/journals`       | Add a new journals     |
| PUT    | `/journals/:id`   | Update journals detail |

### Media/Book cover (Protected by JWT)

| Method | Endpoint          | Description            |
|--------|-------------------|------------------------|
| POST   | `/media`          | Upload Book cover      |
| STATIC | `/media`          | Save Book cover        |

## 🔐 Authentication

All book-related endpoints require a valid JWT token in the `Authorization` header:

```
Authorization: Bearer <token>
```

## 📘 Example Auth Object

```json
{
  "email": "nerodev@gmail.com",
  "password": "adminn"
}
```

## 🙌 Contributing

Contributions are welcome! Feel free to open issues or pull requests to improve this project.

---
