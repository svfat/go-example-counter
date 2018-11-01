// Simple HTTP server which counts requests to itself

package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/svfat/go-example-counter/app/core"
)

// handler for all http requests
func handler(w http.ResponseWriter, r *http.Request) {
	cnt.Inc()
	time.Sleep(time.Duration(rand.Intn(1000))) // load imitation
	log.Println(cnt.Value())
}

var cnt = new(core.Counter)

func init() {
	rand.Seed(time.Now().UnixNano())
	var storage core.Storage
	storage.Init(cnt)
	core.ConfigureSignals(storage)

	http.HandleFunc("/", handler)
}

func main() {
	log.Print("Ready to accept connections...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
