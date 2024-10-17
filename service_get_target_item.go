package animelayer

import (
	"context"
)

func (p *service) GetItemByIdentifier(ctx context.Context, identifier string) (*Item, error) {

	url := formatUrlToItem(identifier)
	doc, err := loadHtmlDocument(p.client, url)
	if err != nil {
		return nil, err
	}

	return p.parser.ParseItem(ctx, doc, identifier)
}
