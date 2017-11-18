# flood fill

## Parameters

	--infile - input file. Mutually exclusive with --indir
    --indir - input dir. Mutually exclusive with --infile
    --remove - remove input file after successful processing

## Run

### run with correct expected data:

	fill_g --infile correct.json

### run with wrong expected data:

	fill_g --infile wrong.json

### process all the files in a directory

	fill_g --indir some_input_dir

