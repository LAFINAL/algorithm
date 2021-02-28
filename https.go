package main

import (
	"fmt"
	"log"
	"net/http"
)

const indexHTML = `<html>
<head>
	<title>Hello</title>
	<script src="/2nd"></script>
</head>
<body>
</body>
</html>
`

func HelloServer(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}


func handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got connection: %s", r.Proto)

	if r.URL.Path == "/2nd" {
		log.Println("Handling 2nd")
		w.Write([]byte("Hello Again"))
		return
	}

	// Handle 1st request
	log.Println("Handling 1st")

	pusher, ok := w.(http.Pusher)
	if !ok {
		log.Println("Can't push to client")
	} else {
		err := pusher.Push("/2nd", nil)
		fmt.Println("Push request has been executed.")
		if err != nil {
			log.Printf("Failed push: %v", err)
		}
	}
	//w.Write([]byte("Hello"))
	fmt.Fprintf(w, indexHTML)
}

func secondHandle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got connection: %s", r.Proto)
	log.Println("Handling 2nd")
	fmt.Println("2")
	w.Write([]byte("Hello Again"))
	fmt.Println("4")
}

func main(){

	srv := &http.Server{Addr: ":8080", Handler: http.HandlerFunc(handle)}

	log.Printf("Serving on https://0.0.0.0:8080")
	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
	http.HandleFunc("/hello", handle)
	//http.HandleFunc("/2nd", secondHandle)
	//err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}
}
