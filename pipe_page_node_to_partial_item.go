package animelayer

import (
	"context"
)

func (p *service) PipePageNodesToPartialItems(ctx context.Context, pageNodes <-chan *CategoryHtml) <-chan ItemPartialWithError {
	return PipeGenericWg(ctx, pageNodes, 100, parseCategoryPageToChan)
}
