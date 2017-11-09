package main

import (
	"github.com/asdf/fill_g/fill"
	"flag"
)

func main() {
	var file = flag.String("infile", "", "input file")
	var dir = flag.String("indir", "", "input file")
	flag.Parse()
	fill.Run(*file, *dir)
}

