package core

import (
	"os"
	"os/signal"
	"syscall"
)

type Cleaner interface {
	Cleanup()
}

// catch SIGINT/SIGTERN and call cleanup()
func ConfigureSignals(cleaners ...Cleaner) {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		cleanup(cleaners)
		os.Exit(1)
	}()
}

// clean all
func cleanup(cleaners []Cleaner) {
	for _, cleaner := range cleaners {
		cleaner.Cleanup()
	}
}
