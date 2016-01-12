package goplserver

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//Set Routing/Handlers
	http.HandleFunc("/lissa/", handlerLissa)
	http.HandleFunc("/surface/", handlerSurface)
	http.HandleFunc("/mandelbrot/", handlerMandelbrot)
	http.HandleFunc("/", handlerIndex)

	//Start the actual server, with logging!
	const address = "localhost:8000"
	log.Printf("Starting Server on %s\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("Served Index")
	msg, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Fprint(w, "Error Loading Page")
	} else {
		fmt.Fprintf(w, "%s", msg)
	}
}
