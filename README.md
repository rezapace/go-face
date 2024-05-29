# Cara Menjalankan Aplikasi Face Recognition

## Persiapan

1. **Instalasi Python dan Anaconda:**
   - Pastikan Python dan Anaconda sudah terinstal di sistem Anda.

2. **Instalasi Dependensi:**
   - Buka Anaconda Prompt dan jalankan perintah berikut:
     ```sh
     conda create -n face_recognition_env python=3.8
     conda activate face_recognition_env
     pip install face_recognition flask flask-ngrok Pillow
     ```

## Menjalankan API Face Recognition

1. **Buka dan Jalankan Notebook:**
   - Buka file `Face Recognition API.ipynb` di Anaconda.
   - Jalankan semua sel di notebook tersebut untuk memulai server Flask.

## Menjalankan Aplikasi Golang

1. **Buka dan Jalankan Program Golang:**
   - Buka file `gofc/main.go` dan `gofc1/main.go`.
   - Jalankan program dengan perintah berikut di terminal:
     ```sh
     go run gofc/main.go
     go run gofc1/main.go
     ```

## Menggunakan Postman

1. **Import Koleksi Postman:**
   - Buka Postman dan import koleksi JSON yang disediakan.

2. **Menambahkan Wajah Dikenal:**
   - Gunakan endpoint `POST /send_request` untuk menambahkan wajah dikenal.
   - Isi body request dengan JSON berikut:
     ```json
     {
       "name": "Nama Orang",
       "encoding": "encoding_string"
     }
     ```

3. **Mengunggah dan Mendeteksi Gambar:**
   - Gunakan endpoint `POST /upload` untuk mengunggah gambar dan mendeteksi wajah.
   - Pilih file gambar yang ingin diunggah.

## Alur Penggunaan

1. **Deteksi Gambar:**
   - Unggah gambar menggunakan endpoint `/upload`.

2. **Input Nama Gambar:**
   - Tambahkan wajah dikenal dengan nama menggunakan endpoint `/send_request`.

3. **Deteksi Gambar:**
   - Unggah kembali gambar untuk mendeteksi wajah yang sudah dikenal.

Selamat mencoba!
deteksi gambar 
