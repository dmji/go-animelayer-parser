package animelayer

import "strconv"

// Category - animelayer category
type Category string
type categories struct{}

// Categories - object to emulate enum class
var Categories = categories{}

func (categories) Anime() Category {
	return "/torrents/anime/"
}

// AnimeHentai category
//
// Deprecated: The category AnimeHentai no more presented by Animelayer
func (categories) AnimeHentai() Category {
	return "/torrents/anime/hentai/"
}
func (categories) Manga() Category {
	return "/torrents/manga/"
}

// MangaHentai category
//
// Deprecated: The category MangaHentai no more presented by Animelayer
func (categories) MangaHentai() Category {
	return "/torrents/manga/hentai/"
}
func (categories) Music() Category {
	return "/torrents/music/"
}
func (categories) Dorama() Category {
	return "/torrents/dorama/"
}

func (categories) All() Category {
	return ""
}

const baseURL = "https://animelayer.ru"

func formatToItemsPageURL(category Category, iPage int) string {
	return baseURL + string(category) + "?page=" + strconv.FormatInt(int64(iPage), 10)
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
