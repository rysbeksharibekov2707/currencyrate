package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://mig.kz/"
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("Ошибка при запросе: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Неверный статус-код: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalf("Ошибка при чтении документа: %v", err)
	}

	fmt.Println("Парсинг курса валют(покупка):")
	doc.Find("td.buy.delta-neutral").EachWithBreak(func(i int, s *goquery.Selection) bool {
		if i >= 10 {
			return false
		}
		title := s.Text()
		link, _ := s.Attr("href")
		fmt.Printf("%d. %s (%s)\n", i+1, title, link)
		return true
	})
	fmt.Println("Парсинг курса валют(продажа):")
	doc.Find("td.sell.delta-neutral").EachWithBreak(func(i int, s *goquery.Selection) bool {
		if i >= 10 {
			return false
		}
		title := s.Text()
		link, _ := s.Attr("href")
		fmt.Printf("%d. %s (%s)\n", i+1, title, link)
		return true
	})
}
