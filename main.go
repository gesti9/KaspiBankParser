package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// JoomTovar()
	KaspiTovar()
}

func JoomTovar() {
	url := "https://www.joom.com/ru/products/62f661838ed09b01ebd4e0e2"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error:  %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	linkAll := doc.Find(".content___QyukV").Find(".label___Z2o2Y")
	productName := doc.Find(".card___XVq8N").Find(".name___uxWcB")
	name, _ := productName.Html()

	price, _ := linkAll.Html()
	fmt.Printf("Товар %s сейчас в состоянии %s", name, price)

}
func KaspiTovar() {
	url := "https://kaspi.kz/shop/nur-sultan/c/smartphones/?q=%3Acategory%3ASmartphones%3AmanufacturerName%3AApple%3ASmartphones*Series%3AApple%20iPhone%2014&sort=price-asc"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error:  %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	linkAll := doc.Find(".item-card__debet ").Find(".item-card__prices-price")
	productName := doc.Find(".item-card__name")
	name, _ := productName.Html()

	price, _ := linkAll.Html()
	fmt.Println(name, price)

}
