package main

import (
	"log"
	"net/http"
	"os"
	"./handlers"
	"time"
)

func main(){

	l :=log.New(os.Stdout,"product-api",log.LstdFlags)

	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	server := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}
	//http.HandleFunc("/", func(rw http.ResponseWriter,r *http.Request) {
	//	log.Println("Hello World")
	//	d, err := ioutil.ReadAll(r.Body)
	//	if err != nil {
	//		http.Error(rw,"oops",http.StatusBadRequest)
	//		//rw.WriteHeader(http.StatusBadRequest)
	//		//rw.Write([]byte("Ooops"))
	//		return
	//	}
	//	fmt.Fprintf(rw,"Heloo %s",d)
	//	log.Printf("data is %s",d)
	//})
	//fmt.Println("welcome")
	server.ListenAndServe()
}