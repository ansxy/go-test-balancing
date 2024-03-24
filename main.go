package main

import (
	"fmt"
	"io/ioutil"
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

	// Get Ip from Here http://checkip.amazonaws.com
	ip, err := http.Get("http://checkip.amazonaws.com")
	if err != nil {
		fmt.Println("Failed to get IP:", err)
		return
	}
	defer ip.Body.Close()

	ipAddress, err := ioutil.ReadAll(ip.Body)
	if err != nil {
		fmt.Println("Failed to read IP:", err)
		return
	}

	fmt.Println("IP Address:", string(ipAddress))
	fmt.Fprintf(w, "Status 200 - OK\n"+"IP Address: "+string(ipAddress)+"\n")
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
