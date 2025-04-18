package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter()

	staticDir := "../frontend/dist"
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(staticDir))))

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" || r.URL.Path == "/index.html" {
			http.ServeFile(w, r, staticDir+"/index.html")
		} else {
			http.NotFound(w, r)
		}
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