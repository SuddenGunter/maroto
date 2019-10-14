package main

import (
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/jung-kurt/gofpdf"
	"log"
)

func main(){
	file := pdf.NewMaroto(gofpdf.OrientationPortrait, consts.A4)
	file.Row(90, func() {
		file.Col(func() {
			file.FileImage("img.png", props.Rect{
				Left:    10,
				Top:     1,
				Center:  false,
				Percent: 30,
			})
		})
	})
	file.Row(90, func() {
		file.Col(func() {
			file.FileImage("img.png", props.Rect{
				Left:    40,
				Top:     1,
				Center:  false,
				Percent: 30,
			})
		})
	})
	file.Row(90, func() {
		file.Col(func() {
			file.FileImage("img.png", props.Rect{
				Left:    50,
				Top:     111,
				Center:  false,
				Percent: 30,
			})
		})
	})

	err := file.OutputFileAndClose("demo.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
