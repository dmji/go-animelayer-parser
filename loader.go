package animelayer

import (
	"bytes"
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func loadHtmlDocument(client *http.Client, urlString string) (*html.Node, error) {

	resp, err := client.Get(urlString)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	fc := doc.FirstChild
	if fc == nil {
		var b bytes.Buffer
		_ = html.Render(&b, fc)
		return nil, fmt.Errorf("unexpected first child: %s", b.String())
	}

	if fc.NextSibling == nil {
		return nil, fmt.Errorf("empty document")
	}

	return doc, nil
}
