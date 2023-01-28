package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
)

var Reset  = "\033[0m"
var Red    = "\033[31m"
var Green  = "\033[32m"
var Yellow = "\033[33m"
var Blue   = "\033[34m"
var Purple = "\033[35m"
var Cyan   = "\033[36m"
var Gray   = "\033[37m"
var White  = "\033[97m"

func main() {
	if runtime.GOOS == "windows" {
		Reset  = ""
		Red    = ""
		Green  = ""
		Yellow = ""
		Blue   = ""
		Purple = ""
		Cyan   = ""
		Gray   = ""
		White  = ""
	}

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}


	dirs := make([]string, 0, len(files))
	for _, f := range files {
		if f.IsDir() {
			if _, err := os.Stat(f.Name() + "/README.md"); err == nil {
				dirs = append(dirs, f.Name())
			}
		}
	}

	log.Println(dirs)

	result := make(map[string]struct{})
	for _, dir := range dirs {
		cmd := exec.Command("go", "test", "./...")
		cmd.Dir = dir
		out, err := cmd.Output()
		log.Println(string(out))
		if err == nil {
			result[dir] = struct{}{}
		} else {
			log.Println(err)
		}
	}

	points := 0
	for idx, dir := range dirs {
		if _, ok := result[dir]; ok {
			points++
			log.Printf("%d. %sTask %s is OK %s\n", idx + 1, Green, dir, Reset)
		} else {
			log.Printf("%d. %sTask %s is FAIL %s\n",idx + 1,  Red, dir, Reset)
		}
	}

	log.Printf("Completed %d/%d\n", points, 20)

	if points != 20 {
		os.Exit(1)
	}
}