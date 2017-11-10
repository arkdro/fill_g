package main

import (
	"github.com/asdf/fill_g/fill"

	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"log"
)

type Request struct {
	Start_point fill.Node
	Color int
	Input_data fill.Plate
	Expected_data fill.Plate
}

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
		request, err := read_request(file)
		if err != nil {
			continue
		}
		log.Printf("process_files, request: %+v", request)
	}
}

func read_request(file string) (Request, error) {
	var request Request
	bin_data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("can't read file '%v': %v", file, err)
		return request, err
	}
	err = json.Unmarshal(bin_data, &request)
	if err != nil {
		log.Printf("can't parse file '%v': %v", file, err)
		return request, err
	}
	if !valid_data(request.Input_data) {
		log.Printf("invalid plate data in file: %v", file)
		return request, errors.New("invalid plate data")
	}
	if !valid_data(request.Expected_data) {
		log.Printf("invalid expected data in file: %v", file)
		return request, errors.New("invalid expected data")
	}
	return request, nil
}

func valid_data(plate fill.Plate) bool {
	return true
}

