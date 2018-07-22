package links

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

// ExtracUrls extracts the URL's from the given URL
// This function could potentially be enhaced to scrap the complete HTML content
func ExtractUrls(url string) ([]string, error) {
	log.Print("Extracting Url's from" + " ->" + url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		if resp != nil && resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil, fmt.Errorf("%d status while getting %s", resp.StatusCode, url)
		}
		return nil, err
	}
	dom, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("Error %v while parsing url %s ", err, url)
	}
	var extractedUrls []string
	// Anonymous function to parse the DOM Element.
	// This also showcases closure concept by having access to extractedUrls in the outer function.
	parseDomElement := func(domEle *html.Node) {
		if domEle.Type == html.ElementNode && domEle.Data == "a" {
			for _, a := range domEle.Attr {
				if a.Key != "href" {
					continue
				}
				url, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				extractedUrls = append(extractedUrls, url.String())
			}
		}
	}
	traverseDomTree(dom, parseDomElement, nil)
	return extractedUrls, nil
}

func traverseDomTree(domEle *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(domEle)
	}
	for c := domEle.FirstChild; c != nil; c = c.NextSibling {
		traverseDomTree(c, pre, post)
	}
	if post != nil {
		post(domEle)
	}

}
