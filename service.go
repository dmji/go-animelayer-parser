package animelayer

import (
	"context"
	"net/http"
)

type logger interface {
	Infow(msg string, keys ...interface{})
	Errorw(msg string, keys ...interface{})
}

type Parser interface {
	PipeItemNodesToDetailedItems(ctx context.Context, itemNodes <-chan PageHtmlNode) <-chan ItemDetailedWithError
	PipePageNodesToPartialItems(ctx context.Context, pageNodes <-chan CategoryHtml) <-chan ItemPartialWithError
	PipePagesFromCategoryToPageNode(ctx context.Context, category Category, pages ...int) <-chan CategoryHtml
	PipePartialItemToItemNode(ctx context.Context, partialItems <-chan ItemPartialWithError) <-chan PageHtmlNode

	PartialItemToDetailedItem(ctx context.Context, partialItem ItemPartial) *ItemDetailed
	CategoryPageToPartialItems(ctx context.Context, category Category, iPage int) ([]ItemPartialWithError, error)
}

type service struct {
	client *http.Client
}

func New(client *http.Client) *service {
	return &service{
		client: client,
	}

}

func (p *service) partialItemToItemNode(item ItemPartial) *PageHtmlNode {
	url := formatUrlToItem(item.Identifier)
	doc, err := loadHtmlDocument(p.client, url)
	return &PageHtmlNode{Node: doc, Identifier: item.Identifier, Error: err}
}

func (p *service) pageTargetToHtmlNode(category Category, iPage int) (*CategoryHtml, error) {
	url := formatUrlToItemsPage(category, iPage)

	doc, err := loadHtmlDocument(p.client, url)
	if err != nil {
		return nil, err
	}

	return &CategoryHtml{doc}, nil
}
