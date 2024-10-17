package animelayer

import (
	"context"
	"fmt"
)

func (p *service) PipePagesFromCategoryToPageNode(ctx context.Context, category Category, pages ...int) <-chan *CategoryHtml {

	documents := make(chan *CategoryHtml, 10)

	getPageIndex := func(i int) (int, bool) {
		if len(pages) == 0 {
			return i, false
		}

		if len(pages) <= i {
			return 0, true
		}

		return pages[i], false
	}

	go func() {
		defer close(documents)

		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				return
			default:
			}

			iPage, bBreak := getPageIndex(i)
			if bBreak {
				break
			}

			doc, err := p.pageTargetToHtmlNode(category, iPage)
			if err != nil {
				break
			}

			documents <- doc
		}

	}()

	return documents
}

func (p *service) PipePageNodesToPartialItems(ctx context.Context, pageNodes <-chan *CategoryHtml) <-chan ItemPartialWithError {
	return PipeGenericWg(ctx, pageNodes, 100, p.parserDetailedItems.ParseCategoryPageToChan)
}

func (p *service) PipePartialItemToItemNode(ctx context.Context, partialItems <-chan ItemPartialWithError) <-chan PageHtmlNode {
	return PipeGeneric(ctx, partialItems, 100, func(prop *ItemPartialWithError) *PageHtmlNode {
		return p.partialItemToItemNode(prop.Item.Identifier)
	})
}

func (p *service) PipeItemNodesToDetailedItems(ctx context.Context, itemNodes <-chan PageHtmlNode) <-chan ItemDetailedWithError {
	return PipeGeneric(ctx, itemNodes, 100, func(itemNode *PageHtmlNode) *ItemDetailedWithError {

		item, err := p.parserDetailedItems.ParseItem(ctx, itemNode.Node, itemNode.Identifier)
		if err != nil {
			return &ItemDetailedWithError{
				Item:  nil,
				Error: fmt.Errorf("got nil item from identifier='%s', err='%v'", itemNode.Identifier, err)}
		}

		return &ItemDetailedWithError{
			Item:  item,
			Error: nil,
		}

	})
}
