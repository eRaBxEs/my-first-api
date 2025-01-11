package transport

import (
	"encoding/json"
	"fmt"
	"log"
	"my-first-api/internal/todo"
	"net/http"
	"strconv"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer(todoSvc *todo.Service) *Server {
	mux := http.NewServeMux()
	counter := 0
	mux.HandleFunc("GET /todo", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(todoSvc.GetAll())
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
		err = todoSvc.Add(t)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}
		w.WriteHeader(http.StatusCreated)
	})

	mux.HandleFunc("DELETE /todo/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			log.Println(err)
			return
		}
		todos := todoSvc.GetAll()
		var found bool
		index := 0
		// get the index of the element to be deleted
		for x, task := range todos {
			if task.ID == id {
				index = x
				found = true
				break
			}
		}

		if !found {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		checks1 := todos[index+1:]
		fmt.Println("checks1:", checks1)

		checks2 := todos[:index]
		fmt.Println("checks2:", checks2)
		// deleting the item from the slice
		//todos = append(todos[:index], todos[index+1:]...)
		todoSvc.Remove(index)
		w.WriteHeader(http.StatusNoContent)

	})

	return &Server{
		mux: mux,
	}
}

func (s *Server) Serve() error {
	return http.ListenAndServe(":8080", s.mux)
}
