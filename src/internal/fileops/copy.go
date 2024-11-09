package fileops

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Copy(src, dest string) error {

	//convert to absolute path
	src, err := filepath.Abs(src)
	if err != nil {
		fmt.Println(err)
		return err
	}
	dest, err = filepath.Abs(dest)
	if err != nil {
		fmt.Println(err)
		return err
	}

	stat, err := os.Stat(src)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if stat.IsDir() {
		err = copyFolder(src, dest)
		if err != nil {
			panic(err)
		}
	}
	err = copyFile(src, dest)
	if err != nil {
		panic(err)
	}
	return nil

}

func copyFolder(src, dest string) error {
	fmt.Printf("copying %s to %s \n", src, dest)
	return nil
}

func copyFile(src, dest string) error {
	fmt.Printf("copying %s to %s\n", src, dest)

	//check destination is a folder, if so create file inside it with same src file name and details
	dstInfo, err := os.Stat(dest)
	if err == nil && dstInfo.IsDir() {
		dest = filepath.Join(dest, filepath.Base(src))
	}

	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Printf("failed to open source file. reason=%s\n", err)
		return err
	}
	defer srcFile.Close()

	// creating a file
	destFile, err := os.Create(dest)
	if err != nil {
		fmt.Printf("failed to create destination file. reason=%s\n", err)
		return err
	}
	defer destFile.Close()

	//copying file content to destination file
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		fmt.Printf("failed to copy file. reason=%s\n", err)
		return err
	}

	//copy file permission
	srcInfo, err := os.Stat(src)
	if err != nil {
		fmt.Printf("failed to stat source file. reason=%s\n", err)
		return err
	}
	os.Chmod(dest, srcInfo.Mode())

	return nil
}
