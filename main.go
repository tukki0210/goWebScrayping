package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)


// webページからHTMLを取得する。
func getHTML(url string) *goquery.Document {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	// エラーが発生した場合は、空のドキュメントを返す
	if err != nil {
		log.Fatal(err)
	}

	return doc
}

func getTitle(url string) string {
	doc := getHTML(url)
	title := ""
	doc.Find("h1").Each(func(i int, s *goquery.Selection) {
		title = s.Text()
	})
	return title
}

func main(){
	url := "https://blog-tukki.com/"
	title := getTitle(url)
	fmt.Println("Article Title", title)
}