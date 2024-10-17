package animelayer

import (
	"bytes"
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// DocGetter - Interface for html loader
type DocGetter interface {
	Get(utl string) (*html.Node, error)
}

// ClientWrapper - Default wrapper over http.Client
type ClientWrapper struct {
	client *http.Client
}

// Get - method implementation
func (c *ClientWrapper) Get(utl string) (*html.Node, error) {
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

// NewClientWrapper - Create default wrapper over http.Client
func NewClientWrapper(client *http.Client) *ClientWrapper {
	return &ClientWrapper{client}
}

func loadDocument(client DocGetter, urlString string) (*html.Node, error) {

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
