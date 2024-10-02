package animelayer

import (
	"context"
	"net/http"
)

type logger interface {
	Infow(msg string, keys ...interface{})
	Errorw(msg string, keys ...interface{})
}

type loggerPlaceholder struct{}

func (l *loggerPlaceholder) Infow(msg string, keys ...interface{})  {}
func (l *loggerPlaceholder) Errorw(msg string, keys ...interface{}) {}

func (s *service) SetLogger(l logger) {
	s.logger = l
}

type service struct {
	client *http.Client
	logger logger
}

func New(client *http.Client) *service {
	return &service{
		client: client,
		logger: &loggerPlaceholder{},
	}

}

type Parser interface {
	PipeItemNodesToDetailedItems(ctx context.Context, itemNodes <-chan ItemNode) <-chan ItemDetailed
	PipePageNodesToPartialItems(ctx context.Context, pageNodes <-chan PageNode) <-chan ItemPartial
	PipePagesTargetFromCategoryToPageNode(ctx context.Context, category category, pages []int) <-chan PageNode
	PipePagesAllFromCategoryToPageNode(ctx context.Context, category category) <-chan PageNode
	PipePartialItemToItemNode(ctx context.Context, partialItems <-chan ItemPartial) <-chan ItemNode

	PartialItemToDetailedItem(ctx context.Context, partialItem ItemPartial) *ItemDetailed
	CategoryPageToPartialItems(ctx context.Context, category category, iPage int) ([]ItemPartial, error)
}
