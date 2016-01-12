package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) //ch is now a channel of strings.
	//The data passed into it is of type string.

	for _, url := range os.Args[1:] {
		go fetch(url, ch) //The go keyword starts a goroutine on its own thread
		//the ch variable is a channel that information is passed through
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) //'Receives' data from the channel
		//Each time a sub go routine sends something to the channel, its processed
		//here.S
	}
	fmt.Println("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
