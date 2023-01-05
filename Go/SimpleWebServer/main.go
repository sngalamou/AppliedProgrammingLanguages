package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)
type User struct {
	FirtsName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
type HomePageHandler struct{}

func (h *HomePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	json.NewDecoder(r.Body).Decode(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	data, _ := json.Marshal(user)
	w.Write(data)
}
func main() {
	http.Handle("/", &HomePageHandler{})

	barHandler := func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "Hello Bar!")
	}
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":3000", nil)
}


