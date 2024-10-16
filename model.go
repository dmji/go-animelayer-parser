package animelayer

import (
	"time"

	"golang.org/x/net/html"
)

type ItemUpdate struct {
	UpdatedDate          *time.Time
	CreatedDate          *time.Time
	SeedLastPresenceDate *time.Time

	ReadFromHtmlKey string
}

type ItemMetrics struct {
	Uploads         string
	Downloads       string
	FilesSize       string
	Author          string
	VisitorCounter  string
	ApprovedCounter string

	ReadFromHtmlKey string
}

type Item struct {
	Identifier  string
	Title       string
	IsCompleted bool

	Metrics ItemMetrics
	Updated ItemUpdate

	RefImagePreview string
	RefImageCover   string

	Notes string
}

type CategoryHtml struct {
	Node *html.Node
}

// ItemPartial for Pipeline
type PageHtmlNode struct {
	Node       *html.Node
	Identifier string
	Error      error
}

// ItemPartial for Pipeline
type ItemPartialWithError struct {
	Item  *Item
	Error error
}

// ItemDetailed for Pipeline
type ItemDetailedWithError struct {
	Item  *Item
	Error error
}
