package animelayer

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func parseTitle(n *html.Node) (string, bool, error) {
	name, bOk := getFirstChildTextData(n)
	if !bOk {
		return "", false, errors.New("failed to get title")
	}

	title := cleanStringFromHtmlSymbols(name)
	bCompleted := false

	if titleCuted, bFound := strings.CutSuffix(title, " Complete"); bFound {
		title = titleCuted
		bCompleted = true
	} else {
		bFound := strings.Contains(title, ") Complete")
		if bFound {
			strings.ReplaceAll(title, ") Complete", ") ")
			bCompleted = true
		}
	}
	return title, bCompleted, nil
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

func parseCategoryPageToChan(ctx context.Context, n *html.Node, chItems chan<- ItemPartialWithError) {

	// cart title
	if isExistAttrWithTargetKeyValue(n, "h3", "class", "h2 m0") {

		item := parseNodeWithTitle(n)
		if item != nil {
			chItems <- ItemPartialWithError{
				Item:  item,
				Error: nil,
			}
		} else {
			var b bytes.Buffer
			_ = html.Render(&b, n)
			chItems <- ItemPartialWithError{
				Item:  nil,
				Error: fmt.Errorf("got nil item from parse string: %s", b.String()),
			}
		}
		return
	}

	// traverses the HTML of the webpage from the first child node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		select {
		case <-ctx.Done():
			return
		default:
			parseCategoryPageToChan(ctx, c, chItems)
		}
	}
}
