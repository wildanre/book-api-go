# Scripts untuk Mengirim 10,000 Data Buku ke API

Folder ini berisi 3 script berbeda untuk mengirim 10,000 data buku ke Books API:

## 📁 Files

1. **`post_books.go`** - Script Go (Recommended) 
2. **`post_10k_books.py`** - Script Python dengan async/await (Advanced)
3. **`simple_post_books.py`** - Script Python sederhana dengan threading
4. **`requirements.txt`** - Dependencies Python

## 🚀 Cara Penggunaan

### Persiapan

1. **Pastikan API server berjalan:**
   ```bash
   cd /Users/danuste/Desktop/kuliah/golang
   go run main.go
   ```

2. **API akan berjalan di:** `http://localhost:8080`

### Option 1: Menggunakan Script Go (Recommended)

```bash
# Masuk ke folder scripts
cd scripts

# Jalankan script Go
go run post_books.go
```

**Keuntungan:**
- ✅ Tidak perlu install dependencies tambahan
- ✅ Performance tinggi dengan goroutines
- ✅ Memory efficient
- ✅ Built-in error handling

### Option 2: Menggunakan Script Python (Advanced)

```bash
# Install dependencies
pip install -r requirements.txt

# Jalankan script async (lebih cepat)
python3 post_10k_books.py
```

**Keuntungan:**
- ✅ Async/await untuk concurrent requests
- ✅ Faker library untuk data yang lebih realistis
- ✅ Progress monitoring yang detail

### Option 3: Menggunakan Script Python Simple

```bash
# Install dependencies
pip install requests

# Jalankan script sederhana
python3 simple_post_books.py
```

**Keuntungan:**
- ✅ Code yang mudah dipahami
- ✅ Threading untuk concurrency
- ✅ Minimal dependencies

## ⚙️ Konfigurasi

Anda dapat mengubah konfigurasi di dalam masing-masing script:

### Script Go (`post_books.go`)
```go
const (
    BaseURL        = "http://localhost:8080"  // URL API
    TotalBooks     = 10000                    // Jumlah buku
    MaxWorkers     = 20                       // Jumlah worker goroutines
    BatchSize      = 100                      // Ukuran batch untuk progress
)
```

### Script Python
```python
BASE_URL = "http://localhost:8080"
TOTAL_BOOKS = 10000
CONCURRENT_REQUESTS = 50  # atau MAX_WORKERS = 20
BATCH_SIZE = 100
```

## 📊 Output Yang Diharapkan

Semua script akan menampilkan:

1. **Progress real-time:**
   ```
   📈 Progress: 1,000/10,000 (10.0%) | Speed: 45.2 books/s | ETA: 199.1s | Success: 998 | Errors: 2
   ```

2. **Hasil akhir:**
   ```
   📊 FINAL RESULTS
   ============================================================
   ✅ Successfully created: 9,987 books
   ❌ Failed to create: 13 books
   ⏱️  Total time: 235.67 seconds
   🚀 Average speed: 42.4 books/second
   📈 Success rate: 99.9%
   ```

3. **Error log:** File JSON berisi detail error (jika ada)

## 🎯 Performance Comparison

| Script | Language | Concurrency | Estimated Time | Memory Usage |
|--------|----------|-------------|----------------|--------------|
| `post_books.go` | Go | Goroutines (20) | ~3-5 minutes | Low |
| `post_10k_books.py` | Python | Async (50) | ~4-7 minutes | Medium |
| `simple_post_books.py` | Python | Threads (20) | ~5-8 minutes | Medium |

## 🔧 Troubleshooting

### API tidak response
```
❌ API is not responding. Please check if the server is running.
💡 Make sure to run: go run main.go
```

**Solusi:**
1. Pastikan server API berjalan di terminal lain
2. Check URL di `http://localhost:8080/health`
3. Pastikan port 8080 tidak digunakan aplikasi lain

### Rate limiting atau timeout
Jika terjadi banyak error:
1. Kurangi jumlah concurrent workers/requests
2. Tambah delay antar request
3. Increase timeout duration

### Memory issues (Python)
Jika script Python kehabisan memory:
1. Kurangi `CONCURRENT_REQUESTS` atau `MAX_WORKERS`
2. Gunakan script Go yang lebih memory efficient

## 📝 Data Yang Digenerate

Script akan membuat data buku dengan format:

```json
{
  "title": "Belajar Python untuk Pemula",
  "author": "Ahmad Fauzi"
}
```

**Variasi Judul:**
- Template: "Belajar Python untuk Pemula"
- Kombinasi: "Praktis Machine Learning"
- Random: "Panduan Database Lanjut"

**Variasi Author:**
- Nama Indonesia yang umum
- Kombinasi nama depan dan belakang

## 🚨 Important Notes

1. **API Server:** Pastikan server berjalan sebelum menjalankan script
2. **Database:** Script akan menambah data ke database yang sama
3. **Duplicate:** Tidak ada pengecekan duplicate, semua data akan di-insert
4. **Performance:** Waktu eksekusi tergantung performa server dan network
5. **Monitoring:** Monitor penggunaan CPU dan memory selama eksekusi

## 📈 Monitoring Progress

Semua script menyediakan:
- ✅ Real-time progress updates
- ✅ Speed monitoring (books/second)
- ✅ ETA (Estimated Time of Arrival)
- ✅ Success/error counting
- ✅ Error logging ke file

Happy coding! 🎉
