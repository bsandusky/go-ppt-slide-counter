package main

import (
	"flag"
	"fmt"
	"github.com/bsandusky/slidecounter/utils"
)

func main() {

	output := flag.String("output", "file", "Select output <file|console>. Defaults to file")
	flag.Parse()

	err, files := utils.ParseArgs()
	if err != nil {
		fmt.Println(err)
		return
	}
	err, presentations := utils.CountSlides(files)
	if err != nil {
		fmt.Println(err)
		return
	}

	if *output == "console" {
		utils.OutputToConsole(presentations)
	} else {
		utils.OutputToFile(presentations)
	}
}
