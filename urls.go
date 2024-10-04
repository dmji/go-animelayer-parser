package animelayer

import "strconv"

type Category string
type categories struct{}

var Categories = categories{}

func (_ categories) Anime() Category {
	return "/torrents/anime/"
}
func (_ categories) AnimeHentai() Category {
	return "/torrents/anime/hentai/"
}
func (_ categories) Manga() Category {
	return "/torrents/manga/"
}
func (_ categories) MangaHentai() Category {
	return "/torrents/manga/hentai/"
}
func (_ categories) Music() Category {
	return "/torrents/music/"
}
func (_ categories) Dorama() Category {
	return "/torrents/dorama/"
}

const baseUrl = "https://animelayer.ru"

func formatUrlToItemsPage(category Category, iPage int) string {
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
