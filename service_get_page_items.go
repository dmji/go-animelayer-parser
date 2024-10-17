package animelayer

import (
	"context"
)

func (p *service) GetItemsFromCategoryPages(ctx context.Context, category Category, iPage int) ([]Item, error) {

	url := formatUrlToItemsPage(category, iPage)

	doc, err := loadHtmlDocument(p.client, url)
	if err != nil {
		return nil, err
	}

	return p.parser.ParseCategoryPage(ctx, doc)
}
