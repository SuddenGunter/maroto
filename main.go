package main

import (
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/jung-kurt/gofpdf"
	"log"
)

func main() {
	file := pdf.NewMaroto(gofpdf.OrientationPortrait, consts.A4)
	file.Row(60, func() {
		file.Col(func() {
			file.QrCode("QR CODE", props.Rect{
				Left:    10,
				Top:     1,
				Center:  false,
				Percent: 30,
			})
		})
	})
	file.Row(60, func() {
		file.Col(func() {
			file.QrCode("QR CODE", props.Rect{
				Left:    50,
				Top:     1,
				Center:  false,
				Percent: 30,
			})
		})
	})
	file.Row(60, func() {
		file.Col(func() {
			file.QrCode("QR CODE", props.Rect{
				Left:    100,
				Top:     -20,
				Center:  false,
				Percent: 30,
			})
		})
	})
	file.Row(60, func() {
		file.Col(func() {
			file.Barcode("BAR CfODE", props.Barcode{
				Left:    60,
				Top:     1,
				Center:  false,
				Percent: 30,
				Proportion: props.Proportion{4,3},
			})
		})
	})
	file.Row(60, func() {
		file.Col(func() {
			file.Barcode("BAR CODE", props.Barcode{
				Left:    50,
				Top:     1,
				Center:  false,
				Percent: 30,
				Proportion: props.Proportion{16,9},
			})
		})
	})
	file.Row(60, func() {
		file.Col(func() {
			file.Barcode("BAR CODE", props.Barcode{
				Left:    1,
				Top:     1,
				Center:  false,
				Percent: 30,
				Proportion: props.Proportion{16,9},
			})
		})
	})
	err := file.OutputFileAndClose("demo.pdf")
	if err != nil {
		log.Fatal(err)
	}
}