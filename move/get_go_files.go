package move

import (
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func getGoFilesList(path string) []string {
	listFiles := []string{}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}

		if file.IsDir() {
			listFiles = append(listFiles, getGoFilesList(path+"/"+file.Name())...)

		} else if bool, err := regexp.Match(`.go`, []byte(file.Name())); err == nil && bool {
			listFiles = append(listFiles, path+"/"+file.Name())
		}
	}

	return listFiles
}
