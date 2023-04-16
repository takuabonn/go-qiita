package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Article struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

func main() {
	fmt.Printf("welcome to article serch app\n")

	// ユーザーのinput値を保持
	var title string
	fmt.Scan(&title)

	// 格納したタイトルでqiita記事の検索
	articleList := FetchArticle(title)
	fmt.Println(articleList)

	// スプレッドシートに書き込む

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
