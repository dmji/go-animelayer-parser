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

	p := animelayer.New(&http.Client{})

	partialItems, err := p.CategoryPageToPartialItems(ctx, animelayer.Categories.Anime(), 1)
	if err != nil {
		panic(err)
	}

	detailedFirstItem := p.PartialItemToDetailedItem(ctx, partialItems[0])
	for i, partialItem := range partialItems {

		log.Printf("%d: %v", i, partialItem)
	}
	log.Printf("%v", detailedFirstItem)
}
