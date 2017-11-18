package main

import (
	"github.com/romana/rlog"

	"github.com/asdf/fill_g/fill"
	"github.com/asdf/fill_g/node"
	"github.com/asdf/fill_g/plate"

	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type Step struct {
	Point node.Node
	Color int
}

type Request struct {
	Steps []Step
	Input_data plate.Plate
	Expected_data plate.Plate
}

func main() {
	var file = flag.String("infile", "", "input file")
	var dir = flag.String("indir", "", "input file")
	var remove = flag.Bool("remove", false, "remove input file on success")
	flag.Parse()
	run(*file, *dir, *remove)
}

func run(file string, dir string, remove bool) {
	files := make([]string, 0)
	if file != "" {
		files = append(files, file)
	} else if dir != "" {
		files = get_files_in_dir(dir)
	} else {
		log.Printf("No parameters given")
		return
	}
	process_files(files, remove)
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
		fname := path.Join(dir, file.Name())
		names = append(names, fname)
	}
	return names
}

func process_files(files []string, remove bool) {
	for _, file := range files {
		rlog.Info("process_files, file:", file)
		request, err := read_request(file)
		if err != nil {
			log.Printf("process_files, can't read request for file '%v': %v",
				file, err)
			continue
		}
		result := run_request(request)
		if !plates_equal(result, request.Expected_data) {
			rlog.Warn("process_files, result mismatch, file:", file,
				"\nresult:", result,
				"\nexpected:", request.Expected_data)
			write_result(file, result)
		} else {
			if remove {
				os.Remove(file)
			}
		}
	}
}

func read_request(file string) (Request, error) {
	rlog.Info("read_request, file:", file)
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
	rlog.Info("write_result, file:", file)
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
	plate := request.Input_data
	for _, step := range request.Steps {
		rlog.Debug("run_request, plate:", plate)
		plate = fill.Run(step.Point, step.Color, plate)
	}
	rlog.Debug("run_request, result plate:", plate)
	return plate
}

