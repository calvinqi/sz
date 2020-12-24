package main

import (
	"flag"
	"fmt"

	"github.com/mattn/go-isatty"

	"io/ioutil"
	"os"
	"path/filepath"

	"./helpers"
)

const DirColor = "\033[1;34m%s\033[0m"

type userArgType struct {
	path string
}

// const progName string = "sz"
// TODO get import to support "go build" and put it in a "work space"
// Citation: https://stackoverflow.com/questions/11720079/how-can-i-see-the-size-of-files-and-directories-in-linux
// go get github.com/mattn/go-isatty

// learned from here https://flaviocopes.com/go-list-files/
func getSize(fi os.FileInfo, origPath string) int64 {
	// returns size of the path in bytes
	// file will return the direct size of the file
	// folder will be the recursive size of all its contents
	if !fi.IsDir() {
		return fi.Size()
	}

	fullPath := origPath + "/" + fi.Name()
	var size int64 = 0
	err := filepath.Walk(fullPath, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	if err != nil {
		panic(err)
	}
	return size
}

func processArgs() userArgType {
	// TODO: also process the various other args, perhaps use a struct
	// textPtr := flag.String("a", "", "Include hidden files")
	flag.Parse()
	// fmt.Println("Text obtained is: " + *textPtr)
	nArg := flag.NArg()
	path, _ := os.Getwd() // default value
	// fmt.Println(nArg)
	if nArg > 0 {
		path = flag.Arg(0)
		if nArg != 1 {
			panic(fmt.Sprintf("Program must be invoked with the flags first, e.g. `program -flag1=val1 -flag2=val2 ... arg1`"))
		}
	}

	pathInfo, err := os.Stat(path)
	if os.IsNotExist(err) || !pathInfo.IsDir() {
		panic("Path arg must be a valid directory. Invalid: " + path)
	}

	return userArgType{path: path}
}

func maxLength(arr []string) int {
	result := -1
	for _, str := range arr {
		if len(str) > result {
			result = len(str)
		}
	}
	return result
}

func main() {
	userArgs := processArgs()
	path := userArgs.path
	fileInfos, _ := ioutil.ReadDir(path)

	for _, fileInfo := range fileInfos {
		sizeStr := helpers.ReadableBytes(getSize(fileInfo, path))
		// fmt.Println(fmt.Sprintf("%s %s", sizeStr, fileInfo.Name()))
		pathName := fileInfo.Name()
		if fileInfo.IsDir() {
			// add trailing slash for clarity
			pathName += "/"
			// if directory and valid terminal, colorize the name as well
			if isatty.IsTerminal(os.Stdout.Fd()) {
				pathName = fmt.Sprintf(DirColor, fileInfo.Name()+"/")
			}
		}

		fmt.Printf(fmt.Sprintf("%-4s   %s\n", sizeStr, pathName))
	}
}
