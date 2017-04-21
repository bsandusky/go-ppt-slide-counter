package utils

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Presentation struct {
	Filename string
	Slides   int `xml:"Slides"`
	Words    int `xml:"Words"`
}

func CountSlides(filepaths []string) (error, []Presentation) {

	// empty slice to collect return values
	var presentations []Presentation

	// File or files exist; iterate and unmarshal slide count into &Presentation{} instance
	for _, v := range filepaths {
		p := &Presentation{}
		p.Filename = v
		filereader, err := zip.OpenReader(v)

		// ignores files that are not valid zips; ie older versions of ppt or
		// non-ppt files with ppt or pptx extensions
		if err != nil {
			continue
		}

		for _, file := range filereader.File {
			if file.Name == "docProps/app.xml" {
				file, err := file.Open()
				if err != nil {
					fmt.Println(err)
				}
				body, err := ioutil.ReadAll(file)
				if err != nil {
					fmt.Println(err)
				}

				err = xml.Unmarshal(body, &p)
				if err != nil {
					fmt.Println(err)
				}
				file.Close()
				filereader.Close()
				presentations = append(presentations, *p)
			}
		}
	}
	return nil, presentations
}
