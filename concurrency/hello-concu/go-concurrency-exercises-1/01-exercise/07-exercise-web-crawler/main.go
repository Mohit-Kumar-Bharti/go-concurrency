package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/html"
)

var (
	fetched map[string]bool
	mu      sync.Mutex
)

type Task struct {
	url   string
	depth int
}

type Result struct {
	url   string
	urls  []string
	err   error
	depth int
}

func Crawl(startURL string, depth int) {

	ch := make(chan *Result)

	var wg sync.WaitGroup

	fetch := func(url string, depth int) {
		defer wg.Done()

		urls, err := findLinks(url)

		ch <- &Result{
			url:   url,
			urls:  urls,
			err:   err,
			depth: depth,
		}
	}

	// outstanding tasks count
	pending := 1

	mu.Lock()
	fetched[startURL] = true
	mu.Unlock()

	wg.Add(1)
	go fetch(startURL, depth)

	for pending > 0 {

		res := <-ch
		pending--

		if res.err != nil {
			fmt.Printf("error fetching %s: %v\n", res.url, res.err)
			continue
		}

		fmt.Printf("crawled: %s\n", res.url)

		if res.depth <= 0 {
			continue
		}

		for _, u := range res.urls {

			mu.Lock()

			if fetched[u] {
				mu.Unlock()
				continue
			}

			fetched[u] = true

			mu.Unlock()

			pending++

			wg.Add(1)
			go fetch(u, res.depth-1)
		}
	}

	wg.Wait()
	close(ch)
}

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	fetched = make(map[string]bool)
	now := time.Now()
	Crawl("https://quotes.toscrape.com/", 1)
	fmt.Println("time taken:", time.Since(now))
}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Print("getting error in http call")
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	return visit(nil, doc), nil
}

// visit appends to links each link found in n, and returns the result.
func visit(links []string, n *html.Node) []string {

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Printf("found link: %s\n", a.Val)
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
