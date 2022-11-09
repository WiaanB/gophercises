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

func main() {
	// fileName := "birthday_001.txt" // => Birthday - 1 of 4.txt
	// newName, err := match(fileName, 4)
	// if err != nil {
	// 	fmt.Println("no match")
	// 	os.Exit(1)
	// }
	// fmt.Println(newName)
	dir := "./sample"
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	count := 0
	var toRename []string
	for _, file := range files {
		if !file.IsDir() {
			_, err := match(file.Name(), 4)
			if err == nil {
				count++
				toRename = append(toRename, file.Name())
			}
		}
	}

	for _, origFileName := range toRename {
		origPath := filepath.Join(dir, origFileName)
		newName, err := match(origFileName, count)
		if err != nil {
			panic(err)
		}
		newPath := filepath.Join(dir, newName)
		fmt.Printf("mv %s => %s\n", origPath, newPath)
		err = os.Rename(origPath, newPath)
		if err != nil {
			panic(err)
		}
	}
}

// match returns the new file name, or an error if the file name
// didn't match our pattern.
func match(filename string, total int) (string, error) {
	pieces := strings.Split(filename, ".")
	ext := pieces[len(pieces)-1]
	tmp := strings.Join(pieces[0:len(pieces)-1], ".")

	pieces = strings.Split(tmp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return "", fmt.Errorf("%s didnt match our pattern", filename)
	}

	return fmt.Sprintf("%s - %d of %d.%s", cases.Title(language.Und, cases.NoLower).String(name), number, total, ext), nil
}
