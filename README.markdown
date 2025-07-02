# Aplikasi Backend Catatan Belajar

Aplikasi ini adalah backend RESTful API yang dibangun menggunakan Go dengan framework Gin dan GORM sebagai ORM untuk PostgreSQL. Aplikasi ini mengelola data pengguna, sumber daya, catatan, kuis, dan kartu flash berdasarkan skema database berikut:

- **users**: Menyimpan informasi pengguna (ID, Google ID, email, nama lengkap, URL avatar, dll.).
- **resources**: Menyimpan sumber daya yang terkait dengan pengguna (misalnya, artikel atau video).
- **notes**: Menyimpan catatan yang terkait dengan sumber daya dan pengguna.
- **quizzes**: Menyimpan kuis yang terkait dengan catatan.
- **flashcards**: Menyimpan kartu flash yang terkait dengan catatan.

## Struktur Direktori

```plaintext
project/
├── config/
│   └── config.go
├── model/
│   └── model.go
├── dto/
│   ├── user_request.go
│   ├── user_response.go
│   ├── resource_request.go
│   ├── resource_response.go
│   ├── note_request.go
│   ├── note_response.go
│   ├── quiz_request.go
│   ├── quiz_response.go
│   ├── flashcard_request.go
│   ├── flashcard_response.go
├── controller/
│   ├── user_controller.go
│   ├── user_controller_impl.go
│   ├── resource_controller.go
│   ├── resource_controller_impl.go
│   ├── note_controller.go
│   ├── note_controller_impl.go
│   ├── quiz_controller.go
│   ├── quiz_controller_impl.go
│   ├── flashcard_controller.go
│   ├── flashcard_controller_impl.go
├── repository/
│   ├── user_repository.go
│   ├── user_repository_impl.go
│   ├── resource_repository.go
│   ├── resource_repository_impl.go
│   ├── note_repository.go
│   ├── note_repository_impl.go
│   ├── quiz_repository.go
│   ├── quiz_repository_impl.go
│   ├── flashcard_repository.go
│   ├── flashcard_repository_impl.go
├── service/
│   ├── user_service.go
│   ├── user_service_impl.go
│   ├── resource_service.go
│   ├── resource_service_impl.go
│   ├── note_service.go
│   ├── note_service_impl.go
│   ├── quiz_service.go
│   ├── quiz_service_impl.go
│   ├── flashcard_service.go
│   ├── flashcard_service_impl.go
├── main.go
├── go.mod
├── go.sum
└── README.md
```

## Prasyarat

- **Go**: Versi 1.16 atau lebih baru.
- **PostgreSQL**: Database yang dijalankan secara lokal atau di cloud.
- **Dependensi Go**:
  ```bash
  go get github.com/gin-gonic/gin
  go get gorm.io/gorm
  go get gorm.io/driver/postgres
  go get github.com/google/uuid
  ```

## Pengaturan

1. **Inisialisasi Proyek**:
   ```bash
   go mod init your_project
   go mod tidy
   ```

2. **Konfigurasi Database**:
   - Buat database PostgreSQL (misalnya, `mydb`).
   - Atur variabel lingkungan `DATABASE_URL` atau ubah `dsn` di `config/config.go`:
     ```go
     dsn := "host=localhost user=postgres password=postgres dbname=mydb port=5432 sslmode=disable"
     ```

3. **Migrasi Database**:
   Tambahkan kode berikut di `main.go` setelah inisialisasi database untuk membuat tabel:
   ```go
   db.AutoMigrate(&model.User{}, &model.Resource{}, &model.Note{}, &model.Quiz{}, &model.Flashcard{})
   ```

4. **Jalankan Aplikasi**:
   ```bash
   go run main.go
   ```
   Server akan berjalan di `http://localhost:8080`.

5. **Mode Produksi**:
   Untuk lingkungan produksi, atur mode release:
   ```bash
   export GIN_MODE=release
   ```
   Atau tambahkan di `main.go`:
   ```go
   gin.SetMode(gin.ReleaseMode)
   ```

   Atur juga proxy tepercaya:
   ```go
   r.SetTrustedProxies([]string{"127.0.0.1"})
   ```

## Endpoint API

API menyediakan operasi CRUD untuk setiap entitas. Berikut adalah daftar endpoint:

### Users
- **GET /api/users**: Mengambil semua pengguna.
- **GET /api/users/:id**: Mengambil pengguna berdasarkan ID.
- **POST /api/users**: Membuat pengguna baru.
  ```json
  {
      "google_id": "12345",
      "email": "user@example.com",
      "full_name": "Nama Pengguna",
      "avatar_url": "https://example.com/avatar.jpg"
  }
  ```
