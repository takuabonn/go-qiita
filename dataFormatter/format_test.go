package dataformatter

import (
	"go-qiita/types"
	"testing"
)

func TestConvertToArticleCsv(t *testing.T) {
    var data []types.Article
    test1 := types.Article{Title: "test 1", Url: "https://test.com/1"}
    test2 := types.Article{Title: "test 2", Url: "https://test.com/2"}
    data = append(data, test1, test2)
    csvData := ConvertToCsv(data)
    header := csvData[0]

    if header[0] != "title" {
        t.Error("headerがありません")
    }

    if len(csvData) == 1 {
        t.Error("変換データがありません")
    }
}