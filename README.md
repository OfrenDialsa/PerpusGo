# 📚 Perpus Go

Perpus Go is a backend project developed in **Golang** using the **Fiber** web framework and **PostgreSQL** as the database. It provides a simple RESTful API with features like **JWT-based authentication** and **book management**. This project is designed as a foundational backend for a library management system or similar use cases.

## 🚀 Features

- 🔐 **JWT Authentication**
  - User login with secure token generation
  - Middleware-protected routes
- 📘 **Book Management**
  - CRUD operations for books (Create, Read, Update, Delete)
  - Book listing and search functionality
- 🛠️ Built with:
  - Go (Golang)
  - Fiber (Web framework)
  - PostgreSQL (Database)

## 🧱 Project Structure

```
PerpusGo/
├── .github/
├── domain/
│   ├── auth.go
│   ├── customer.go
│   └── user.go
├── dto/
│   ├── auth_data.go
│   ├── customer_data.go
│   └── response.go
├── internal/
│   ├── api/
│   │   ├── auth.go
│   │   └── customer.go
│   ├── config/
│   │   ├── loader.go
│   │   └── model.go
│   ├── connection/
│   │   └── database.go
│   ├── repository/
│   │   ├── customer.go
│   │   └── user.go
│   ├── service/
│   │   ├── auth.go
│   │   └── customer.go
│   └── util/
│       └── validation.go
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
PORT=3000
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=perpus_go
JWT_SECRET=your_jwt_secret
```

3. **Install dependencies**

```bash
go mod tidy
```

4. **Run the application**

```bash
go run main.go
```

The server will start on `http://localhost:3000`.

## 🧪 API Endpoints

### Auth

| Method | Endpoint       | Description        |
|--------|----------------|--------------------|
| POST   | `/auth`        | User login         |

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

## 🔐 Authentication

All book-related endpoints require a valid JWT token in the `Authorization` header:

```
Authorization: Bearer <token>
```

## 📘 Example Book Object

```json
{
  "code": "M-0001",
  "name": "10 Dosa Besar Soeharto"
}
```

## 📄 License

This project is open-source and available under the [MIT License](LICENSE).

## 🙌 Contributing

Contributions are welcome! Feel free to open issues or pull requests to improve this project.

---
