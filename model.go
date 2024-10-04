package animelayer

import "golang.org/x/net/html"

type ItemPartial struct {
	Identifier  string
	Title       string
	IsCompleted bool
}

type ItemDetailed struct {
	Identifier  string
	Title       string
	IsCompleted bool

	TorrentFilesSize string

	RefImagePreview string
	RefImageCover   string

	UpdatedDate string
	CreatedDate string

	Notes string
}

type CategoryHtml struct {
	Node *html.Node
}

type PageHtmlNode struct {
	Node       *html.Node
	Identifier string
	Error      error
}

type ItemPartialWithError struct {
	Item  *ItemPartial
	Error error
}

type ItemDetailedWithError struct {
	Item  *ItemDetailed
	Error error
}
