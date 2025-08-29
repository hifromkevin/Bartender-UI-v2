package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(enableCORS)

	staticDir := "../frontend/dist"

	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.Dir(staticDir))))
	router.Handle("/images/*", http.StripPrefix("/images/", http.FileServer(http.Dir("backend/public/images"))))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, staticDir+"/index.html")
	})

	router.Get("/catchphrase/string", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Mix smarter, drink better."))
	})

	router.Get("/catchphrase", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"catchphrase": "Mix smarter.", "byline": "Drink better."}`))
	})

	port := os.Getenv("GO_PORT")
	if port == "" {
		port = "3003"
	}

	log.Printf("Server started on :%s\n", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
