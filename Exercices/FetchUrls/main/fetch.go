package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type SafeVisited struct {
	mu      sync.Mutex
	visited map[string]bool
}

func (s *SafeVisited) Visit(url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.visited[url] {
		return false
	}
	s.visited[url] = true
	return true
}

func Crawl(url string, depth int, fetcher Fetcher, v *SafeVisited, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 {
		return
	}
	if !v.Visit(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, v, wg)
	}
}

func main() {
	v := &SafeVisited{visited: make(map[string]bool)}
	var wg sync.WaitGroup

	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, v, &wg)
	wg.Wait()
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{"https://golang.org/pkg/", "https://golang.org/cmd/"},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{"https://golang.org/", "https://golang.org/pkg/"},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{"https://golang.org/", "https://golang.org/pkg/"},
	},
}
