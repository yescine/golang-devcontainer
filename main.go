package main

import (
	"log"
	"net/http"
)

func logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("REQ", r.Method, r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}

func main() {

	mux := http.NewServeMux()
	muxWithLogger := logger(mux)

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World !!"))
	})

	mux.HandleFunc("GET /posts/{postId}", func(w http.ResponseWriter, r *http.Request) {
		postId := r.PathValue("postId")
		log.Println("POST ID", postId)
		w.Write([]byte("success"))

	})

	log.Println("server stared on port 8080")
	if err := http.ListenAndServe(":8080", muxWithLogger); err != nil {
		panic(err)
	}

}
