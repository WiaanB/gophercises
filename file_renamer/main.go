package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var re = regexp.MustCompile("^(.+?) ([0-9]{4}) [(]([0-9]+) of ([0-9]+)[)][.](.+?)$")
var replaceStr = "$2 - $1 - $3 of $4.$5"

func main() {
	walkDir := "sample"
	var toRename []string
	filepath.Walk(walkDir, func(path string, info os.FileInfo, err error) error {
		// dont want to rename directories
		if info.IsDir() {
			return nil
		}
		if _, err := match(info.Name()); err == nil {
			toRename = append(toRename, path)
		}
		return nil
	})

	for _, oldPath := range toRename {
		dir := filepath.Dir(oldPath)
		fileName := filepath.Base(oldPath)
		newFileName, _ := match(fileName)
		newPath := filepath.Join(dir, newFileName)
		fmt.Printf("mv %s => %s\n", oldPath, newPath)

		// err := os.Rename(oldPath, newPath)
		// if err != nil {
		// 	fmt.Println("Error renaming: ", oldPath, newPath, err.Error())
		// }
	}
}

// match returns the new file name, or an error if the file name
// didn't match our pattern.
func match(filename string) (string, error) {
	if !re.MatchString(filename) {
		return "", fmt.Errorf("%s didnt match our pattern", filename)
	}
	return re.ReplaceAllString(filename, replaceStr), nil
}
