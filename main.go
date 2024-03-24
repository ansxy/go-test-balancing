package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/consume-memory/", consumeMemoryHandler)
	http.HandleFunc("/status/200", status200Handler)
	http.HandleFunc("/status/400", status400Handler)
	http.HandleFunc("/status/500", status500Handler)

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func status200Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Status 200 - OK\n")
}

func status400Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "Status 400 - Bad Request\n")
}

func status500Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Status 500 - Internal Server Error\n")
}

func consumeMemoryHandler(w http.ResponseWriter, r *http.Request) {
	mbStr := r.URL.Path[len("/consume-memory/"):]
	mb, err := strconv.Atoi(mbStr)
	if err != nil {
		http.Error(w, "Invalid number of megabytes", http.StatusBadRequest)
		return
	}

	// Allocate memory
	data := make([]byte, mb*1024*1024)
	if data == nil {
		http.Error(w, "Failed to allocate memory", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Allocated %d MB of memory", mb)
}
