package core

import (
	"encoding/binary"
	"io"
	"log"
	"os"
)

const filename = "/tmp/counter.dat"

type Storage struct {
	f   *os.File
	cnt *Counter
}

// Initialize storage
func (s *Storage) Init(cnt *Counter) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	s.f = f
	s.cnt = cnt
	s.LoadCounter()
}

// Convert counter value to bytes, save it and close file
func (s Storage) Cleanup() {
	log.Printf("Saving counter value: %v", s.cnt.Value())
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, s.cnt.Value())
	s.f.Seek(0, io.SeekStart)
	s.f.Write(b)
	s.f.Close()
}

func (s *Storage) LoadCounter() {
	var counter [8]byte
	_, err := io.ReadFull(s.f, counter[:])
	if err != nil {
		if err == io.EOF {
			s.cnt.Set(0)
		} else {
			panic(err)
		}
	} else {
		s.cnt.Set(uint64(binary.LittleEndian.Uint64(counter[:])))
	}
	log.Printf("Counter value is %v\n", s.cnt.Value())
}
