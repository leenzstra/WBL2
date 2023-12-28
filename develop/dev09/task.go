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
	"sync"

	"github.com/PuerkitoBio/goquery"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const (
	// Глубина рекурсии/ссылок
	defaultDepth = -1
	// Исходная ссылка
	defaultUrl = ""
	// Путь сохранения
	defaultDir = "."
	// Конкурентный запуск
	defaultConcur = false
)

var (
	verbose = false
)

// Ссылка и ее глубина от исходной
type linkDepth struct {
	url   string
	Depth int
}

type wget struct {
	url   *url.URL
	dir   string
	depth int

	// Посещенные ссылки
	visited map[string]struct{}
}

func NewWget(url *url.URL, dir string, depth int) *wget {
	return &wget{
		url:     url,
		dir:     dir,
		depth:   depth,
		visited: make(map[string]struct{}),
	}
}

// Запуск конкурентного скачивания сайта
func (w *wget) RunAsync() error {
	// канал, в который приходят все ссылки
	linksChan := make(chan linkDepth)
	errors := make(chan error)
	wg := sync.WaitGroup{}

	// Запуск первой ссылки
	wg.Add(1)
	go func() {
		err := w.runAsync(linkDepth{w.url.String(), w.depth}, linksChan, &wg)
		if err != nil {
			errors <- err
		}
	}()

	// Закрытие канала, когда закончили работу
	go func() {
		wg.Wait()
		close(linksChan)
		close(errors)
	}()

	// Получаем ссылки из канала
	// Проверяем были ли уже там
	// Если не были, то отмечаемся, уменьшаем глубину и идем глубже по сайту
	go func() {
		for link := range linksChan {
			link := link
			if _, ok := w.visited[link.url]; !ok {
				w.visited[link.url] = struct{}{}
				if link.Depth > 0 {
					link.Depth = link.Depth - 1
				}
				wg.Add(1)
				go func() {
					err := w.runAsync(link, linksChan, &wg)
					if err != nil {
						errors <- err
					}
				}()
			}
		}
	}()

	return <-errors
}

// Реализация конкурентного запуска
func (w *wget) runAsync(root linkDepth, linksChan chan<- linkDepth, wg *sync.WaitGroup) error {
	defer wg.Done()

	// Получение ссылок из html
	links, err := w.run(root.url, root.Depth)
	if err != nil {
		return err
	}

	// Запись ссылок в общий канал
	for _, link := range links {
		linksChan <- linkDepth{link, root.Depth}
	}

	return nil
}

// Синхронный/последовательный запуск скачивания сайта
func (w *wget) RunSync() error {
	return w.runSync(w.url.String(), w.depth+1)
}

// Реализация синхронного запуска
// Используется рекурсия
func (w *wget) runSync(root string, depth int) error {
	// Получаем ссылки из html
	links, err := w.run(root, depth)
	if err != nil {
		return err
	}

	// Проверяем, были ли уже там
	// Если нет, то идем глубже
	for _, link := range links {
		if _, ok := w.visited[link]; !ok {
			w.visited[link] = struct{}{}

			var err error
			if depth > 0 {
				err = w.runSync(link, depth-1)
			} else {
				err = w.runSync(link, depth)
			}

			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Главная функция обработки html документа
// И получения из него ссылок
func (w *wget) run(root string, depth int) ([]string, error) {
	// Если достигли максимальной глубины, то конец
	if depth == 0 {
		return []string{}, nil
	}

	if verbose {
		fmt.Println("Visit", root, "depth", depth)
	}

	// Получем html
	p, err := w.download(root)
	if err != nil {
		return nil, err
	}

	url, err := url.Parse(root)
	if err != nil {
		return nil, err
	}

	// Парсим html
	doc, err := w.parseHtml(bytes.NewReader(p))
	if err != nil {
		return nil, err
	}

	// Сохранаяем документ
	err = w.save(url.Path, bytes.NewReader(p))
	if err != nil {
		return nil, err
	}

	// Ищем все ссылки
	links, err := w.findLinks(doc)
	if err != nil {
		return nil, err
	}

	return links, nil
}

// Получение html документа
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

// Парсинг html документа
func (w *wget) parseHtml(html io.Reader) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(html)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// Сохранение html документа
func (w *wget) save(url string, body io.Reader) error {
	if strings.HasSuffix(url, "/") {
		url += "index"
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

// Получение ссылок из html
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
		// паралелльная работа
		concur  bool
		verbose bool
	)

	flag.StringVar(&siteUrl, "url", defaultUrl, "root url")
	flag.StringVar(&dir, "dir", defaultUrl, "save dir")
	flag.IntVar(&depth, "depth", defaultDepth, "max recurstion depth")
	flag.BoolVar(&concur, "conc", defaultConcur, "run concurrent")
	flag.BoolVar(&verbose, "v", verbose, "verbose")

	flag.Parse()

	if siteUrl == "" {
		fmt.Println("no url provided")
		os.Exit(1)
	}

	u, err := url.Parse(siteUrl)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	wget := NewWget(u, dir, depth)

	if concur {
		err = wget.RunAsync()
	} else {
		err = wget.RunSync()
	}

	if err != nil {
		fmt.Println(err)
	}
}