- **PUT /api/users/:id**: Memperbarui pengguna.
- **DELETE /api/users/:id**: Menghapus pengguna.

### Resources
- **GET /api/resources**: Mengambil semua sumber daya.
- **GET /api/resources/:id**: Mengambil sumber daya berdasarkan ID.
- **POST /api/resources**: Membuat sumber daya baru.
  ```json
  {
      "user_id": "UUID_PENGGUNA",
      "type": "article",
      "source_url": "https://example.com",
      "original_title": "Judul Artikel",
      "status": "active"
  }
  ```
- **PUT /api/resources/:id**: Memperbarui sumber daya.
- **DELETE /api/resources/:id**: Menghapus sumber daya.

### Notes
- **GET /api/notes**: Mengambil semua catatan.
- **GET /api/notes/:id**: Mengambil catatan berdasarkan ID.
- **POST /api/notes**: Membuat catatan baru.
  ```json
  {
      "resource_id": "UUID_SUMBER_DAYA",
      "user_id": "UUID_PENGGUNA",
      "title": "Judul Catatan",
      "summary": "Ringkasan catatan",
      "full_text": "Teks lengkap catatan"
  }
  ```
- **PUT /api/notes/:id**: Memperbarui catatan.
- **DELETE /api/notes/:id**: Menghapus catatan.

### Quizzes
- **GET /api/quizzes**: Mengambil semua kuis.
- **GET /api/quizzes/:id**: Mengambil kuis berdasarkan ID.
- **POST /api/quizzes**: Membuat kuis baru.
  ```json
  {
      "note_id": "UUID_CATATAN",
      "question": "Apa ibu kota Indonesia?",
      "options": "[\"Jakarta\", \"Bandung\", \"Surabaya\", \"Medan\"]",
      "correct_answer_index": 0,
      "explanation": "Jakarta adalah ibu kota Indonesia."
  }
  ```
- **PUT /api/quizzes/:id**: Memperbarui kuis.
- **DELETE /api/quizzes/:id**: Menghapus kuis.

### Flashcards
- **GET /api/flashcards**: Mengambil semua kartu flash.
- **GET /api/flashcards/:id**: Mengambil kartu flash berdasarkan ID.
- **POST /api/flashcards**: Membuat kartu flash baru.
  ```json
  {
      "note_id": "UUID_CATATAN",
      "front_text": "Ibu kota Indonesia",
      "back_text": "Jakarta"
  }
  ```
- **PUT /api/flashcards/:id**: Memperbarui kartu flash.
- **DELETE /api/flashcards/:id**: Menghapus kartu flash.

## Contoh Pengujian dengan curl

1. **Membuat Kuis**:
   ```bash
   curl -X POST http://localhost:8080/api/quizzes \
   -H "Content-Type: application/json" \
   -d '{"note_id": "UUID_CATATAN", "question": "Apa ibu kota Indonesia?", "options": "[\"Jakarta\", \"Bandung\", \"Surabaya\", \"Medan\"]", "correct_answer_index": 0, "explanation": "Jakarta adalah ibu kota Indonesia."}'
   ```

2. **Membuat Kartu Flash**:
   ```bash
   curl -X POST http://localhost:8080/api/flashcards \
   -H "Content-Type: application/json" \
   -d '{"note_id": "UUID_CATATAN", "front_text": "Ibu kota Indonesia", "back_text": "Jakarta"}'
   ```

3. **Mengambil Semua Kuis**:
   ```bash
   curl http://localhost:8080/api/quizzes
   ```

## Catatan

- **Validasi Kunci Asing**: Pastikan `user_id`, `resource_id`, dan `note_id` valid sebelum membuat entitas terkait.
- **Kesalahan Umum**:
  - Jika muncul error `record not found`, periksa apakah ID yang digunakan ada di database.
  - Jika koneksi database gagal, pastikan `DATABASE_URL` atau `dsn` di `config/config.go` benar.
- **Ekstensi**: Untuk menambahkan otentikasi, validasi JSON untuk `options` pada kuis, atau fitur lain, silakan modifikasi kode sesuai kebutuhan.

## Kontribusi

1. Fork repositori ini.
2. Buat branch baru (`git checkout -b fitur-baru`).
3. Commit perubahan Anda (`git commit -m "Menambahkan fitur baru"`).
4. Push ke branch (`git push origin fitur-baru`).
5. Buat Pull Request.

## Lisensi

Proyek ini dilisensikan di bawah [MIT License](LICENSE).