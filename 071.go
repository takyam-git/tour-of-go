package main

import (
    "fmt"
)

type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string) (body string, urls []string, err error)
}

// map for check fetched
var existsPages map[string]bool

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
    // Run crawl via gorutine
    finish := make(chan bool)
    go crawl(url, depth, fetcher, finish)
    <- finish
}

func crawl(url string, depth int, fetcher Fetcher, finish chan bool){
    // return when fetched or non depth 
    if _, exists := existsPages[url]; exists || depth <= 0 {
        finish <- false
        return
    }
    
    // fetch the page
    body, urls, err := fetcher.Fetch(url)
    
    // add to fetched pages map for check exists
    existsPages[url] = true
    
    // on error
   	if err != nil {
        fmt.Println(err)
        finish <- false
        return
    }
    
    // export result
    fmt.Printf("found: %s %q\n", url, body)
    
    // deps pages countor
    depsPagesCount := 0
    
    // deps pages check channel
    depsPagesFinishCh := make(chan bool)
    
    // do crawl all deps pages
    for _, u := range urls {
        // crawl depth page
        go crawl(u, depth-1, fetcher, depsPagesFinishCh)
        
        // increment deps pages countor for wait
        depsPagesCount++
    }
    
    // wait all deps pages crawling
    for i := 0; i < depsPagesCount; i++ {
        <- depsPagesFinishCh
    }
    
    // finish crawl
	finish <- true
    return
}

func main() {
    existsPages = make(map[string]bool)
    Crawl("http://golang.org/", 4, fetcher)
    fmt.Println("Finish!!")
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
