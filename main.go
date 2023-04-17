package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Article struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

func main() {
	godotenv.Load(".env")


	fmt.Printf("welcome to article serch app\n")

	// ユーザーのinput値を保持
	var title string
	fmt.Scan(&title)

	// 格納したタイトルでqiita記事の検索
	articleList := FetchArticle(title)
	fmt.Println(articleList)

	// csvに書き込む
	SaveArticle(articleList)

}

func FetchArticle(title string) []Article {
	apiUrl := fmt.Sprintf("https://qiita.com/api/v2/items?page=1&per_page=10&query=title:%s", title)
	
	// apiを叩く
	resp, err := http.Get(apiUrl)
	if err != nil {
		// エラー時の処理
		log.Fatal(err)
		fmt.Println(err)
        os.Exit(1)
	}

	defer resp.Body.Close()


	var data []Article

    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        log.Fatal(err)
    }
	
	return data
}

func SaveArticle(articleList []Article) {
	
	csvData := ConvertToCsv(articleList)

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

func ConvertToCsv(articleList []Article) [][]string {

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
