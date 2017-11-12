package main

import (
	"github.com/asdf/fill_g/fill"
	"github.com/asdf/fill_g/node"
	"github.com/asdf/fill_g/plate"

	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"log"
)

type Request struct {
	Start_point node.Node
	Color int
	Input_data plate.Plate
	Expected_data plate.Plate
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
		result := run_request(request)
		if !plates_equal(result, request.Expected_data) {
			log.Printf("process_files, result mismatch for file: %v", file)
			write_result(file, result)
		}
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
	if !request.Input_data.Valid_data() {
		log.Printf("invalid plate data in file: %v", file)
		return request, errors.New("invalid plate data")
	}
	if !request.Expected_data.Valid_data() {
		log.Printf("invalid expected data in file: %v", file)
		return request, errors.New("invalid expected data")
	}
	return request, nil
}

func plates_equal(result plate.Plate, expected plate.Plate) bool {
	if result.Width != expected.Width {
		return false
	}
	if result.Height != expected.Height {
		return false
	}
	if len(result.Data) != len(expected.Data) {
		return false
	}
	for y, _ := range result.Data {
		if len(result.Data[y]) != len(expected.Data[y]) {
			return false
		}
		for x, _ := range result.Data[y] {
			if result.Data[y][x] != expected.Data[y][x] {
				return false
			}
		}
	}
	return true
}

func write_result(file string, result plate.Plate) {
	data, err := json.Marshal(result)
	if err != nil {
		log.Printf("can't encode result to json for file '%v': %v\n%+v",
			file, err, result)
		return
	}
	fname := file + "-result"
	ioutil.WriteFile(fname, data, 0644)
}

func run_request(request Request) plate.Plate {
	result := fill.Run(request.Start_point, request.Color, request.Input_data)
	return result
}

