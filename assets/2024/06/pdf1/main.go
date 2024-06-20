// pdf - PDF
//
// Remarks
//   - none
//
// See pdf.txt for program notes.
//
// See CHANGELOG.txt for revision history.
package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf/v2"
)

// imports

// typedefs

// common declarations

func main() {

	// test 1: single page, configuration a4 portrait
	test1()

	// test 2: single page, configuration a4 portrait, security
	test2()

	// test 3: single page, configuration a4 portrait, margins
	test3()

	// test 4: single page, configuration a4 portrait, header, footer, line
	test4()

	// test 5: multiple page, configuration a4 portrait, header, footer, line
	test5()
}

func test1() {

	fmt.Printf("\n---test1---\n")

	// setup
	// font family, style = regular, size
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Arial", "", 16)

	// page
	pdf.AddPage()

	// simple cell
	pdf.Cell(40, 10, "Hello, World 1")
	pdf.Cell(40, 10, "Hello, World 2")
	pdf.Cell(40, 10, "Hello, World 3")

	// multi cell to print text with line breaks
	pdf.MultiCell(0, 10, "", "", "", false)
	pdf.MultiCell(0, 10, "ana voli milovana multi cell to print text with line breaks multi cell to print text with line breaks multi cell to print text with line breaks multi cell to print text with line breaks multi cell to print text with line breaks", "", "C", false)
	// w/ border
	pdf.MultiCell(0, 10, "Hello, World 1", "1", "", false)
	// w/ alignment
	pdf.MultiCell(0, 10, "Hello, World 2", "", "R", false)

	fileName := "doc/test1.pdf"
	err := pdf.OutputFileAndClose(fileName)
	if err != nil {
		fmt.Printf("Cannot output document, err = %v\n", err)
		return
	}

	fmt.Printf("Saved document, file name = %v\n", fileName)
}

func test2() {

	fmt.Printf("\n---test2---\n")

	// setup
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Arial", "", 16)
	pdf.SetProtection(gofpdf.CnProtectPrint, "1234", "abcd")

	// page
	pdf.AddPage()
	pdf.Cell(40, 10, "Hello, World!")

	fileName := "doc/test2.pdf"
	err := pdf.OutputFileAndClose(fileName)
	if err != nil {
		fmt.Printf("Cannot output document, err = %v\n", err)
		return
	}

	fmt.Printf("Saved document, file name = %v\n", fileName)
}

func test3() {

	fmt.Printf("\n---test3---\n")

	// setup
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Arial", "", 16)
	pdf.SetMargins(40, 30, 20)

	// page
	pdf.AddPage()
	pdf.Cell(40, 10, "Hello, World!")

	fileName := "doc/test3.pdf"
	err := pdf.OutputFileAndClose(fileName)
	if err != nil {
		fmt.Printf("Cannot output document, err = %v\n", err)
		return
	}

	fmt.Printf("Saved document, file name = %v\n", fileName)
}

func test4() {

	fmt.Printf("\n---test4---\n")

	// setup
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Arial", "", 16)

	// header
	pdf.SetHeaderFunc(func() {
		pdf.SetFont("Arial", "", 10)
		pdf.Cell(0, 20, "Header 1")
		pdf.Ln(15)
	})
	pdf.SetFooterFunc(func() {
		// 1.5 cm from bottom
		pdf.SetFont("Arial", "", 10)
		pdf.SetY(-15)
		pdf.Cell(0, 10, fmt.Sprintf("Page %d", pdf.PageNo()))
	})

	// page
	pdf.AddPage()
	pdf.Cell(40, 10, "Hello, World!")

	fileName := "doc/test4.pdf"
	err := pdf.OutputFileAndClose(fileName)
	if err != nil {
		fmt.Printf("Cannot output document, err = %v\n", err)
		return
	}

	fmt.Printf("Saved document, file name = %v\n", fileName)
}

func test5() {

	fmt.Printf("\n---test5---\n")

	// setup
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Arial", "", 16)

	// header
	pdf.SetHeaderFunc(func() {
		pdf.SetFont("Arial", "B", 10)
		pdf.Cell(0, 20, "Header 1")
		pdf.Ln(15)
	})
	pdf.SetFooterFunc(func() {
		// 1.5 cm from bottom
		pdf.SetFont("Arial", "I", 10)
		pdf.SetY(-15)
		pdf.Cell(0, 10, fmt.Sprintf("Page %d", pdf.PageNo()))
	})

	// page
	pdf.AddPage()

	// add rows
	for i := 0; i < 100; i++ {
		pdf.MultiCell(0, 5, fmt.Sprintf("%5d", i), "", "", false)
	}

	fileName := "doc/test5.pdf"
	err := pdf.OutputFileAndClose(fileName)
	if err != nil {
		fmt.Printf("Cannot output document, err = %v\n", err)
		return
	}

	fmt.Printf("Saved document, file name = %v\n", fileName)
}
