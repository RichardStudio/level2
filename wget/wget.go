package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mywget <url>")
		return
	}
	url := os.Args[1]
	err := downloadSite(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func downloadSite(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Создание директории для сохранения сайта
	dir := "./downloaded_site"
	os.MkdirAll(dir, os.ModePerm)

	// Сохранение главной страницы
	mainPage, err := os.Create(path.Join(dir, "index.html"))
	if err != nil {
		return err
	}
	defer mainPage.Close()

	_, err = io.Copy(mainPage, resp.Body)
	if err != nil {
		return err
	}

	// Парсинг HTML для поиска других ресурсов
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	var downloadLinks func(*html.Node)
	downloadLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "link" || n.Data == "script" || n.Data == "img") {
			for _, attr := range n.Attr {
				if attr.Key == "href" || attr.Key == "src" {
					link := attr.Val
					if strings.HasPrefix(link, "http") {
						// Скачиваем ресурс и сохраняем его
						err := downloadResource(link, dir)
						if err != nil {
							fmt.Println("Error downloading", link, ":", err)
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			downloadLinks(c)
		}
	}
	downloadLinks(doc)

	return nil
}

func downloadResource(url, dir string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fileName := path.Base(url)
	file, err := os.Create(path.Join(dir, fileName))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}
