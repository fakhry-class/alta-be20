package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type article struct {
	ID      int
	Title   string
	Content string
}

// dummy data
var data = []article{
	article{1, "lorem", "lorem"},
	article{2, "ipsum", "ipsum"},
	article{3, "alterra", "academy"},
}

func articles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var result, err = json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	} else if r.Method == "POST" {
		var response = map[string]any{
			"status":  "succes",
			"message": "success add data",
		}
		var result, err = json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return

	}
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	//define endpoint
	http.HandleFunc("/articles", articles)
	http.HandleFunc("/dashboard", articles)
	http.HandleFunc("/products", articles)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
