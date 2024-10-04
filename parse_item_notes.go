package animelayer

import (
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func parseIdentifierFromStyleAttr(n *html.Node, prefix string) (string, bool) {
	for _, a := range n.Attr {
		if a.Key == "style" {
			return strings.CutPrefix(a.Val, fmt.Sprintf("view-transition-name: %s-", prefix))
		}
	}
	return "", false
}

func parseItemNotes(n *html.Node, item *ItemDetailed) {

	identifier, bFound := parseIdentifierFromStyleAttr(n, "description")
	if !bFound {
		return
	}

	item.Identifier = identifier

	var b bytes.Buffer

	for c := n.FirstChild; c != nil; c = c.NextSibling {

		err := html.Render(&b, c)
		if err != nil {
			panic(err)
		}

	}

	item.Notes = cleanStringFromHtmlSymbols(b.String())
}
