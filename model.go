package animelayer

import (
	"time"
)

// ItemUpdate - data from update dates block
type ItemUpdate struct {
	UpdatedDate          *time.Time
	CreatedDate          *time.Time
	SeedLastPresenceDate *time.Time

	DebugReadFromElementClass string `json:"ReadFromHtmlKey"`
}

// ItemMetrics - data from metrics block
type ItemMetrics struct {
	Uploads         string
	Downloads       string
	FilesSize       string
	Author          string
	VisitorCounter  string
	ApprovedCounter string

	DebugReadFromElementClass string `json:"ReadFromHtmlKey"`
}

// Item - main item of package
type Item struct {
	Identifier  string
	Title       string
	IsCompleted bool

	Metrics ItemMetrics
	Updated ItemUpdate

	RefImagePreview string
	RefImageCover   string

	Category Category

	Notes           string
	NotesSematizied *NotesSematizied
}

type NotesSematizied struct {
	Taged    []NotesSematiziedItem `json:"Taged,omitempty"`
	Untagged []string              `json:"Untagged,omitempty"`
}

func (n *NotesSematizied) Extend(from *NotesSematizied) {
	for _, t := range from.Taged {
		n.Taged = append(n.Taged, t)
	}
	for _, t := range from.Untagged {
		n.Untagged = append(n.Untagged, t)
	}
}

type NotesSematiziedItem struct {
	Tag    string           `json:"Tag"`
	Text   string           `json:"Text,omitempty"`
	Childs *NotesSematizied `json:"Childs,omitempty"`
}
