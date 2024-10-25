package fileops

import (
	"os"
	"path/filepath"
)

func ListFile() ([]string, error) {
	var files []string

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		files = append(files, path)
		//if !info.IsDir() {
		//	files = append(files, path)
		//}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
