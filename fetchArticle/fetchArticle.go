package fetchArticle

import (
	"encoding/json"
	"fmt"
	"go-qiita/types"
	"log"
	"net/http"
	"os"
)


func Fetch(title string) []types.Article {
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


	var data []types.Article

    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        log.Fatal(err)
    }
	
	return data
}