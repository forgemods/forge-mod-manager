package install

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/forgemods/forge-mod-manager/command/launcher/install/pack"
	"log"
	"net/http"
	"reflect"
	"runtime/debug"
	"strings"
)

const scrapeBaseUrl = "https://minecraft.curseforge.com"

var printedMods = make(map[string]bool)

type Closeable interface {
	Close() error
}

func Close(c Closeable) {
	err := c.Close()
	if err != nil {
		instanceType := reflect.TypeOf(c)
		log.Printf("WARN: Failed while closing stream of type `%s`, err: %s", instanceType.String(), err)
		debug.PrintStack()
	}
}

func getModDependencies(mod string) []string {
	r, err := http.DefaultClient.Get(fmt.Sprintf("%s/projects/%s/relations/dependencies?filter-related-dependencies=3", scrapeBaseUrl, mod))
	if err != nil {
		log.Fatal(err)
	}

	defer Close(r.Body)

	if r.StatusCode != 200 {
		log.Printf("%s: status: %d", mod, r.StatusCode)
		return nil
	}

	var deps []string

	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".project-listing.project-relations .project-list-item .info.name .name-wrapper a").Each(func(i int, s *goquery.Selection) {
		if href, ok := s.Attr("href"); !ok {
			log.Print("ERR: No href for element")
		} else {
			depModName := strings.Replace(href, fmt.Sprintf("%s/projects/", scrapeBaseUrl), "", -1)

			if _, ok := allModInstances[depModName]; !ok {
				//log.Printf("dep: %s, mod: %s", depModName, mod)
				deps = append(deps, depModName)
				allModInstances[depModName] = &Mod{name: depModName}

				subDeps := getModDependencies(depModName)
				allModInstances[depModName].Dependencies = subDeps
			} else {
				//log.Printf("!! dep: %s, mod: %s", depModName, mod)
				deps = append(deps, depModName)
			}
		}
	})

	return deps
}

const GameVersionTypeID = "2020709689"

var GameVersions = map[string]string{
	"1.12.2": "6756",
}

func documentFromURL(url string) (doc *goquery.Document, err error) {
	r, err := http.DefaultClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	if r.StatusCode != 200 {
		return &goquery.Document{}, fmt.Errorf("unable to get http document from `%s`, err: %s", url, r.Status)
	}

	return goquery.NewDocumentFromReader(r.Body)
}

func getDownloadUrlFromVersion(mod string, modVersion string, mcVersion string) (resultUrl string, err error) {
	doc, err := documentFromURL(fmt.Sprintf(
		"%s/projects/%s/files?filter-game-version=%s:%s",
		scrapeBaseUrl,
		mod,
		GameVersionTypeID,
		GameVersions[mcVersion],
	))

	if err != nil {
		return "", err
	}

	doc.Find("tr.project-file-list-item .project-file-name .project-file-name-container a").Each(func(i int, s *goquery.Selection) {
		if val, ok := s.Attr("data-name"); ok {
			if !strings.Contains(val, modVersion) {
				return
			}

			resultUrl, _ = s.Attr("href")
		}
	})

	if resultUrl == "" {
		err = fmt.Errorf("unable to find a suitable download link for specified mod version")
	}

	resultUrl = fmt.Sprintf("%s%s/download", scrapeBaseUrl, resultUrl)

	return resultUrl, err
}

func getLatestReleaseUrl(mod string, mcVersion string) (resultUrl string, err error) {
	doc, err := documentFromURL(fmt.Sprintf(
		"%s/projects/%s/files?filter-game-version=%s:%s",
		scrapeBaseUrl,
		mod,
		GameVersionTypeID,
		GameVersions[mcVersion],
	))

	if err != nil {
		return "", err
	}

	doc.Find("tr.project-file-list-item:first-child .version-label").Each(func(i int, s *goquery.Selection) {
		ver := strings.Trim(s.Text(), " ")
		if ver != mcVersion {
			log.Fatalf("mod: `%s`, ver: `%s`", mod, ver)
			if loaded, ok := printedMods[mod]; !ok || !loaded {
				log.Printf("%s GAME VERSION: %s", mod, ver)
				printedMods[mod] = true
			}
		}
	})

	doc.Find("tr.project-file-list-item:first-child .project-file-download-button a").Each(func(i int, s *goquery.Selection) {
		if href, ok := s.Attr("href"); !ok {
			log.Print("ERR: No href for element")
		} else {
			resultUrl = fmt.Sprintf("%s%s", scrapeBaseUrl, href)
		}
	})

	return resultUrl, err
}

func GetMod(mod *pack.PackModInfo, outputDir string, mcVersion string) error {
	var downloadUrl string
	var err error

	if mod.Version == "" {
		downloadUrl, err = getLatestReleaseUrl(mod.Mod, mcVersion)
	} else {
		downloadUrl, err = getDownloadUrlFromVersion(mod.Mod, mod.Version, mcVersion)
	}

	if err != nil {
		return fmt.Errorf("unable to get version download url for mod `%s`, err: %s", mod.Mod, err)
	}

	if downloadUrl != "" {
		err := DownloadFile(
			fmt.Sprintf("%s/%s.jar", outputDir, mod.Mod),
			downloadUrl,
		)
		if err != nil {
			return fmt.Errorf("error downloading mod %s,  err: %s", mod, err)
		}
	}

	return nil
}
