package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddHandler(response http.ResponseWriter, request *http.Request) {
	queryParams := mux.Vars(request)
	response.WriteHeader(http.StatusOK)
	value1Int, _ := strconv.Atoi(queryParams["value1"])
	value2Int, _ := strconv.Atoi(queryParams["value2"])
	fmt.Fprintf(response, strconv.Itoa(value1Int+value2Int))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/add/{value1}/{value2}", AddHandler)
	log.Fatal(http.ListenAndServe(":8888", router))
}
