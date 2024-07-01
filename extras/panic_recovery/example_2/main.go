package main

import (
	"log"
	"net/http"
)

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handler middleware started")
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[PANIC] Recovered: %v\n", r)
				// debug.PrintStack()
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)

			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("Panic")
	})

	log.Println("Listening on", ":3000")
	if err := http.ListenAndServe(":3000", recoverMiddleware(mux)); err != nil {
		log.Fatalf("Could not liste on %s: %v\n", ":3000", err)
	}
}
