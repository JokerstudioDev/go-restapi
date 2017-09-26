package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var tododac = TodoDAC{}
var todoHandler = NewTodoHandler(tododac)

func registerHandlers() {
	route := mux.NewRouter().StrictSlash(true)
	route.HandleFunc("/", index)
	route.HandleFunc("/todos", errorHandler(getAllTodos)).Methods("GET")
	route.HandleFunc("/todos/{id}", errorHandler(getTodo)).Methods("GET")
	http.Handle("/", route)
}

type badRequest struct{ error }
type notFound struct{ error }

func errorHandler(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err == nil {
			return
		}
		switch err.(type) {
		case badRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case notFound:
			http.Error(w, "not found", http.StatusNotFound)
		default:
			log.Println(err)
			http.Error(w, "oops", http.StatusInternalServerError)
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func getAllTodos(w http.ResponseWriter, r *http.Request) error {
	todos, err := todoHandler.GetAllTodos()
	if err != nil {
		return notFound{err}
	}
	return json.NewEncoder(w).Encode(todos)
}

func getTodo(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]
	todo, err := todoHandler.GetTodo(id)
	if err != nil {
		return notFound{err}
	}
	return json.NewEncoder(w).Encode(todo)
}

func main() {
	registerHandlers()
	http.ListenAndServe(":8080", nil)
	log.Println("starting server port 8080")
}
