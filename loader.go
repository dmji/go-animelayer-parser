package animelayer

import (
	"bytes"
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

type HtmlDocGetter interface {
	Get(utl string) (*html.Node, error)
}

type HttpClientWrapper struct {
	client *http.Client
}

func (c *HttpClientWrapper) Get(utl string) (*html.Node, error) {
	resp, err := c.client.Get(utl)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func NewHttpClientWrapper(client *http.Client) *HttpClientWrapper {
	return &HttpClientWrapper{client}
}

func loadHtmlDocument(client HtmlDocGetter, urlString string) (*html.Node, error) {

	doc, err := client.Get(urlString)
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
