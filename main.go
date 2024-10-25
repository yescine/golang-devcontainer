package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World !!"))
	})

	mux.HandleFunc("GET /posts/{postId}", func(w http.ResponseWriter, r *http.Request) {
		postId := r.PathValue("postId")
		log.Println(postId)
		w.Write([]byte("success"))

	})

	log.Println("server stared on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}

}
