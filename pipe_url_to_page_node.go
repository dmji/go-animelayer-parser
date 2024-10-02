package animelayer

import (
	"context"
)

func (p *pipe) loadHtmlToChan(category category, iPage int, pageNodes chan<- PageNode) bool {

	url := formatUrlToItemsPage(category, iPage)
	p.logger.Infow("Started reading form", "url", url)

	doc, err := loadHtmlDocument(p.client, url)
	if err != nil {
		p.logger.Errorw("Error:", err)
		return false
	}

	pageNodes <- PageNode{
		Node: doc,
		Page: iPage,
	}
	return true
}

func (p *pipe) PipeTargetPagessFromCategoryToPageNode(ctx context.Context, category category, pages []int) <-chan PageNode {
	pageNodes := make(chan PageNode, 10)

	go func() {
		defer close(pageNodes)

		for _, i := range pages {
			select {
			case <-ctx.Done():
				return
			default:
				ok := p.loadHtmlToChan(category, i, pageNodes)
				if !ok {
					break
				}
			}
		}
	}()

	return pageNodes
}

func (p *pipe) PipeAllPagessFromCategoryToPageNode(ctx context.Context, category category) <-chan PageNode {
	documents := make(chan PageNode, 10)

	go func() {
		defer close(documents)
		i := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				i++
				ok := p.loadHtmlToChan(category, i, documents)
				if !ok {
					break
				}
			}
		}

	}()

	return documents
}
