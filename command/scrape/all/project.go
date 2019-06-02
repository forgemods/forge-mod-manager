package all

import (
	"github.com/PuerkitoBio/goquery"
	"log"
)

func GetProjectList(doc *goquery.Document) (list []string) {
	doc.Find(".project-list-item .name-wrapper a").Each(func(i int, s *goquery.Selection) {
		if href, ok := s.Attr("href"); !ok {
			log.Print("ERR: No href for element")
		} else {
			list = append(list, href)
		}
	})

	return list
}