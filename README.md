## sz
`sz` is a command line program for listing the file sizes of all the files+folders in the current directory. Example:
```
github/music-sharing calvinqi$ sz
6.1K   .DS_Store
47M    .git/
218M   client/
27M    node_modules/
695B   package.json
7.8K   server.js
88K    yarn.lock
```

## Usage
* `sz` will activate in the current working directory
* `sz <FOLDER_NAME>` will run sz in the specified folder
* `sz -s` will sort in decreasing order of file size

For example,
```
calvinqi$ sz -s ~/github/music-sharing
218M   client/
47M    .git/
27M    node_modules/
88K    yarn.lock
7.8K   server.js
6.1K   .DS_Store
695B   package.json
```


## Purpose
* Understand the disk space footprint for all your directories
* Figure out which large files/folders to delete when freeing up space (lots of bloat shows up, often hiding in packages+caches)
* Existing options didn't seem sufficient (`ls -l` only shows file sizes and not the contents of folders, and for including folders, piping `ls` into `xargs du -sh`, piped again into `sort` was getting cumbersome)

## Installation
Interested? Sweet! Should be a quick installation. `sz` is written with Golang, partially because it was the first example I used to learn the language. Go is cross platform so installation should be easy!
1. Install Go [here](https://golang.org/doc/install) if you don't already have it. Supports Linux+Mac+Windows
2. Clone the repo and build the program
```bash
git clone https://github.com/calvinqi/sz.git
cd sz
go build sz.go
```
3. Copy the executable into a user PATH. I use `/usr/local/bin` on Mac+Linux and it works great
```bash
cp sz /usr/local/bin
```
4. Have fun! You can now use the `sz` command freely when navigating your shell just like you would use `ls` or `cd`.
