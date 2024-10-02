package animelayer

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

type logger interface {
	Infow(msg string, keys ...interface{})
	Errorw(msg string, keys ...interface{})
}

type loggerPlaceholder struct{}

func (l *loggerPlaceholder) Infow(msg string, keys ...interface{})  {}
func (l *loggerPlaceholder) Errorw(msg string, keys ...interface{}) {}

func (s *pipe) SetLogger(l logger) {
	s.logger = l
}

type pipe struct {
	client *http.Client
	logger logger
}

func New(client *http.Client) *pipe {
	return &pipe{
		client: client,
		logger: &loggerPlaceholder{},
	}

}

func (p *pipe) pipeDocumentToDescription(ctx context.Context, documents <-chan PageNode) {

	items := p.PipePageNodesToPartialItems(ctx, documents)
	descriptionNodes := p.PipePartialItemToItemNode(ctx, items)

	for descriptionNode := range descriptionNodes {

		var b bytes.Buffer
		err := html.Render(&b, descriptionNode.Node)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(fmt.Sprintf("items/%s.html", descriptionNode.Identifier), b.Bytes(), 0644)
		if err != nil {
			panic(err)
		}

	}
}

func (p *pipe) StartPipeForAllPages(ctx context.Context, category category) {
	documents := p.PipeAllPagessFromCategoryToPageNode(ctx, category)
	p.pipeDocumentToDescription(ctx, documents)
}

func (p *pipe) StartPipeForTargetPages(ctx context.Context, category category, pages []int) {
	documents := p.PipeTargetPagessFromCategoryToPageNode(ctx, category, pages)
	p.pipeDocumentToDescription(ctx, documents)
}
