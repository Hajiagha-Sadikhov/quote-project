package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"github.com/rs/cors"
)

// Quote, her bir alıntıyı temsil eder.
type Quote struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Text   string `json:"text"`
}

// quotes slice, JSON dosyasından yüklenecek alıntıları tutar.
var quotes []Quote

// Her sayfada gösterilecek alıntı sayısı.
const pageSize = 10

// loadQuotes, belirtilen dosyadan alıntıları yükler.
func loadQuotes(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&quotes)
}

// getQuotesHandler, GET /quotes endpoint'ini pagination desteğiyle işler.
// ?page=n parametresini kontrol eder; belirtilmezse 1. sayfa kullanılır.
func getQuotesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pageStr := r.URL.Query().Get("page")
	page := 1
	if pageStr != "" {
		p, err := strconv.Atoi(pageStr)
		if err == nil && p > 0 {
			page = p
		}
	}

	start := (page - 1) * pageSize
	end := start + pageSize

	// Eğer başlangıç indeksi mevcut verinin dışında ise boş dizi döndür.
	if start >= len(quotes) {
		json.NewEncoder(w).Encode([]Quote{})
		return
	}

	if end > len(quotes) {
		end = len(quotes)
	}

	json.NewEncoder(w).Encode(quotes[start:end])
}

// getQuoteByIDHandler, GET /quotes/{id} endpoint'ini işler.
func getQuoteByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// URL: /quotes/{id} şeklinde beklendiği için, id değeri path'den ayrıştırılır.
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	idStr := pathParts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid quote ID", http.StatusBadRequest)
		return
	}

	// Alıntılar arasında id eşleşmesini arar.
	for _, q := range quotes {
		if q.ID == id {
			json.NewEncoder(w).Encode(q)
			return
		}
	}

	// Eğer bulunamazsa 404 döndür.
	http.Error(w, "Quote not found", http.StatusNotFound)
}

// router, gelen isteğe göre doğru handler'a yönlendirme yapar.
func router(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/quotes" {
		getQuotesHandler(w, r)
		return
	} else if strings.HasPrefix(r.URL.Path, "/quotes/") {
		getQuoteByIDHandler(w, r)
		return
	}

	http.NotFound(w, r)
}

func main() {
	// quotes.json dosyasından alıntıları yükle.
	err := loadQuotes("quotes.json")
	if err != nil {
		log.Fatalf("Error loading quotes: %v", err)
	}
	log.Printf("Loaded %d quotes.", len(quotes))

	// CORS middleware'ini ekle
	mux := http.NewServeMux()
	mux.HandleFunc("/quotes", router)
	mux.HandleFunc("/quotes/", router)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST).
	handler := cors.Default().Handler(mux)

	// API server'ı başlat
	log.Println("Starting API server on port 5000...")
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
