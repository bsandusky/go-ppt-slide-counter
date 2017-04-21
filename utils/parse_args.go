package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ParseArgs() (error, []string) {

	// No file(s) input
	if len(os.Args) == 1 {
		return errors.New("No file(s) or directory selected"), nil
	}

	var filepaths []string

	// filepath.WalkFunc
	traverseDirectory := func(path string, f os.FileInfo, err error) error {
		if strings.HasSuffix(path, "ppt") || strings.HasSuffix(path, "pptx") {
			filepaths = append(filepaths, path)
		}
		return nil
	}

	// Deal with dir(s)
	parseDirectory := func(arg string) {
		err := filepath.Walk(arg, traverseDirectory)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Deal with file(s)
	parseRegular := func(arg string) {
		// Check file extensions for ppt and pptx
		if strings.HasSuffix(arg, "ppt") || strings.HasSuffix(arg, "pptx") {
			// Check that file(s) actually exists
			if _, err := os.Stat(arg); err == nil {
				filepaths = append(filepaths, arg)
			}
		}
	}

	// File(s) or dir(s) input
	for _, arg := range os.Args[1:] {
		input, err := os.Stat(arg)

		// Ignore non files or directories; this ignores command line flags
		if err != nil {
			continue
		}
		switch mode := input.Mode(); {
		case mode.IsDir():
			parseDirectory(arg)
		case mode.IsRegular():
			parseRegular(arg)
		}
	}

	// No file(s) found
	if len(filepaths) == 0 {
		return errors.New("No .ppt or .pptx files found"), nil
	}
	return nil, filepaths
}
