package models

type PageLimitParams struct {
	Offset int    `json:"offset"` // data start with, default is 0
	Limit  int    `json:"limit"`  // setup return data items, default 30
	Total  string `json:"total"`  // unknow, default is `true`
}
