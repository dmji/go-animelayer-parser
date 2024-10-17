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

	p := animelayer.New(animelayer.NewHttpClientWrapper(&http.Client{}), animelayer.WithNoteClassOverride("i", ""))

	id := "56c1b194e1cf6851038b493b"
	detailedFirstItem, err := p.GetItemByIdentifier(ctx, id)

	if err != nil {
		panic(err)
	}

	log.Printf("%v", *detailedFirstItem)
}
