package all

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
)

func GetPageCount(url string) (int64, error) {
	r, err := http.DefaultClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	if r.StatusCode != 200 {
		log.Printf("status: %d", r.StatusCode)
		return -1, errors.New(r.Status)
	}

	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Get element after dots (last page number)
	lastPage, err := strconv.ParseInt(doc.Find(".b-pagination li.dots").Next().First().Text(), 10, 32)
	if err != nil {
		return -1, err
	}

	return lastPage, nil
}