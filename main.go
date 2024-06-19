package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/baonguyendinh07/go_fw/cmd"
)

type student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func returnJson(w http.ResponseWriter, r *http.Request) {
	students := []student{
		{Id: 1, Name: "Super Handsome Man"},
		{Id: 2, Name: "So beautiful flower"},
		{Id: 3, Name: "Super Super Hero"},
	}

	studentsJson, err := json.Marshal(students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(studentsJson)
}

func main() {
	cmd.Execute()
	http.HandleFunc("/", returnJson)
	log.Fatal(http.ListenAndServe(":3000", nil)) // http://localhost:3000/
}
