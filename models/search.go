package models

type Search struct {
	KeyWord string
	Type    int
	Limit   int
	Offset  int
}

const (
	searchUrl = "/weapi/search/get"
)

func SearchSongs(search Search) string {
	// params := utils.TransformStructToStr(search)

	// searchResult, _, _ := utils.NeteaseCloudRequest(searchUrl, params, http.MethodGet)
	// return searchResult.(string)
	return ""
}
