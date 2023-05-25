package main

import (
	"context"
	"os"
	"os/signal"
)

var host = "0.0.0.0:1000"

func main() {
	_, cansel := context.WithCancel(context.Background())

	go handlerSignal(cansel)

}

func handlerSignal(cancel context.CancelFunc) {
	sigCh := make(chan os.Signal)

	signal.Notify(sigCh, os.Interrupt)

	for {
		sig := <-sigCh
		switch sig {
		case os.Interrupt:
			cancel()
			return
		}
	}
}

//func startserver
