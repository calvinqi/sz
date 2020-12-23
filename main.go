package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"./helpers"
)

type userArgType struct {
	path string
}

// const progName string = "sz"
// TODO get import to support "go build" and put it in a "work space"

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

func formatLine(sizeStr string, name string, maxLenStr string) string {
	return fmt.Sprintf("%-"+maxLenStr+"s  %s", sizeStr, name)
}

func main() {
	userArgs := processArgs()
	path := userArgs.path
	fileInfos, _ := ioutil.ReadDir(path)
	numFiles := len(fileInfos)
	sizeStrs := make([]string, numFiles) // one size string per file

	for i, fileInfo := range fileInfos {
		sizeStr := helpers.ReadableBytes(getSize(fileInfo, path))
		sizeStrs[i] = sizeStr
		// fmt.Println(fmt.Sprintf("%s %s", sizeStr, fileInfo.Name()))
	}
	maxLenStr := strconv.Itoa(maxLength(sizeStrs))
	for i, fileInfo := range fileInfos {
		fmt.Println(formatLine(sizeStrs[i], fileInfo.Name(), maxLenStr))
	}
}
