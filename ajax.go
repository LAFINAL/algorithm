package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Computer struct {
	CPU string
	RAM int
	Graphic string
}


func TestCache(w http.ResponseWriter, r *http.Request){
	a := "test"
	w.Write([]byte(a))
}

func AjaxTest(w http.ResponseWriter, r *http.Request){
	myCom := new(Computer)
	err := json.NewDecoder(r.Body).Decode(myCom)
	if err != nil{
		log.Fatal("Decode error: ", err)
	}
	bMycom, err := json.Marshal(myCom)
	if err != nil{
		log.Fatal("Marshal error: ", err)
	}
	fmt.Println("요청은 왔냐?")
	http.Redirect(w, r, "/redirectTest",302)
	//w.WriteHeader(302)
	w.Write(bMycom)
	// postman으로 할 경우 성공
}

func RedirectTest(w http.ResponseWriter, r *http.Request){
	fmt.Println("리다이렉트 왔냐?")
	text := "Redirect end successfully"
	status := 200
	w.WriteHeader(status)
	w.Write([]byte(text))
}