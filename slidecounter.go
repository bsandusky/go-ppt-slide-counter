package main

import (
	"fmt"
	"github.com/bsandusky/slidecounter/utils"
)

func main() {
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

	utils.OutputToFile(presentations)
}
