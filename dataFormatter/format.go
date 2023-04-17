package dataformatter

import "go-qiita/types"

func ConvertToCsv(articleList []types.Article) [][]string {

	var csvData [][]string

	// ヘッダーを追加
	csvData = append(csvData, []string{"title", "url"})

	// csvフォーマットに変換
	for _, article := range articleList {
		var currentArticle []string
		title := article.Title
		url := article.Url
		csvData = append(csvData, append(currentArticle, title, url))
	}

	return csvData
}