package utils

import (
	"fmt"
	"os"
	"time"
)

func OutputToFile(presentations []Presentation) {
	f, err := os.Create("output_" + time.Now().String() + ".csv")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	var sumSlides int
	var sumWords int

	f.WriteString(fmt.Sprintln("Filename,Num Slides,Num Words"))
	for _, p := range presentations {
		f.WriteString(fmt.Sprintf("%s,%d,%d\n", p.Filename, p.Slides, p.Words))
		sumSlides += p.Slides
		sumWords += p.Words
	}

	f.WriteString(fmt.Sprintf("Totals,%d,%d", sumSlides, sumWords))
}

func OutputToConsole(presentations []Presentation) {

	var sumSlides int
	var sumWords int

	for _, p := range presentations {
		fmt.Printf("Filename: %s\t Slides: %d\t Words: %d\n", p.Filename, p.Slides, p.Words)
		sumSlides += p.Slides
		sumWords += p.Words
	}
	fmt.Printf("Total slides in %d file(s): %d, Words: %d\n", len(presentations), sumSlides, sumWords)
}
