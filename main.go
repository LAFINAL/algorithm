package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type QWER struct {
	Name string `json:"name"`
	CI  string `json:"ci"`
	DI  string `json:"di"`
	CIF string `json:"CI"`
}
func main(){
	fmt.Println("test")
	http.HandleFunc("/hello", asdf)

	http.ListenAndServe(":8080", nil)
}

func asdf(w http.ResponseWriter, r *http.Request){
	erg := new(QWER)
	if err := json.NewDecoder(r.Body).Decode(&erg); err != nil{
		log.Fatal(err)
	}
	fmt.Println(erg)

	w.Write([]byte("Hello World"))
}