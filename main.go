package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Data untuk template HTML
type PageData struct {
	Ingredients string
	Response    string
	Error       string
}

// Struktur API

// RequestBox: Marshal (Bungkus ke JSON)
type RequestBox struct {
	// Request Marshal di sini
}

// ResponseBox: Unmarshal (Bongkar dari JSON)
type ResponseBox struct {
	// Response Unmarshal di sini
}

func main() {
	godotenv.Load()

	http.HandleFunc("/", handleHome)

	fmt.Println("Program jalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/index.html"))
	data := PageData{}

	if r.Method == http.MethodPost {
		// Imput dari form dan panggil Gemini API
	}

	tmpl.Execute(w, data)
}

func askGemini(userInput string) (string, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash:generateContent?key=" + apiKey

	// MARSHAL (Membungkus ke JSON)

	// KIRIM DATA (HTTP POST) ---

	// UNMARSHAL (Bongkar JSON)

	return "Hasilnya akan muncul di sini!", nil
}
