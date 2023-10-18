package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/go-playground/validator"
)

type Comment struct {
	Message  string `validate:" required,min=1, max=140"`
	UserName string `validate:" required,min=1, max=15"`
}

func main() {
	var mutex = &sync.RWMutex{}
	comments := make([]Comment, 0, 100)
	http.HandleFunc("/comments", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			mutex.RLock()

			if err := json.NewEncoder(w).Encode(comments); err != nil {
				http.Error(w, fmt.Sprintf(`{"status": "%s"}`, err), http.StatusInternalServerError)
				return
			}

			mutex.RUnlock()

		case http.MethodPost:
			var c Comment
			if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
				http.Error(w, fmt.Sprintf(`{"status": "%s"}`, err), http.StatusInternalServerError)
				return
			}
			validate := validator.New()
			if err := validate.Struct(c); err != nil {
				http.Error(w, fmt.Sprintf(`{"status": "%s"}`, err), http.StatusBadRequest)
				return
			}
			mutex.Lock()
			comments = append(comments, c)
			mutex.Unlock()

			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"status": "created"}`))

		default:
			http.Error(w, `{"status": "permits only GET or POST"}`, http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8888", nil)
}
