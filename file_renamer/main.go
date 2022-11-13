package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type fileStruct struct {
	path string
	name string
}

func main() {
	dir := "sample"
	var toRename []fileStruct
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		// dont want to rename directories
		if info.IsDir() {
			return nil
		}
		if _, err := match(info.Name()); err == nil {
			toRename = append(toRename, fileStruct{
				name: info.Name(),
				path: path,
			})
		}
		return nil
	})

	for _, orig := range toRename {
		var n fileStruct
		var err error
		n.name, err = match(orig.name)
		if err != nil {
			fmt.Println("Error matching:", orig.path, err.Error())
		}
		n.path = filepath.Join(dir, n.name)
		fmt.Printf("mv %s => %s\n", orig.path, n.path)
		err = os.Rename(orig.path, n.path)
		if err != nil {
			fmt.Println("Error renaming:", orig.path, err.Error())
		}
	}
}

// match returns the new file name, or an error if the file name
// didn't match our pattern.
func match(filename string) (string, error) {
	pieces := strings.Split(filename, ".")
	ext := pieces[len(pieces)-1]
	tmp := strings.Join(pieces[0:len(pieces)-1], ".")

	pieces = strings.Split(tmp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return "", fmt.Errorf("%s didnt match our pattern", filename)
	}

	return fmt.Sprintf("%s - %d.%s", cases.Title(language.Und, cases.NoLower).String(name), number, ext), nil
}
