package all

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"sync"
)

var allProjects []string

// Command for scraping all download links, for all mods on CurseForge
// The download links will be saved to files in a project directory created in your current working directory
func GetAllCommand() *cobra.Command {
	// TODO: Make this skip already existing file lists (controllable by boolean cli flag)
	// TODO: Keep state, to make it possible to continue if for some reason stopped (persistent queue e.g. github.com/beeker1121/goque)

	return &cobra.Command{
		Use:   "all",
		Short: "Scrapes all mods from CurseForge",
		Long:  "Creates a directory and scrapes information about all CurseForge mods into it",
		Run: func(cmd *cobra.Command, args []string) {
			getAllMods()
		},
	}
}

type ModFile struct {
	ModName      string `json:"mod"`
	URL          string `json:"url"`
	MCVersion    string `json:"mcVersion"`
	FileName     string `json:"fileName"`
	Size         string `json:"size"`
	ReleasePhase string `json:"releasePhase"`
	UploadTime   string `json:"uploadTime"`
}

type ModList struct {
	mu sync.Mutex
	li []*ModFile
}

func NewModList() *ModList {
	return &ModList{
		mu: sync.Mutex{},
	}
}

func (m *ModList) add(files ...*ModFile) {
	m.mu.Lock()
	m.li = append(m.li, files...)
	m.mu.Unlock()
}

func getAllMods() {
	projPageNum, err := GetPageCount("https://minecraft.curseforge.com/mc-mods")
	if err != nil {
		log.Printf("Error: %s", err)
	}

	projectChan := make(chan string, 128)
	ctx, ctxCancel := context.WithCancel(context.Background())
	waitGroup := sync.WaitGroup{}

	go func() {
		continueLoop := true
		for continueLoop {
			select {
			case <-ctx.Done():
				continueLoop = false
			case project := <-projectChan:
				projectDir := fmt.Sprintf(".%s", project)
				_ = os.MkdirAll(projectDir, 0755)

				var modFiles = NewModList()
				var wg = sync.WaitGroup{}

				modPagenum, _ := GetPageCount(fmt.Sprintf("https://minecraft.curseforge.com%s/files", project))
				if modPagenum > 0 {
					var i int64
					log.Printf("Get mod file list for %s (%d pages)", project, modPagenum)
					for i = 1; i <= modPagenum; i++ {
						wg.Add(1)
						go func() {
							files := GetModFiles(project, fmt.Sprintf("https://minecraft.curseforge.com%s/files?page=%d", project, i))
							modFiles.add(files...)
							wg.Done()
						}()
					}
				} else {
					log.Printf("Get mod file list for %s, (1 page)", project)
					wg.Add(1)
					go func() {
						files := GetModFiles(project, fmt.Sprintf("https://minecraft.curseforge.com%s/files", project))
						modFiles.add(files...)
						wg.Done()
					}()
				}

				wg.Wait()
				modFilesJson, err := json.Marshal(modFiles.li)
				if err != nil {
					log.Printf("Error marshalling mod file list json. err: %s", err)
				}

				err = ioutil.WriteFile(
					path.Join(projectDir, "files.json"),
					modFilesJson,
					0644,
				)
				if err != nil {
					log.Printf("error writing mod file list json file. err: %s", err)
				}

				waitGroup.Done()
			}
		}
	}()

	var i int64
	var wg = sync.WaitGroup{}

	for i = 1; i <= projPageNum; i++ {
		wg.Add(1)
		go func(pagenum int64) {
			r, err := http.DefaultClient.Get(fmt.Sprintf("https://minecraft.curseforge.com/mc-mods?page=%d", pagenum))
			if err != nil {
				log.Printf("Error fetching page %d, Error: %s", i, err)
				return
			}

			doc, err := goquery.NewDocumentFromReader(r.Body)
			if err != nil {
				log.Fatal(err)
			}
			projList := GetProjectList(doc)

			allProjects = append(allProjects, projList...)

			for _, proj := range projList {
				waitGroup.Add(1)
				projectChan <- proj
			}

			log.Printf("%d/%d", pagenum, projPageNum)
		}(i)
	}

	wg.Wait()
	log.Printf("Number of mods: %d", len(allProjects))

	jsonBytes, err := json.Marshal(allProjects)
	if err != nil {
		log.Printf("error marshalling json: %s", err)
	}

	err = ioutil.WriteFile("allprojectlinks.json", jsonBytes, 0644)
	if err != nil {
		log.Printf("error writing json file: %s", err)
	}

	waitGroup.Wait()
	ctxCancel()
	close(projectChan)
}
