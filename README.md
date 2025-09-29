# Go API Boilerplate

Boilerplate siap pakai untuk aplikasi:

- Sekolah
- Inventaris / Gudang
- Kasir
- Pengelolaan kerja

## Fitur

- MySQL support (XAMPP & production)
- Modular structure (`internal/`)
- CORS enabled
- Standard JSON response
- Environment-based config

## Setup Development (XAMPP)

1. Jalankan MySQL di XAMPP
2. Buat database (misal: `myapp_db`)
3. Salin `.env.example` → `.env`, sesuaikan
4. Jalankan:
   ```bash
   go run cmd/server/main.go
   ```
