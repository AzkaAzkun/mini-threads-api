# Mini Threads API

Mini Threads API adalah clone API untuk aplikasi seperti Threads Instagram atau Twitter. API ini memungkinkan pengguna untuk membuat, membaca, memperbarui, dan menghapus postingan serta fitur lainnya seperti like dan komentar.

## Fitur Utama
- **Autentikasi Pengguna** (Registrasi & Login)
- **Manajemen Postingan** (Buat, Edit, Hapus, Lihat Postingan)
- **Manajemen Gambar** (Upload & Hapus Gambar Postingan)
- **Fitur Interaksi** (Like, Komentar)

## Teknologi yang Digunakan
- **Golang** (Gin Framework)
- **GORM** (ORM untuk database)
- **PostgreSQL** (Database utama)
- **UUID & ULID** (Untuk identifikasi unik)
- **JWT** (Autentikasi Token)

## Struktur Database (Tabel Utama)

![alt text](image-1.png)

1. **Users** - Menyimpan data pengguna
2. **Posts** - Menyimpan data postingan
3. **PostImages** - Menyimpan gambar dari postingan
4. **Likes** - Menyimpan siapa yang menyukai postingan
5. **Comments** - Menyimpan komentar pada postingan

## Getting Started
### 1. Clone Repository
```bash
git clone https://github.com/AzkaAzkun/mini-threads-api.git
cd mini-threads-api
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Konfigurasi Environment
Buat file `.env` dan atur variabel yang diperlukan seperti berikut:
```env
DB_HOST = localhost
DB_USER = postgres
DB_PASS = password
DB_NAME = db
DB_PORT = 5432
API_KEY = your_secret
```

### 4. Jalankan Server
```bash
go run main.go
```
Server akan berjalan di `http://localhost:8000`

## Endpoint API
Untuk list endpoint silahkan mengakses dokumentasi [ini](https://documenter.getpostman.com/view/34227976/2sAYkKGwzT)

## Kontribusi
Jika ingin berkontribusi, silakan fork repository ini dan buat pull request dengan perubahan yang diusulkan.

## Lisensi
Proyek ini menggunakan lisensi **MIT**.

---
**Author**: [Azka Rizqullah R.](https://github.com/AzkaAzkun)

