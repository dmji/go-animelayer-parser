package animelayer

//go:generate go-stringer -type=Category -trimprefix=Category -nametransform=snake_case_lower -fromstringgenfn -outputtransform=snake_case_lower -marshaljson

import (
	"errors"
	"strconv"
)

// Category - animelayer category
type Category int8

const (
	CategoryAnime Category = iota
	CategoryAnimeHentai
	CategoryManga
	CategoryMangaHentai
	CategoryMusic
	CategoryDorama
	CategoryAll
)

func (c Category) Presentation() string {
	switch c {
	case CategoryAnime:
		return "аниме"
	case CategoryAnimeHentai:
		return "аниме"
	case CategoryManga:
		return "манга"
	case CategoryMangaHentai:
		return "манга"
	case CategoryMusic:
		return "музыка"
	case CategoryDorama:
		return "дорама"
	default:
		return ""
	}
}

func (c *Category) Url() string {
	switch *c {
	case CategoryAnime:
		return "/torrents/anime/"
	case CategoryAnimeHentai:
		return "/torrents/anime/hentai/"
	case CategoryManga:
		return "/torrents/manga/"
	case CategoryMangaHentai:
		return "/torrents/manga/hentai/"
	case CategoryMusic:
		return "/torrents/music/"
	case CategoryDorama:
		return "/torrents/dorama/"
	case CategoryAll:
		return "/"
	default:
		return ""
	}
}

func categoryFromPresentationString(s string) (Category, error) {
	switch s {

	case CategoryAnime.Presentation():
		return CategoryAnime, nil
	case CategoryAnimeHentai.Presentation():
		return CategoryAnimeHentai, nil
	case CategoryManga.Presentation():
		return CategoryManga, nil
	case CategoryMangaHentai.Presentation():
		return CategoryMangaHentai, nil
	case CategoryMusic.Presentation():
		return CategoryMusic, nil
	case CategoryDorama.Presentation():
		return CategoryDorama, nil
	}

	return CategoryAnime, errors.New("string not match any of categories")
}

// Url heplers
const baseURL = "https://animelayer.ru"

func formatToItemsPageURL(category Category, iPage int) string {
	return baseURL + category.Url() + "?page=" + strconv.FormatInt(int64(iPage), 10)
}

func formatToItemURL(identifier string) string {
	return baseURL + "/torrent/" + identifier
}

func formatToItemDownloadURL(identifier string) string {
	return baseURL + "/torrent/" + identifier + "/download"
}

// GetTorrentURL - direct download link
func (m *Item) GetTorrentURL() string {
	return formatToItemURL(m.Identifier)
}
