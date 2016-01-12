package main

import (
	"log"
	"net/http"
  "gopl/chp01/lissa"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler( w http.ResponseWriter, r *http.Request ){
  lissa.Lissajous(w)
}
