package animelayer

import (
	"context"
)

func (p *service) pageTargetToHtmlNode(category category, iPage int) (*PageNode, error) {
	url := formatUrlToItemsPage(category, iPage)
	p.logger.Infow("Started reading form", "url", url)

	doc, err := loadHtmlDocument(p.client, url)
	if err != nil {
		return nil, err
	}

	return &PageNode{
		Node: doc,
		Page: iPage,
	}, nil
}

func (p *service) loadHtmlToChan(category category, iPage int, pageNodes chan<- PageNode) bool {

	item, err := p.pageTargetToHtmlNode(category, iPage)
	if err != nil {
		p.logger.Errorw("Error:", err)
		return false
	}

	pageNodes <- *item
	return true
}

func (p *service) PipePagesTargetFromCategoryToPageNode(ctx context.Context, category category, pages []int) <-chan PageNode {
	pageNodes := make(chan PageNode, 10)

	go func() {
		defer close(pageNodes)

	loop:
		for _, i := range pages {
			select {
			case <-ctx.Done():
				return
			default:
				ok := p.loadHtmlToChan(category, i, pageNodes)
				if !ok {
					break loop
				}
			}
		}
	}()

	return pageNodes
}

func (p *service) PipePagesAllFromCategoryToPageNode(ctx context.Context, category category) <-chan PageNode {
	documents := make(chan PageNode, 10)

	go func() {
		defer close(documents)
		i := 0

	loop:
		for {
			select {
			case <-ctx.Done():
				return
			default:
				i++
				ok := p.loadHtmlToChan(category, i, documents)
				if !ok {
					break loop
				}
			}
		}

	}()

	return documents
}
