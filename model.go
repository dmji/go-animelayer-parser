package animelayer

import (
	"time"
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
