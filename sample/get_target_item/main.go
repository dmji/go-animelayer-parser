package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dmji/go-animelayer-parser"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		<-exit
		cancel()
	}()

	p := animelayer.New(animelayer.NewClientWrapper(&http.Client{}), animelayer.WithNoteClassOverride("i", ""))

	id := "60ddae82fd89780039313044"
	firstItem, err := p.GetItemByIdentifier(ctx, id)

	if err != nil {
		panic(err)
	}

	log.Printf("%v", *firstItem)
}
