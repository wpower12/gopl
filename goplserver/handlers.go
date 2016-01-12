package goplserver

import (
	"gopl/chp01/lissa"
	"gopl/chp03/mandelbrot"
	"gopl/chp03/surface"
	"log"
	"net/http"
	"strconv"
	"math/rand"
)

func handlerLissa(w http.ResponseWriter, r *http.Request) {
	const fScale = 100.0
	freq := r.URL.Path[len("/lissa/"):]
	f := 0.0
	if freq == "" {
		f = rand.Float64() * 3.0
	} else {
		a, _ := strconv.ParseFloat(freq, 64)
		f = (a/fScale) * 3.0
	}
	lissa.Lissajous(w, f)
	log.Printf("Served Lissa w/ %g\n", f)
}

func handlerSurface(w http.ResponseWriter, r *http.Request) {
	log.Println("Served Surface")
	w.Header().Set("Content-Type", "image/svg+xml")
	surface.Surface(w)
}

func handlerMandelbrot(w http.ResponseWriter, r *http.Request) {
	log.Println("Served Mandelbrot")
	mandelbrot.Mandelbrot(w)
}
