package animelayer

import (
	"context"
)

type logger interface {
	Infow(msg string, keys ...interface{})
	Errorw(msg string, keys ...interface{})
}

type ParserPipeline interface {
	PipeItemNodesToDetailedItems(ctx context.Context, itemNodes <-chan PageHtmlNode) <-chan ItemDetailedWithError
	PipePageNodesToPartialItems(ctx context.Context, pageNodes <-chan *CategoryHtml) <-chan ItemPartialWithError
	PipePagesFromCategoryToPageNode(ctx context.Context, category Category, pages ...int) <-chan *CategoryHtml
	PipePartialItemToItemNode(ctx context.Context, partialItems <-chan ItemPartialWithError) <-chan PageHtmlNode
}

type Parser interface {
	GetItemByIdentifier(ctx context.Context, identifier string) (*Item, error)
	GetItemsFromCategoryPages(ctx context.Context, category Category, iPage int) ([]Item, error)
}

type service struct {
	client              HtmlDocGetter
	parserDetailedItems parserDetailedItems
}

type ServiceOptionsApplier func(s *service)

func WithNoteClassOverride(noteElem, noteClass string) ServiceOptionsApplier {
	return func(s *service) {
		s.parserDetailedItems.NotePlaintTextElementInterceptor = noteElem
		s.parserDetailedItems.NotePlaintTextElementClassInterceptor = noteClass
	}
}

func New(client HtmlDocGetter, ServiceOptionsApplier ...ServiceOptionsApplier) *service {
	s := &service{
		client: client,
	}

	for _, apply := range ServiceOptionsApplier {
		apply(s)
	}

	return s
}

func (p *service) partialItemToItemNode(identifier string) *PageHtmlNode {
	url := formatUrlToItem(identifier)
	doc, err := loadHtmlDocument(p.client, url)
	return &PageHtmlNode{Node: doc, Identifier: identifier, Error: err}
}

func (p *service) pageTargetToHtmlNode(category Category, iPage int) (*CategoryHtml, error) {
	url := formatUrlToItemsPage(category, iPage)

	doc, err := loadHtmlDocument(p.client, url)
	if err != nil {
		return nil, err
	}

	return &CategoryHtml{doc}, nil
}
