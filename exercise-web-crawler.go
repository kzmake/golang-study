package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher) {
	type Fetched struct {
		v   map[string]bool
		mux sync.Mutex
	}

	fetched := Fetched{v: make(map[string]bool)}

	var crawler func(fetched Fetched, url string, depth int, done chan struct{}, fetcher Fetcher)
	crawler = func(fetched Fetched, url string, depth int, done chan struct{}, fetcher Fetcher) {
		defer close(done)

		fetched.mux.Lock()
		if _, ok := fetched.v[url]; ok {
			fetched.mux.Unlock()
			return
		}
		fetched.v[url] = true
		fetched.mux.Unlock()

		if depth <= 0 {
			return
		}
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("found: %s %q\n", url, body)

		childrenDone := make([](chan struct{}), len(urls))
		for i, url := range urls {
			childrenDone[i] = make(chan struct{}, 0)
			go crawler(fetched, url, depth-1, childrenDone[i], fetcher)
		}
		for _, done := range childrenDone {
			<-done
		}
		return
	}

	done := make(chan struct{}, 0)
	go crawler(fetched, url, depth, done, fetcher)
	<-done
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
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
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
