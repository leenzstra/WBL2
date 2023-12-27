package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/panjf2000/ants/v2"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type wget struct {
	url *url.URL
	dir string
	depth int

	pool    *ants.Pool
	visited map[string]struct{}
}

func NewWget(url *url.URL, dir string, pool *ants.Pool, depth int) *wget {
	return &wget{
		url:  url,
		dir:  dir,
		pool: pool,
		depth: depth,
		visited: make(map[string]struct{}),
	}
}

func (w *wget) RunSync() error {
	return w.run(w.url.String(), w.depth)
}

// func (w *wget) RunConc() error {

// }

func (w *wget) run(root string, depth int) error {
	if depth == 0 {
		// fmt.Println(root, "max depth reached")
		return nil
	}
	
	fmt.Println("Visit", root, "depth", depth)

	p, err := w.download(root)
	if err != nil {
		return err
	}

	url, err := url.Parse(root)
	if err != nil {
		return err
	}

	doc, err := w.parseHtml(bytes.NewReader(p))
	if err != nil {
		return err
	}

	err = w.save(url.Path, bytes.NewReader(p))
	if err != nil {
		return err
	}

	links, err := w.findLinks(doc)
	if err != nil {
		return err
	}

	for _, link := range links {
		if _, ok := w.visited[link]; !ok {
			w.visited[link] = struct{}{}
			if depth > 0 {
				depth--;
			}

			err := w.run(link, depth)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (w *wget) download(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, err
	}

	return io.ReadAll(res.Body)
}

func (w *wget) parseHtml(html io.Reader) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(html)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func (w *wget) save(url string, body io.Reader) error {
	if url == "/" {
		url = "index"
	}

	p := filepath.Join(w.dir, url+".html")
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
        return err
    }

	file, err := os.Create(p)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, body)
	if err != nil {
		return err
	}

	return nil
}

func (w *wget) findLinks(doc *goquery.Document) ([]string, error) {
	links := make([]string, 0)

	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		if strings.HasPrefix(href, "/") {
			u, err := url.Parse(href)
			if err != nil {
				return 
			}
			href, _ = url.JoinPath(w.url.String(), u.Path)
			links = append(links, href)
		}
	})

	return links, nil
}

func main() {
	var (
		// Ссылка на сайт, директория сохранения файов
		siteUrl, dir string
		// Макс. глубина рекурсии
		depth int
	)

	flag.StringVar(&siteUrl, "url", "", "root url")
	flag.StringVar(&dir, "dir", ".", "save dir")
	flag.IntVar(&depth, "depth", -1, "max recurstion depth")

	flag.Parse()

	if siteUrl == "" {
		fmt.Println("no url provided")
		os.Exit(1)
	}

	pool, err := ants.NewPool(25)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	u, err := url.Parse(siteUrl)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	wget := NewWget(u, dir, pool, depth)

	err = wget.RunSync()
	if err != nil {
		fmt.Println(err)
	}
}
