package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello the DevOps World 🚀")
}

func main() {
	http.HandleFunc("/", hello)
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

