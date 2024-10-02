package animelayer

import (
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

func (p *parser) parseItemNotes(n *html.Node, item *ItemDetailed) {

	identifier, bFound := parseIdentifierFromStyleAttr(n, "description")
	if !bFound {
		return
	}

	item.Identifier = identifier

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		t, bOk := getFirstChildTextData(c)
		if !bOk {
			continue
		}

		t = cleanStringFromHtmlSymbols(t)
		if len(t) <= 0 {
			continue
		}

		switch c.Data {
		case "u":
		case "strong":
			t, _ = strings.CutSuffix(t, ":")
			item.Notes = append(item.Notes, Note{Name: t})
		default:
			n := len(item.Notes) - 1
			if n < 0 {
				p.logger.Errorw("ParseDescriptionNode", "error", "description bold part not found")
			}

			if len(item.Notes[n].Name) == 0 {
				p.logger.Errorw("ParseDescriptionNode", "error", "description bold part not found")
			}

			value := item.Notes[n].Text
			if len(value) > 0 {
				value += "\n"
			}
			value += t

			item.Notes[n].Text = value
		}
	}
}
