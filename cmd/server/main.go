package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "at root!")
	fmt.Println("back home!")
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong!")
	fmt.Println("received a ping request!")
}

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/ping", ping)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed!", err)
	}
}
