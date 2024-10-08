package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/dmji/go-animelayer-parser"
	"golang.org/x/net/html"
)

type loggerBasic struct{}

func (l *loggerBasic) kts(keys ...interface{}) string {
	args := ""
	for _, key := range keys {

		switch key.(type) {

		case string:
			args += " " + key.(string)
		case int:
			args += " " + strconv.FormatInt(int64(key.(int)), 10)
		case []interface{}:
			args += l.kts(key.([]interface{})...)
		}

	}
	return args
}

func (l *loggerBasic) Infow(msg string, keys ...interface{}) {
	log.Print("Info  | ", msg, l.kts(keys))
}
func (l *loggerBasic) Errorw(msg string, keys ...interface{}) {
	log.Print("Error | ", msg, l.kts(keys))
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		sig := <-exit
		log.Printf("Catched signal: %s", sig.String())
		cancel()
	}()

	p := animelayer.New(&http.Client{})

	// get first 3 anime pages
	pageNodes := p.PipePagesFromCategoryToPageNode(ctx, animelayer.Categories.Anime(), 1, 2, 3)

	// intercept page html result to files
	pageNodes2 := animelayer.PipeGenericInterceptor(ctx, pageNodes, 100, func(pageNode **animelayer.CategoryHtml) {

		var b bytes.Buffer
		err := html.Render(&b, (**pageNode).Node)
		if err != nil {
			panic(err)
		}

		os.Mkdir("~category_anime", 0700)
		err = os.WriteFile(fmt.Sprintf("~category_anime/page_%.3d.html", pageNode), b.Bytes(), 0644)
		if err != nil {
			panic(err)
		}

	})

	// parse partial items
	partialItems := p.PipePageNodesToPartialItems(ctx, pageNodes2)

	// got from partial items url to detailed html nodes
	itemNodes := p.PipePartialItemToItemNode(ctx, partialItems)

	// intercept page html result to files
	itemNodes2 := animelayer.PipeGenericInterceptor(ctx, itemNodes, 100, func(itemNode *animelayer.PageHtmlNode) {

		var b bytes.Buffer
		err := html.Render(&b, itemNode.Node)
		if err != nil {
			panic(err)
		}

		os.Mkdir("~item_anime", 0700)
		err = os.WriteFile(fmt.Sprintf("~item_anime/item_%s.html", itemNode.Identifier), b.Bytes(), 0644)
		if err != nil {
			panic(err)
		}

	})

	// get detailed items form item nodes
	detailedItemProps := p.PipeItemNodesToDetailedItems(ctx, itemNodes2)

	for {

		select {
		case <-ctx.Done():
			return
		case prop, bOpen := <-detailedItemProps:

			if !bOpen && len(detailedItemProps) == 0 {
				return
			}

			log.Printf("Title: %s (%s)", prop.Item.Title, prop.Item.Identifier)
		}

	}

}
