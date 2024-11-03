package fileops

import (
	"io/ioutil"
)

func ListFile(dirOnly, fileOnly bool) ([]string, error) {
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
