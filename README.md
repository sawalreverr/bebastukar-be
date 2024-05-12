# Bebastukar Backend

## About Project

Bebastukar adalah platform inovatif yang memfasilitasi pertukaran barang bekas secara online dengan cara yang efisien dan praktis. Dengan Bebastukar, pengguna dapat dengan mudah mendaftar dan masuk ke akun mereka, membuka diskusi menarik tentang berbagai topik terkait barang bekas, dan berbagi informasi dengan pengguna lain melalui fitur komentar yang interaktif.

Selain itu, pengguna dapat menambahkan postingan untuk barang bekas yang ingin mereka tukarkan, menjadikan proses pertukaran lebih mudah dan menyenangkan. Kami juga menyediakan fitur chat yang nyaman, memungkinkan pengguna untuk berkomunikasi secara langsung dengan pengguna lain tentang barang yang ingin ditukar. Dengan Bebastukar, pengalaman pertukaran barang bekas anda akan menjadi lebih menyenangkan dan efisien daripada sebelumnya.

## Features

### MVP Features

#### Autentikasi:

- Pengguna dapat membuat akun baru dengan menyediakan nama pengguna, email, dan kata sandi.
- Pengguna dapat masuk ke dalam akun mereka menggunakan kredensial yang mereka daftarkan sebelumnya.

#### Profil pengguna:

- Pengguna dapat melihat data mereka pada halaman profil.
- Pengguna dapat mengubah data pribadi mereka pada halaman ubah profil.
- Pengguna dapat mengunggah avatar/foto pribadi mereka pada halaman profil.

#### Forum Diskusi:

- Pengguna dapat membuat dan mengikuti topik diskusi terkait pertukaran barang bekas.
- Pengguna dapat memposting pertanyaan, saran, atau pengalaman terkait dengan pertukaran barang bekas.
- Pengguna dapat mengubah konten dari topik diskusi mereka jikalau adanya sesuatu yang perlu diubah.
- Pengguna dapat menghapus topik diskusi mereka jikalau diskusi tersebut sudah dianggap tidak perlu atau diluar batasan peraturan.
- Pengguna dapat memberikan, mengubah, dan menghapus komentar terhadap postingan pengguna lain.
- Pengguna dapat membalas, mengubah, menghapus komentar dari pengguna lain (reply comment).

#### Chat Bot:

- Pengguna dapat menanyakan terkait pengelolaan atau contoh barang bekas yang layak untuk ditukar kepada AI.

### Penggunaan Publik:

- Pengguna dapat melihat profil pengguna lain dengan menggunakan ID.
- Pengguna dapat melihat diskusi pengguna lain dengan menggunakan ID.
- Pengguna dapat melihat semua topik diskusi pengguna lain dengan userID.
- Pengguna dapat melihat semua topik diskusi dengan pagination (page, limit, sort_by, sort_type).
- Pengguna dapat melihat semua komentar yang ada pada 1 topik diskusi dan dapat melihat reply komen yang ada.

### Other Features

#### Posting Barang Bekas:

- Pengguna dapat membuat postingan untuk menawarkan barang bekas yang mereka ingin tukarkan.
- Postingan barang dapat berisi informasi seperti deskripsi barang, kondisi barang, lokasi, dan foto.

#### Pencarian dan Filter Postingan:

- Pengguna dapat mencari postingan berdasarkan kata kunci, kategori barang, atau lokasi.
- Pengguna dapat menyaring hasil pencarian berdasarkan kriteria tertentu seperti kategori atau lokasi.

#### Pesan Pribadi:

- Pengguna dapat mengirim dan menerima pesan pribadi dari pengguna lain untuk mendiskusikan detail pertukaran barang.

#### Admin:

- Admin dapat melihat keseluruhan data pengguna.
- Admin dapat memblokir pengguna yang melanggar peraturan komunitas.

## Tech Stacks

- [Echo](https://github.com/labstack/echo) (Web Framework Go)
- [Cloudinary](https://github.com/cloudinary/cloudinary-go/) (Cloud storage free)
- [Viper](https://github.com/spf13/viper) (Configuration)
- [Validator](https://github.com/go-playground/validator) (Type validation)
- [JWT](https://github.com/golang-jwt/jwt) (Middleware)
- [Generative AI](https://github.com/google/generative-ai-go) (Chat Bot)
- MySQL (SQL)
- [GORM](https://gorm.io/docs/) (ORM)

## API Documentation

sertakan dokumentasi API yang dibuat dengan menggunakan postman / swagger.

## ERD

![ERD - Picture](<docs/(ERD)%20Mini%20Project%20-%20BebasTukar.drawio.png>)

[ERD - draw.io](https://drive.google.com/file/d/1njk9cH9IgRjSxqUGDRB2-1rVvhaYRWvo/view?usp=sharing)

## Setup

sebutkan cara menggunakan project ini di lokal
