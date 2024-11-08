package fileops

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func getFile(dirOnly, fileOnly bool) ([]string, error) {
	var files []string
	entries, err := ioutil.ReadDir(".")
	for _, entry := range entries {
		if dirOnly && entry.IsDir() {
			files = append(files, entry.Name())
		} else if fileOnly && !entry.IsDir() {
			files = append(files, entry.Name())
		} else if !fileOnly && !dirOnly {
			files = append(files, entry.Name())
		}
	}
	if err != nil {
		return nil, err
	}
	return files, nil
}

func ListFiles(dirOnly, fileOnly bool) {
	files, err := getFile(dirOnly, fileOnly)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {

		info, err := os.Stat(file)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		//icon and color for files
		icon := ""
		color := ""
		if info.IsDir() {
			icon = "📁"
			color = "\033[34m"
		} else {
			ext := filepath.Ext(file)
			switch ext {
			case ".tar", ".zip", ".gz":
				icon = "📦"
			case ".pdf":
				icon = "📑"
			case ".jpg", ".jpeg", ".png", ".gif", ".svg":
				icon = "🖼️"
			default:
				icon = "📄"
			}
			color = "\033[32m"
		}
		fmt.Printf("%s%s %s \033[0m \n", color, icon, file)
	}
}
