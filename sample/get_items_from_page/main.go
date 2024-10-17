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

	items, err := p.GetItemsFromCategoryPages(ctx, animelayer.Categories.Anime(), 1)
	if err != nil {
		panic(err)
	}

	for i, item := range items {
		log.Printf("%d: %v", i, item)
	}

}
