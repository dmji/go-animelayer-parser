package animelayer

import "strconv"

type category string
type categories struct{}

var Categories = categories{}

func (_ categories) Anime() category {
	return "/torrents/anime/"
}
func (_ categories) AnimeHentai() category {
	return "/torrents/anime/hentai/"
}
func (_ categories) Manga() category {
	return "/torrents/manga/"
}
func (_ categories) MangaHentai() category {
	return "/torrents/manga/hentai/"
}
func (_ categories) Music() category {
	return "/torrents/music/"
}
func (_ categories) Dorama() category {
	return "/torrents/dorama/"
}

const baseUrl = "https://animelayer.ru"

func formatUrlToItemsPage(category category, iPage int) string {
	return baseUrl + string(category) + "?page=" + strconv.FormatInt(int64(iPage), 10)
}

func formatUrlToItem(identifier string) string {
	return baseUrl + "/torrent/" + identifier
}
func formatUrlToItemDownload(identifier string) string {
	return baseUrl + "/torrent/" + identifier + "/download"
}

func (m *ItemPartial) GetTorrentUrl() string {
	return formatUrlToItem(m.Identifier)
}

func (m *ItemDetailed) GetTorrentUrl() string {
	return formatUrlToItem(m.Identifier)
}
