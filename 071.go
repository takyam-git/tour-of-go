package main

import (
    "fmt"
)

type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string) (body string, urls []string, err error)
}

type Page struct {
    Body string
    Urls []string
    Err error
}

var existsPages map[string]bool

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, fin chan int) {
    // TODO: Fetch URLs in parallel.
    // TODO: Don't fetch the same URL twice.
    // This implementation doesn't do either:
    if _, exists := existsPages[url]; exists || depth <= 0 {
        fin <- 1
        return
    }
    
    
    body, urls, err := fetcher.Fetch(url)
    existsPages[url] = true
    
   	if err != nil {
        fmt.Println(err)
        fin <- 1
        return
    }
    
    fmt.Printf("found: %s %q\n", url, body)
    depsPagesCount := 0
    finCh := make(chan int)
    for _, u := range urls {
        go Crawl(u, depth-1, fetcher, finCh)
        depsPagesCount++
    }
    for i := 0; i < depsPagesCount; i++ {
        <- finCh
    }
	fin <- 1
	
    return
}

func main() {
    existsPages = make(map[string]bool)
    comp := make(chan int)
    go Crawl("http://golang.org/", 4, fetcher, comp)
    <-comp
    fmt.Println("finish!")
}

// fakeFetcher is Fetcher that returns canned results.
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

// fetcher is a populated fakeFetcher.
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
