package main

import (
	"testing"
)

func TestFetchArticle(t *testing.T) {
	// expected := []Article{{Title: "test1"}, {Title: "test2"}}

    // ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    //     w.WriteHeader(http.StatusOK)
    //     if err := json.NewEncoder(w).Encode(expected); err != nil {
    //         t.Fatal(err)
    //     }
    // }))
    // defer ts.Close()

	// apiUrl := ts.URL + "/api/v2/items?page=1&per_page=10&query=title:test"

	// articles := FetchArticle("test")
    // if len(articles) != len(expected) {
    //     t.Errorf("unexpected number of articles, got %d, expected %d", len(articles), len(expected))
    // }
}

