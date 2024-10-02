package animelayer

import "golang.org/x/net/html"

type ItemPartial struct {
	Identifier  string
	Title       string
	IsCompleted bool
}

type Note struct {
	Name string
	Text string
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

	LastCheckedDate string

	Notes []Note
}

type PageNode struct {
	Node *html.Node
	Page int
}

type ItemNode struct {
	Node       *html.Node
	Identifier string
}
