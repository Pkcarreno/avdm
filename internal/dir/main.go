package dir

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func SetExecutablePermissions(location string) error {
	// Check if it's windows
	OS := strings.ToLower(runtime.GOOS)
	if OS == "windows" {
		return nil
	}

	// change permissions of bin files
	err := filepath.Walk(location, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if !info.IsDir() {
			os.Chmod(path, 0777)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func CheckPath(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func RemoveDir(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
