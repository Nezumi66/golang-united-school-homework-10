package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/bad", handleBadRequest).Methods(http.MethodGet)
	router.HandleFunc("/name/{PARAM}", handleHelloRequest).Methods(http.MethodGet)
	router.HandleFunc("/data", handleBodyDataRequest).Methods(http.MethodPost)
	router.HandleFunc("/headers", handleHeaderRequest).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func handleBadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func handleHelloRequest(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["PARAM"]
	_, err := w.Write([]byte(fmt.Sprintf("Hello, %v!", param)))
	if err != nil {
		log.Println("handleHelloRequest error")
	}
}

func handleBodyDataRequest(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("handleBodyDataRequest error")
	}
	str := "I got message:\n" + string(data)
	_, err = w.Write([]byte(str))
	if err != nil {
		log.Println("handleBodyDataRequest error")
	}
}

func handleHeaderRequest(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.Header.Get("a"))
	b, _ := strconv.Atoi(r.Header.Get("b"))
	w.Header().Add("a+b", strconv.Itoa(a+b))
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
