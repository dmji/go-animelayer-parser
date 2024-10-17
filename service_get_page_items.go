package animelayer

import (
	"context"
	"errors"
)

func (p *service) GetItemsFromCategoryPages(ctx context.Context, category Category, iPage int) ([]Item, error) {

	url := formatUrlToItemsPage(category, iPage)

	doc, err := loadHtmlDocument(p.client, url)
	if err != nil {
		return nil, err
	}

	items := make(chan ItemPartialWithError, 100)
	go func() {
		defer close(items)
		p.parserDetailedItems.ParseCategoryPageToChan(ctx, &CategoryHtml{doc}, items)
	}()

	errs := make([]error, 0, 100)
	res := make([]Item, 0, 100)
	for item := range items {
		if item.Error != nil {
			errs = append(errs, item.Error)
			continue
		}

		res = append(res, *item.Item)
	}

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return res, nil
}
