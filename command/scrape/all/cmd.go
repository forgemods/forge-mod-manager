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

func GetAllCommand() *cobra.Command {
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

func getAllMods() {
	projPageNum, err := GetPageCount("https://minecraft.curseforge.com/mc-mods")
	if err != nil {
		log.Printf("Error: %s", err)
	}

	projectChan := make(chan string)
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

				var modFileList []*ModFile
				modPagenum, _ := GetPageCount(fmt.Sprintf("https://minecraft.curseforge.com%s/files", project))
				if modPagenum > 0 {
					var i int64
					for i = 1; i <= modPagenum; i++ {
						log.Printf("Get mod file list %s, %d/%d", project, i, modPagenum)
						modFiles := GetModFiles(project, fmt.Sprintf("https://minecraft.curseforge.com%s/files?page=%d", project, i))
						modFileList = append(modFileList, modFiles...)
					}
				} else {
					log.Printf("Get mod file list %s, 1/1", project)
					modFiles := GetModFiles(project, fmt.Sprintf("https://minecraft.curseforge.com%s/files", project))
					modFileList = append(modFileList, modFiles...)
				}

				modFilesJson, err := json.Marshal(modFileList)
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
	for i = 1; i <= projPageNum; i++ {
		r, err := http.DefaultClient.Get(fmt.Sprintf("https://minecraft.curseforge.com/mc-mods?page=%d", i))
		if err != nil {
			log.Printf("Error fetching page %d, Error: %s", i, err)
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

		log.Printf("%d/%d", i, projPageNum)
	}

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
