package animelayer

import (
	"errors"
	"strconv"
)

// Category - animelayer category
type Category string

// Categories - object to emulate enum class
var Categories = struct {
	Anime       Category
	AnimeHentai Category
	Manga       Category
	MangaHentai Category
	Music       Category
	Dorama      Category
	All         Category
}{
	Anime:       "anime",
	AnimeHentai: "anime_hentai",
	Manga:       "manga",
	MangaHentai: "manga_henai",
	Music:       "music",
	Dorama:      "dorama",
	All:         "",
}

func (c *Category) Presentation() string {
	switch *c {
	case Categories.Anime:
		return "аниме"
	case Categories.AnimeHentai:
		return "аниме"
	case Categories.Manga:
		return "манга"
	case Categories.MangaHentai:
		return "манга"
	case Categories.Music:
		return "музыка"
	case Categories.Dorama:
		return "дорама"
	case Categories.All:
		return ""
	default:
		return ""
	}
}

func (c *Category) Url() string {
	switch *c {
	case Categories.Anime:
		return "/torrents/anime/"
	case Categories.AnimeHentai:
		return "/torrents/anime/hentai/"
	case Categories.Manga:
		return "/torrents/manga/"
	case Categories.MangaHentai:
		return "/torrents/manga/hentai/"
	case Categories.Music:
		return "/torrents/music/"
	case Categories.Dorama:
		return "/torrents/dorama/"
	case Categories.All:
		return ""
	default:
		return ""
	}
}

func CategoryFromString(s string) (Category, error) {
	switch s {

	case string(Categories.Anime):
		return Categories.Anime, nil
	case string(Categories.AnimeHentai):
		return Categories.AnimeHentai, nil
	case string(Categories.Manga):
		return Categories.Manga, nil
	case string(Categories.MangaHentai):
		return Categories.MangaHentai, nil
	case string(Categories.Music):
		return Categories.Music, nil
	case string(Categories.Dorama):
		return Categories.Dorama, nil
	case string(Categories.All):
		return Categories.All, nil
	}

	return Categories.Anime, errors.New("string not match any of categories")
}

func categoryFromPresentationString(s string) (Category, error) {
	switch s {

	case Categories.Anime.Presentation():
		return Categories.Anime, nil
	case Categories.AnimeHentai.Presentation():
		return Categories.AnimeHentai, nil
	case Categories.Manga.Presentation():
		return Categories.Manga, nil
	case Categories.MangaHentai.Presentation():
		return Categories.MangaHentai, nil
	case Categories.Music.Presentation():
		return Categories.Music, nil
	case Categories.Dorama.Presentation():
		return Categories.Dorama, nil
	case Categories.All.Presentation():
		return Categories.All, nil
	}

	return Categories.Anime, errors.New("string not match any of categories")
}

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
