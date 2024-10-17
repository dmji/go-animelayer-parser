package animelayer_test

import (
	"errors"

	"github.com/dmji/go-animelayer-parser"
)

type TestGetItemParams struct {
	Identifier string

	NoteElem  string
	NoteClass string
}

type TestParseFirstPageParams struct {
	Category      animelayer.Category
	Page          int
	ExpectedError error

	NoteElem  string
	NoteClass string
}

func testListIdentifiers() []TestGetItemParams {

	return []TestGetItemParams{
		{
			Identifier: "670ad89197553b37b913cea8",
		},
		{
			Identifier: "67018ad1315d323c564c1f02",
		},
		{
			Identifier: "5dbf1df7fd89787d9a2c4344",
		},
		{
			Identifier: "5e0f28ecfd89782ceb562a12",
		},
		{
			Identifier: "5786907de1cf68975f8b4568",
		},
		{
			Identifier: "65966cf587060e43aa1f3366",
		},
		{
			Identifier: "670ad89197553b37b913cea8",
			NoteElem:   "i",
		},
		{
			Identifier: "67018ad1315d323c564c1f02",
			NoteElem:   "i",
		},
		{
			Identifier: "5dbf1df7fd89787d9a2c4344",
			NoteElem:   "i",
		},
		{
			Identifier: "5e0f28ecfd89782ceb562a12",
			NoteElem:   "i",
		},
		{
			Identifier: "5786907de1cf68975f8b4568",
			NoteElem:   "i",
		},
		{
			Identifier: "65966cf587060e43aa1f3366",
			NoteElem:   "i",
		},
		{
			Identifier: "670ad89197553b37b913cea8",
			NoteElem:   "i",
			NoteClass:  "test",
		},
		{
			Identifier: "67018ad1315d323c564c1f02",
			NoteElem:   "i",
			NoteClass:  "test",
		},
		{
			Identifier: "5dbf1df7fd89787d9a2c4344",
			NoteElem:   "i",
			NoteClass:  "test",
		},
		{
			Identifier: "5e0f28ecfd89782ceb562a12",
			NoteElem:   "i",
			NoteClass:  "test",
		},
		{
			Identifier: "5786907de1cf68975f8b4568",
			NoteElem:   "i",
			NoteClass:  "test",
		},
		{
			Identifier: "65966cf587060e43aa1f3366",
			NoteElem:   "i",
			NoteClass:  "test",
		},
	}

}

func testCategoryPages() []TestParseFirstPageParams {

	return []TestParseFirstPageParams{
		{
			Category: animelayer.Categories.Anime(),
			Page:     1,
		},
		{
			Category: animelayer.Categories.Manga(),
			Page:     2,
		},
		{
			Category:  animelayer.Categories.Anime(),
			Page:      1,
			NoteElem:  "i",
			NoteClass: "test",
		},
		{
			Category:  animelayer.Categories.Manga(),
			Page:      2,
			NoteElem:  "i",
			NoteClass: "test",
		},
		{
			Category:      animelayer.Categories.Anime(),
			Page:          1500,
			ExpectedError: errors.New("empty document"),
		},
	}

}
