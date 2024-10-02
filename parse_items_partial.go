package animelayer

import (
	"bytes"
	"context"
	"errors"
	"strings"

	"golang.org/x/net/html"
)

func parseTitle(n *html.Node) (string, bool, error) {
	name, bOk := getFirstChildTextData(n)
	if !bOk {
		return "", false, errors.New("failed to get title")
	}

	title, bFound := strings.CutSuffix(cleanStringFromHtmlSymbols(name), " Complete")
	return title, bFound, nil
}

func parseNodeWithTitle(n *html.Node) *ItemPartial {

	identifier, bOk := parseIdentifierFromStyleAttr(n, "title")
	if !bOk {
		return nil
	}

	ref := getFirstChildHrefNode(n)
	if ref == nil {
		return nil
	}

	title, bCompleted, err := parseTitle(ref)
	if err != nil {
		return nil
	}

	return &ItemPartial{
		Identifier:  identifier,
		Title:       title,
		IsCompleted: bCompleted,
	}
}

func (p *parser) ParsePartialItemsToChan(ctx context.Context, n *html.Node, chItems chan<- ItemPartial) {

	// cart title
	if isElementNodeData(n, "h3") {

		if isExistAttrWithTargetKeyValue(n, "class", "h2 m0") {

			item := parseNodeWithTitle(n)
			if item != nil {
				chItems <- (*item)
			} else {
				var b bytes.Buffer
				err := html.Render(&b, n)
				if err == nil {
					p.logger.Errorw("AnimeLayer ParseCategoryToChan | Warning: Got nil item", "node", b.String())
				} else {
					p.logger.Errorw("AnimeLayer ParseCategoryToChan | Warning: Got nil item but error on html.Render: ", "error", err)
				}
			}
			return
		}

	}

	// traverses the HTML of the webpage from the first child node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		select {
		case <-ctx.Done():
			return
		default:
			p.ParsePartialItemsToChan(ctx, c, chItems)
		}
	}
}
