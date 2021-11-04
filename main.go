package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	data := []struct {
		Id    int
		Title string
	}{
		{1, "List Artikel"},
		{2, "List Artikel 2"},
		{3, "List Artikel 3"},
		{4, "List Artikel 4"},
	}

	jsonInBytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonInBytes)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	fmt.Println("Server Running in http://127.0.0.1:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
