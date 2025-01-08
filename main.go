package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type TodoItem struct {
	ID   int    `json:"id"`
	Item string `json:"item"`
}

func main() {
	var todos = make([]TodoItem, 0)
	mux := http.NewServeMux()
	counter := 0
	mux.HandleFunc("GET /todo", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(todos)
		if err != nil {
			log.Println(err)
		}
		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
		}
	})

	mux.HandleFunc("POST /todo", func(w http.ResponseWriter, r *http.Request) {
		var t TodoItem
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}

		counter++ // increment
		t.ID = counter
		todos = append(todos, t)
		w.WriteHeader(http.StatusCreated)
		return
	})

	mux.HandleFunc("DELETE /todo", func(w http.ResponseWriter, r *http.Request) {

	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
