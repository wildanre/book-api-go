# Books API - Refactored

Sebuah REST API sederhana untuk manajemen buku menggunakan Go, Gin, GORM, dan PostgreSQL.

## 🏗️ Arsitektur

Aplikasi ini telah direfactor dengan arsitektur yang lebih terstruktur:

```
├── internal/
│   ├── database/        # Database connection
│   ├── handlers/        # HTTP handlers
│   ├── middleware/      # Custom middleware
│   ├── models/          # Data models
│   ├── routes/          # Route definitions
│   └── services/        # Business logic
├── main.go             # Entry point
├── .env               # Environment variables
└── .gitignore         # Git ignore rules
```

## 🚀 Fitur Baru

### ✅ Struktur yang Lebih Baik
- **Separation of Concerns**: Model, Handler, Service terpisah
- **Clean Architecture**: Dependency injection pattern
- **Modular**: Mudah untuk di-maintain dan extend

### ✅ Error Handling yang Lebih Baik
- Consistent error responses
- Proper HTTP status codes
- Validation errors

### ✅ Enhanced Model
- Timestamps (created_at, updated_at)
- Soft delete (deleted_at)
- Request validation

### ✅ Middleware
- Custom logger
- CORS support
- Panic recovery

### ✅ API Versioning
- `/api/v1/` prefix untuk endpoint baru
- Backward compatibility dengan endpoint lama

## 📡 Endpoints

### Health Check
```
GET /health
```

### API v1 (Recommended)
```
POST   /api/v1/books      # Create book
GET    /api/v1/books      # Get all books
GET    /api/v1/books/:id  # Get book by ID
PUT    /api/v1/books/:id  # Update book
DELETE /api/v1/books/:id  # Delete book
```

### Legacy Endpoints (Backward Compatible)
```
POST   /books      # Create book
GET    /books      # Get all books
GET    /books/:id  # Get book by ID
PUT    /books/:id  # Update book
DELETE /books/:id  # Delete book
```

## 📋 Response Format

### Success Response
```json
{
  "message": "Success message",
  "data": { ... }
}
```

### Error Response
```json
{
  "error": "Error message"
}
```

## 🗄️ Database Schema

```sql
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);
```

## 🛠️ Setup & Installation

1. **Clone dan masuk ke direktori**
   ```bash
   cd /path/to/your/golang/project
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Setup environment variables**
   Buat file `.env`:
   ```env
   DATABASE_URL=your_postgresql_connection_string
   PORT=8080
   GIN_MODE=debug
   ```

4. **Run aplikasi**
   ```bash
   go run main.go
   ```

## 📮 Testing dengan Postman

Import file `books_api.postman_collection.json` ke Postman untuk testing.

Collection includes:
- Health check
- CRUD operations (v1 dan legacy)
- Proper error handling examples

## 🔧 Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DATABASE_URL` | PostgreSQL connection string | Required |
| `PORT` | Server port | 8080 |
| `GIN_MODE` | Gin mode (debug/release) | debug |

## 🧪 Example Usage

### Create Book
```bash
curl -X POST http://localhost:8080/api/v1/books \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Belajar Go",
    "author": "Danu"
  }'
```

### Get All Books
```bash
curl http://localhost:8080/api/v1/books
```

## 🔄 Migration dari Versi Lama

Jika Anda menggunakan versi lama, aplikasi akan otomatis:
- Menambahkan kolom `created_at`, `updated_at`, `deleted_at`
- Mempertahankan data yang sudah ada
- Support endpoint lama untuk backward compatibility

## 🛡️ Security & Best Practices

- ✅ Input validation
- ✅ Proper error handling
- ✅ Environment variables untuk sensitive data
- ✅ CORS middleware
- ✅ Request logging
- ✅ Panic recovery

## 📚 Dependencies

- **Gin**: HTTP web framework
- **GORM**: ORM untuk Go
- **PostgreSQL**: Database driver
- **Validator**: Request validation
- **Godotenv**: Environment variables loading

## 🚧 Future Improvements

- [ ] JWT Authentication
- [ ] Rate limiting
- [ ] Pagination
- [ ] Full-text search
- [ ] API documentation (Swagger)
- [ ] Unit tests
- [ ] Docker support
