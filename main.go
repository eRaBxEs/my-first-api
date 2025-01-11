package main

import (
	"encoding/json"
	"log"
	"my-first-api/internal/todo"
	"net/http"
	"strconv"
)

func main() {
	svc := todo.NewService()
	mux := http.NewServeMux()
	counter := 0
	mux.HandleFunc("GET /todo", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(svc.GetAll())
		if err != nil {
			log.Println(err)
		}
		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
		}
	})

	mux.HandleFunc("POST /todo", func(w http.ResponseWriter, r *http.Request) {
		var t todo.TodoItem
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		counter++ // increment
		t.ID = counter
		svc.Add(t)
		w.WriteHeader(http.StatusCreated)
	})

	mux.HandleFunc("DELETE /todo/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			log.Println(err)
			return
		}
		todos := svc.GetAll()
		found := false
		for _, x := range svc.GetAll() {
			if x.ID == id {
				found = true
				todos = append(todos[:id-1], todos[id:]...)
				break
			}
		}
		if !found {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
