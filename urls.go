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
func (categories) AnimeHentai() Category {
	return "/torrents/anime/hentai/"
}
func (categories) Manga() Category {
	return "/torrents/manga/"
}
func (categories) MangaHentai() Category {
	return "/torrents/manga/hentai/"
}
func (categories) Music() Category {
	return "/torrents/music/"
}
func (categories) Dorama() Category {
	return "/torrents/dorama/"
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
