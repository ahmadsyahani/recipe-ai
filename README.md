# 👨‍🍳 Recipe AI: Generator Resep (Golang + Gemini API)

Selamat datang di **Study Jam #2 : Golang API Integration**!  
Project ini adalah aplikasi web sederhana berbasis **Go** yang terintegrasi dengan **Gemini API** untuk membantu kamu meracik resep masakan hanya berdasarkan bahan-bahan yang tersisa di kulkas.

---

## 🛠️ Persiapan Awal
Sebelum kita ngoding bareng, pastikan kamu sudah siapin "alat masak"-nya:
1. **Golang** sudah ter-install (Versi 1.20 ke atas disarankan).
2. **API Key Gemini**: Dapatkan secara gratis di [Google AI Studio](https://aistudio.google.com/app/apikey).
3. **Code Editor**: VS Code atau editor favoritmu.

---

## 🚀 Cara Setup

Setelah kamu melakukan **Clone** repository ini, buka terminal di dalam folder project dan ikuti perintah berikut:

### 1. Inisialisasi Module
Kalau kamu memulai dari nol atau file `go.mod` belum ada:
```bash
go mod init dapur-ai
```
### 2. Download Library
Kita butuh library `godotenv` untuk membaca file konfigurasi rahasia
```bash
go get github.com/joho/godotenv
```
### 3. Setup API Key
Buat file baru bernama `.env` di folder paling luar. Masukkan API Key kamu di dalamnya
```bash
GEMINI_API_KEY=MASUKKAN_API_KEY_KAMU_DISINI
```
### 4. Nyalakan Aplikasinya dengan perintah
```bash
go run main.go
```
Buka browser dan akses: `http://localhost:8080`

---

### 🎯 Apa yang Akan Kita Pelajari?
Di project ini, kita nggak cuma bikin UI yang bagus, tapi kita belajar inti dari Software Engineering:

`HTTP Server`: Membuat web server lokal pakai package bawaan Go (net/http).

`JSON Marshal`: Proses membungkus data Go menjadi paket JSON untuk dikirim ke Google.

`JSON Unmarshal`: Proses membongkar paket balasan dari Google kembali menjadi data yang bisa dibaca Go.

`Prompt Engineering`: Cara memberikan instruksi yang tepat ke AI agar hasilnya rapi tanpa format bintang-bintang.

`Dark Mode & Modern UI`: Implementasi UI ala Gemini dengan Tailwind CSS dan Phosphor Icons.
