package all

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func GetFileList(project string, doc *goquery.Document) (list []*ModFile) {
	doc.Find(".project-file-listing tr.project-file-list-item").Each(func(i int, s *goquery.Selection) {
		downloadURL, _ := s.Find(".project-file-download-button a").Attr("href")
		mcVersion := strings.Trim(s.Find(".project-file-game-version .version-label").Text(), " \t\n\r")
		fileName := strings.Trim(s.Find(".project-file-name .project-file-name-container").Text(), " \t\n\r")
		fileSize := strings.Trim(s.Find(".project-file-size").Text(), " \t\n\r")
		uploadTime := s.Find(".project-file-date-uploaded abbr.standard-date").AttrOr("data-epoch", "")

		if uploadTime == "" {
			uploadTime = strings.Trim(s.Find(".project-file-date-uploaded").Text(), " \t\n\r")
		}

		var releasePhase = "R"

		if s.Find(".project-file-release-type div").HasClass("alpha-phase") {
			releasePhase = "A"
		} else if s.Find(".project-file-release-type div").HasClass("beta-phase") {
			releasePhase = "B"
		}

		list = append(list, &ModFile{
			ModName:      project,
			URL:          downloadURL,
			MCVersion:    mcVersion,
			FileName:     fileName,
			Size:         fileSize,
			ReleasePhase: releasePhase,
			UploadTime:   uploadTime,
		})
	})

	return list
}

func GetModFiles(project string, url string) (list []*ModFile) {
	r, err := http.DefaultClient.Get(url)
	if err != nil {
		log.Printf("Error fetching url `%s`, Error: %s", url, err)
	}

	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	return GetFileList(project, doc)
}
