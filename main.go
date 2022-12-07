package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	price := make([]string, 0)
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		}
		fmt.Println()
		for i, link := range links {
			if string(link) == "₸" {
				// if len(links[i-2]) == 3 {
				price = append(price, string(links[i-2])+string(links[i-1])+string(links[i]))

				// }

			}

		}

	}

	for _, v := range price {
		if len(v) == 10 {
			fmt.Println(v)
		}
	}

}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
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

func visit(links []string, n *html.Node) []string {
	data := strings.Split(n.Data, " ")

	for _, v := range data {
		links = append(links, string(v))
		// fmt.Println(string(v), i)
	}

	// if n.Type == html.ElementNode && n.Data == "spasn" {

	// 	for _, a := range n.Attr {

	// 		// if a.Key == "" {
	// 		fmt.Println(a.Key)
	// 		links = append(links, a.Val)
	// 		// }
	// 	}
	// }

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}
