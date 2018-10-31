// Simple HTTP server which counts requests to itself

package main

import (
	"encoding/binary"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Counter struct {
	value     uint64
	countLock sync.Mutex
}

// Increment counter
func (cnt *Counter) Inc() {
	cnt.countLock.Lock()
	cnt.value++
	cnt.countLock.Unlock()
}

// Get the value of counter
func (cnt *Counter) Value() uint64 {
	cnt.countLock.Lock()
	defer cnt.countLock.Unlock()
	return cnt.value
}

type Storage struct {
	f *os.File
}

// Initialize storage
func (s *Storage) Init() {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	s.f = f
}

// Convert counter value to bytes, save it and close file
func (s *Storage) SaveCounterAndClose(c Counter) {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, c.Value())
	s.f.Seek(0, io.SeekStart)
	s.f.Write(b)
	s.f.Close()
}

func (s *Storage) LoadCounter(c Counter) {
	var counter [8]byte
	_, err := io.ReadFull(s.f, counter[:])
	if err != nil {
		if err == io.EOF {
			cnt.value = 0
		} else {
			panic(err)
		}
	} else {
		cnt.value = uint64(binary.LittleEndian.Uint64(counter[:]))
	}
	log.Printf("Counter value is %v\n", cnt.value)
}

// catch SIGINT/SIGTERN and call cleanup()
func configureSignals() {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		cleanup()
		os.Exit(1)
	}()
}

// save counter on interruption
func cleanup() {
	log.Printf("Saving counter value: %v", cnt.Value())
	storage.SaveCounterAndClose(cnt)
}

// handler for all http requests
func handler(w http.ResponseWriter, r *http.Request) {
	cnt.Inc()
	time.Sleep(time.Duration(rand.Intn(1000))) // load imitation
	log.Println(cnt.Value())
}

var cnt Counter
var storage Storage

const filename = "/tmp/counter.dat"

func init() {
	rand.Seed(time.Now().UnixNano())
	configureSignals()
	storage.Init()
	storage.LoadCounter(cnt)
	http.HandleFunc("/", handler)
}

func main() {
	log.Print("Ready to accept connections...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
