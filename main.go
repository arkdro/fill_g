package main

import (
	"github.com/asdf/fill_g/fill"
	"flag"
	"log"
	"io/ioutil"
)

func main() {
	var file = flag.String("infile", "", "input file")
	var dir = flag.String("indir", "", "input file")
	flag.Parse()
	run(*file, *dir)
}

func run(file string, dir string) {
	files := make([]string, 0)
	if file != "" {
		files = append(files, file)
	} else if dir != "" {
		files = get_files_in_dir(dir)
	} else {
		log.Printf("No parameters given")
		return
	}
	process_files(files)
}

func get_files_in_dir(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("get_files_in_dir, read dir error: %v", err)
		files := make([]string, 0)
		return files
	}
	names := make([]string, 0)
	for _, file := range files {
		names = append(names, file.Name())
	}
	return names
}

func process_files(files []string) {
	for _, file := range files {
		log.Printf("process_files, file: %v", file)
	}
}

