package repositories

import (
	"encoding/csv"
	"go-qiita/dataformatter"
	"go-qiita/types"
	"log"
	"os"
	"strconv"
	"time"
)

func SaveArticle(articleList []types.Article) {
	
	csvData := dataformatter.ConvertToCsv(articleList)

	var fileName string = "csv/articles/" + strconv.FormatInt(time.Now().Unix(), 10) + "file.csv"

	f, err := os.Create(fileName)
    if err != nil {
        log.Fatal(err)
    }
    w := csv.NewWriter(f)

    w.WriteAll(csvData)

    w.Flush()

    if err := w.Error(); err != nil {
        log.Fatal(err)
    }
}
