package fileops

import (
	"fmt"
	"os"
)

func Copy(src, dest string) error {
	stat, err := os.Stat(src)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if stat.IsDir() {
		copyFolder(src, dest)
	}
	copyFile(src, dest)
	return nil

}

func copyFolder(src, dest string) error {
	fmt.Printf("copying %s to %s \n", src, dest)
	return nil
}

func copyFile(src, dest string) error {
	fmt.Printf("copying %s to %s\n", src, dest)
	return nil
}
