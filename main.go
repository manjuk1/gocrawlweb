package main

import (
	"fmt"
	"github.com/manjuk1/gocrawlweb/links"
	"log"
	"os"
)

// Calling the URL extactor
func crawl(url string) []string {
	list, err := links.ExtractUrls(url)
	if err != nil {
		errStr := fmt.Errorf("Error Extracting URL %s -> %v", url, err)
		log.Print(errStr)
	}
	return list
}

func main() {
	log.Print("Hello World! Starting a Web crawler project")

	log.Print("Crawling the below links.....")
	for _, link := range os.Args[1:] {
		log.Print(link)
	}

	// Creating Channels to send and receive the urls
	// between main go routine and link Extractor
	urlListCh := make(chan []string)
	unVisitedUrlCh := make(chan string)

	// Push input URL to the channel
	go func() { urlListCh <- os.Args[1:] }()

	// Spawning 20 GO routines to extact the URL's from the given URL
	// Each of these GO routine will wait for UnVisited URL and extracts URL's from within
	go func() {
		for i := 0; i < 20; i++ {
			go func() {
				for link := range unVisitedUrlCh {
					links := crawl(link)
					// If the below code is not wrapped in go routine
					// Leads to deadlock
					go func() { urlListCh <- links }()
				}
			}()
		}
	}()

	// Filtering out the UnVisited URL and pass to Go routine which extracts URLS
	var urlVisits map[string]bool = map[string]bool{}
	for urlLinks := range urlListCh {
		for _, link := range urlLinks {
			if !urlVisits[link] {
				urlVisits[link] = true
				unVisitedUrlCh <- link
			}
		}
	}
}
