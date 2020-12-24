package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	"./helpers"
	"github.com/mattn/go-isatty"
	flag "github.com/spf13/pflag"
)

// DirColor commento
const DirColor = "\033[1;34m%s\033[0m"

type userArgType struct {
	path    string
	sorted  bool
	grouped bool
}

// FileInfoWithSize
type FileInfoWithSize struct {
	fileInfo os.FileInfo
	size     int64
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
	sortedBoolPtr := flag.BoolP("sort", "s", false, "sort the results by file size")
	groupBoolPtr := flag.BoolP("group", "g", false, "group the results by file extension (WIP)")
	flag.Parse()
	// fmt.Println("Text obtained is: " + *textPtr)
	// fmt.Println("Sorted Bool obtained is:", *sortedBoolPtr)
	// fmt.Println("Group Bool obtained is:", *groupBoolPtr)
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

	return userArgType{path: path, sorted: *sortedBoolPtr, grouped: *groupBoolPtr}
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

func printFileInfo(fileInfo os.FileInfo, size int64, path string) {
	outputIsTerminal := isatty.IsTerminal(os.Stdout.Fd())

	sizeStr := helpers.ReadableBytes(size)
	// fmt.Println(fmt.Sprintf("%s %s", sizeStr, fileInfo.Name()))
	pathName := fileInfo.Name()
	if fileInfo.IsDir() {
		// add trailing slash for clarity
		pathName += "/"
		// if directory and valid terminal, colorize the name as well
		if outputIsTerminal {
			pathName = fmt.Sprintf(DirColor, fileInfo.Name()+"/")
		}
	}

	fmt.Printf(fmt.Sprintf("%-4s   %s\n", sizeStr, pathName))
}

func main() {
	userArgs := processArgs()
	path := userArgs.path
	fileInfos, _ := ioutil.ReadDir(path)

	if userArgs.sorted {
		// in the sorted case, precompute sizes so we can sort, then print all at once
		fileInfosWithSizes := make([]FileInfoWithSize, len(fileInfos))
		for i, fileInfo := range fileInfos {
			fileInfosWithSizes[i] = FileInfoWithSize{fileInfo: fileInfos[i], size: getSize(fileInfo, path)}
		}
		// custom comparator sort in descending order
		// TODO: ascending order as an option as well?
		sort.SliceStable(fileInfosWithSizes, func(i, j int) bool {
			return fileInfosWithSizes[i].size > fileInfosWithSizes[j].size
		})
		for _, fileInfoWithSize := range fileInfosWithSizes {
			fileInfo, size := fileInfoWithSize.fileInfo, fileInfoWithSize.size
			printFileInfo(fileInfo, size, path)
		}
	} else {
		// in the non-sorted, non-grouped case, print as they come
		for _, fileInfo := range fileInfos {
			size := getSize(fileInfo, path)
			printFileInfo(fileInfo, size, path)
		}
	}
}
