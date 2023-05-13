package move

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func replaceInFile(listFiles []string, oldText, newText string) {
	for _, pathFile := range listFiles {
		data, err := ioutil.ReadFile(pathFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		lines := strings.Split(string(data), "\n")
		output := []string{}

		inBlockImport := false
		for _, line := range lines {
			if inBlockImport && strings.Contains(line, ")") {
				inBlockImport = false
			} else if strings.Contains(line, "import") && strings.Contains(line, "(") {
				inBlockImport = true
			} else if strings.HasPrefix(line, "import") || inBlockImport {
				output = append(output, strings.Replace(line, oldText, newText, -1))
				continue
			}
			output = append(output, line)
		}

		if err = ioutil.WriteFile(pathFile, []byte(strings.Join(output, "\n")), 0666); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
