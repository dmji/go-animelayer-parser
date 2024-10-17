package animelayer

import (
	"context"
)

func (p *service) GetItemByIdentifier(ctx context.Context, identifier string) (*Item, error) {

	url := formatUrlToItem(identifier)
	doc, err := loadHtmlDocument(p.client, url)
	detailedItemNode := &PageHtmlNode{Node: doc, Identifier: identifier, Error: err}

	return p.parserDetailedItems.ParseItem(ctx, detailedItemNode.Node, detailedItemNode.Identifier)
}
