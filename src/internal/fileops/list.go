package fileops

import (
	"os"
	"path/filepath"
)

func ListFile(dirOnly, fileOnly bool) ([]string, error) {
	var files []string
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if dirOnly && info.IsDir() {
			files = append(files, path)
		} else if fileOnly && !info.IsDir() {
			files = append(files, path)
		} else if !fileOnly && !dirOnly {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
