package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

func (u User) SayHello() string {
	name := u.Name
	return "Привет, " + name
}

func handlerGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Привет")
}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	w.Write([]byte(user.SayHello()))
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", handlerGet)
	http.HandleFunc("/sayHello", handlerPost)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
