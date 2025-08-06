# Books API - Go Gin CRUD Application

A RESTful API for managing books built with Go, Gin, GORM, and PostgreSQL.

## 🚀 Live Demo

**Production URL:** `https://book-api-go.zeabur.app`
**Swagger Documentation:** `https://book-api-go.zeabur.app/swagger/index.html`

## 🚀 Features

Aplikasi ini telah direfactor dengan arsitektur yang lebih terstruktur:

```
├── docs/               # Swagger documentation
├── internal/
│   ├── database/        # Database connection
│   ├── handlers/        # HTTP handlers  
│   ├── middleware/      # Custom middleware
│   ├── models/          # Data models
│   ├── routes/          # Route definitions
│   └── services/        # Business logic
├── main.go             # Entry point
├── .env.example        # Environment variables template
├── Dockerfile          # Docker configuration
├── render.yaml         # Render deployment config
└── .gitignore          # Git ignore rules
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

### ✅ API Documentation
- Swagger/OpenAPI 3.0 documentation
- Interactive API testing
- Deployed documentation available online

### ✅ Cloud Deployment
- Deployed on Zeabur platform
- Automated CI/CD pipeline
- Production-ready configuration

## 📡 Endpoints

### Health Check
```
GET /health             # Root health check
GET /api/health         # API health check
```

### Books API
```
POST   /api/books      # Create book
GET    /api/books      # Get all books (with pagination)
GET    /api/books/:id  # Get book by ID
PUT    /api/books/:id  # Update book
DELETE /api/books/:id  # Delete book
```

### Documentation
```
GET /swagger/*         # Swagger UI documentation
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

### Local Development

1. **Clone repository**
   ```bash
   git clone https://github.com/wildanre/book-api-go.git
   cd book-api-go
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Setup environment variables**
   Copy `.env.example` to `.env` dan isi dengan values yang sesuai:
   ```env
   DATABASE_URL=postgresql://username:password@localhost:5432/dbname?sslmode=disable
   PORT=8080
   GIN_MODE=debug
   ```

4. **Run aplikasi**
   ```bash
   go run main.go
   ```

5. **Generate Swagger docs (optional)**
   ```bash
   # Install swag tool
   go install github.com/swaggo/swag/cmd/swag@latest
   
   # Generate docs
   swag init
   ```

### Docker Deployment

1. **Build Docker image**
   ```bash
   docker build -t book-api-go .
   ```

2. **Run container**
   ```bash
   docker run -p 8080:8080 \
     -e DATABASE_URL="your_db_url" \
     -e GIN_MODE="release" \
     book-api-go
   ```

## 📮 API Testing

### Using Swagger UI (Recommended)
Buka `https://book-api-go.zeabur.app/swagger/index.html` untuk interactive API documentation dan testing.

### Using Postman
Import file `books_api.postman_collection.json` ke Postman untuk testing.

Collection includes:
- Health check endpoints
- Complete CRUD operations
- Proper error handling examples
- Pagination examples

### Using cURL
```bash
# Test health check
curl https://book-api-go.zeabur.app/health

# Create a book
curl -X POST https://book-api-go.zeabur.app/api/books \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Belajar Go Programming",
    "author": "John Doe"
  }'

# Get all books
curl https://book-api-go.zeabur.app/api/books

# Get book by ID
curl https://book-api-go.zeabur.app/api/books/1

# Update book
curl -X PUT https://book-api-go.zeabur.app/api/books/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Title",
    "author": "Updated Author"
  }'

# Delete book
curl -X DELETE https://book-api-go.zeabur.app/api/books/1
```

## 🔧 Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DATABASE_URL` | PostgreSQL connection string | Required |
| `PORT` | Server port | 8080 |
| `GIN_MODE` | Gin mode (debug/release) | debug |

## 🧪 Example Usage

### Create Book
```bash
curl -X POST https://book-api-go.zeabur.app/api/books \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Belajar Go Programming",
    "author": "John Doe"
  }'
```

**Response:**
```json
{
  "message": "Book created successfully",
  "data": {
    "id": 1,
    "title": "Belajar Go Programming",
    "author": "John Doe",
    "created_at": "2025-08-06T18:20:00Z",
    "updated_at": "2025-08-06T18:20:00Z"
  }
}
```

### Get All Books with Pagination
```bash
curl "https://book-api-go.zeabur.app/api/books?page=1&limit=10"
```

**Response:**
```json
{
  "message": "Books retrieved successfully",
  "data": [
    {
      "id": 1,
      "title": "Belajar Go Programming",
      "author": "John Doe",
      "created_at": "2025-08-06T18:20:00Z",
      "updated_at": "2025-08-06T18:20:00Z"
    }
  ]
}
```

## 🚀 Deployment

### Zeabur (Current)
Aplikasi ini di-deploy menggunakan [Zeabur](https://zeabur.com):

1. Connect repository ke Zeabur
2. Set environment variables di dashboard
3. Deploy otomatis dari branch `master`

### Environment Variables untuk Production
```env
DATABASE_URL=postgresql://user:password@host:port/database?sslmode=require
PORT=8080
GIN_MODE=release
```

## 🔄 Migration dari Versi Lama

Jika Anda menggunakan versi lama, aplikasi akan otomatis:
- Menambahkan kolom `created_at`, `updated_at`, `deleted_at`
- Mempertahankan data yang sudah ada
- Endpoints sudah menggunakan struktur `/api/books` yang lebih sederhana

## 🛡️ Security & Best Practices

- ✅ Input validation dengan struct tags
- ✅ Proper error handling dan response codes
- ✅ Environment variables untuk sensitive data
- ✅ CORS middleware untuk cross-origin requests
- ✅ Request logging untuk monitoring
- ✅ Panic recovery middleware
- ✅ Structured logging
- ✅ Database connection pooling
- ✅ Production-ready Docker configuration

## 📚 Dependencies

- **[Gin](https://github.com/gin-gonic/gin)**: HTTP web framework
- **[GORM](https://gorm.io/)**: ORM untuk Go
- **[PostgreSQL Driver](https://github.com/jackc/pgx)**: Database driver
- **[Validator](https://github.com/go-playground/validator)**: Request validation
- **[Godotenv](https://github.com/joho/godotenv)**: Environment variables loading
- **[Swagger](https://github.com/swaggo/gin-swagger)**: API documentation

## 🏗️ Architecture

Aplikasi ini menggunakan clean architecture dengan layers:

```
┌─────────────────┐
│   HTTP Layer    │ ← Gin handlers, middleware, routing
├─────────────────┤
│  Service Layer  │ ← Business logic, validation
├─────────────────┤
│   Model Layer   │ ← Data structures, GORM models
├─────────────────┤
│ Database Layer  │ ← PostgreSQL, GORM connection
└─────────────────┘
```

## 🚧 Future Improvements

- [x] ~~JWT Authentication~~ → Planned for v2.0
- [x] ~~Rate limiting~~ → Planned for v2.0
- [x] ✅ Pagination (Done)
- [ ] Full-text search
- [x] ✅ API documentation (Swagger) (Done)
- [ ] Unit tests
- [x] ✅ Docker support (Done)
- [ ] Redis caching
- [ ] Database migrations
- [ ] File upload for book covers

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📞 Support

Jika ada pertanyaan atau masalah, silakan buat issue di [GitHub Issues](https://github.com/wildanre/book-api-go/issues).
