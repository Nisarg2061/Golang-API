package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
  http.HandleFunc("/", home)
  http.HandleFunc("/route", route)

  addr := "localhost:8080"
  log.Printf("Started server at %s", addr)

  err := http.ListenAndServe(addr, nil)

  if err!= nil{
    log.Fatal(err)
  }
}

func message(w http.ResponseWriter, r_string string){
  response := []byte(r_string)
   _, err := w.Write(response)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(response)
}

func home(w http.ResponseWriter, r *http.Request){
  if r.Method != "GET"{
    http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
  }
  message(w, "Hello World!\n")  
}

func route(w http.ResponseWriter, r *http.Request){
  if r.Method != "GET"{
    http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
  }
  message(w, "Hidden route found!\n")
}


