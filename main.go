package main

import (
	"fmt"

	"go-qiita/fetchArticle"
	"go-qiita/repositories"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")


	fmt.Printf("welcome to article serch app\n")

	// ユーザーのinput値を保持
	var title string
	fmt.Scan(&title)

	// 格納したタイトルでqiita記事の検索
	articleList := fetchArticle.Fetch(title)
	fmt.Println(articleList)

	// csvに書き込む
	repositories.SaveArticle(articleList)

}
