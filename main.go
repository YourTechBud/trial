package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleGreeting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"greeting": fmt.Sprintf("Hello %s", name)})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/greeting/{name}", HandleGreeting)

	fmt.Println("Starting greeter server")
	http.ListenAndServe(":8080", r)
}
