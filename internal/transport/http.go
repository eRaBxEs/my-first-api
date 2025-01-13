package transport

import (
	"encoding/json"
	"log"
	"my-first-api/internal/todo"
	"net/http"
	"strconv"
)

type TodoItem struct {
	Item string `json:"item"`
}
type Server struct {
	mux *http.ServeMux
}

func NewServer(todoSvc *todo.Service) *Server {
	mux := http.NewServeMux()
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
		var t TodoItem
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = todoSvc.Add(t.Item)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}
		w.WriteHeader(http.StatusCreated)
	})

	mux.HandleFunc("GET /search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		if query == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		results := todoSvc.Search(query)
		b, err := json.Marshal(results) // marshal it into a byte stream
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
			return
		}

	})

	mux.HandleFunc("DELETE /todo/{id}", func(w http.ResponseWriter, r *http.Request) {
		_, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			log.Println(err)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	})

	return &Server{
		mux: mux,
	}
}

func (s *Server) Serve() error {
	return http.ListenAndServe(":8080", s.mux)
}
