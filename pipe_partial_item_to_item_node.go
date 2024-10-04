package animelayer

import (
	"context"
)

func (p *service) PipePartialItemToItemNode(ctx context.Context, partialItems <-chan ItemPartialWithError) <-chan PageHtmlNode {
	return PipeGeneric(ctx, partialItems, 100, func(prop *ItemPartialWithError) *PageHtmlNode {
		return p.partialItemToItemNode(*prop.Item)
	})
}
