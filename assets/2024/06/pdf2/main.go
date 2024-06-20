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
	"time"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/consts/protection"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/repository"
)

// imports

// typedefs

// common declarations

func main() {

	// test 1: single page, default configuration
	test1()

	// test 2: single page, configuration a4 vertical, security
	test2()

	// test 3: single page, configuration a4 vertical, margins
	test3()

	// test 4: single page, configuration a4 vertical, header, footer, line
	test4()

	// test 5: single page, configuration a4 vertical, header, footer, line, text, qr
	test5()

	// test 6: multi page, configuration a4 vertical, custom font, header, footer, text, metrics
	test6()

	// test 7: multi page, configuration a4 horizontal, fonts, header, footer, text
	test7()

}

func test1() {

	fmt.Printf("\n---test1---\n")

	// setup core
	mrt := maroto.New()
	fmt.Printf("Core = %+v\n", mrt)

	// create document
	document, err := mrt.Generate()
	if err != nil {
		fmt.Printf("Cannot generate document, err = %v\n", err)
		return
	}

	// save document
	fileName := "doc/test1.pdf"
	err = document.Save(fileName)
	if err != nil {
		fmt.Printf("Cannot save document, err = %v\n", err)
		return
	}
	fmt.Printf("Saved document, file name = %v\n", fileName)
}

func test2() {

	fmt.Printf("\n---test2---\n")

	// setup configuration
	builder := config.NewBuilder()
	builder.WithPageSize(pagesize.A4)
	builder.WithOrientation(orientation.Vertical)
	builder.WithAuthor("Branko", true)
	builder.WithSubject("Sales", true)
	builder.WithTitle("Simple report", true)
	builder.WithDebug(true) // debug = true creates frames
	builder.WithProtection(protection.Copy, "root", "1234")
	fmt.Printf("Builder = %+v\n", builder)

	// build configuration
	config := builder.Build()
	fmt.Printf("Config = %+v\n", config)

	// setup core
	mrt := maroto.New(config)
	fmt.Printf("Core = %+v\n", mrt)

	// create document
	document, err := mrt.Generate()
	if err != nil {
		fmt.Printf("Cannot generate document, err = %v\n", err)
		return
	}

	// save document
	fileName := "doc/test2.pdf"
	err = document.Save(fileName)
	if err != nil {
		fmt.Printf("Cannot save document, err = %v\n", err)
		return
	}
	fmt.Printf("Saved document, file name = %v\n", fileName)

}

func test3() {

	fmt.Printf("\n---test3---\n")

	// setup configuration
	builder := config.NewBuilder()
	builder.WithPageSize(pagesize.A4)
	builder.WithOrientation(orientation.Vertical)
	builder.WithAuthor("Branko", true)
	builder.WithSubject("Sales", true)
	builder.WithTitle("Simple report", true)
	builder.WithDebug(true) // debug = true creates frames
	builder.WithLeftMargin(25)
	builder.WithRightMargin(10)
	builder.WithTopMargin(5)
	fmt.Printf("Builder = %+v\n", builder)

	// build configuration
	config := builder.Build()
	fmt.Printf("Config = %+v\n", config)

	// setup core
	mrt := maroto.New(config)
	fmt.Printf("Core = %+v\n", mrt)

	// create document
	document, err := mrt.Generate()
	if err != nil {
		fmt.Printf("Cannot generate document, err = %v\n", err)
		return
	}

	// save document
	fileName := "doc/test3.pdf"
	err = document.Save(fileName)
	if err != nil {
		fmt.Printf("Cannot save document, err = %v\n", err)
		return
	}
	fmt.Printf("Saved document, file name = %v\n", fileName)

}

func test4() {

	fmt.Printf("\n---test4---\n")

	// setup configuration
	pn := props.PageNumber{
		Pattern: "Page {current} of {total}",
		Place:   props.RightBottom,
	}

	builder := config.NewBuilder()
	builder.WithPageSize(pagesize.A4)
	builder.WithOrientation(orientation.Vertical)
	builder.WithPageNumber(pn)
	builder.WithAuthor("Branko", true)
	builder.WithSubject("Sales", true)
	builder.WithTitle("Simple report", true)
	builder.WithDebug(true) // debug = true creates frames
	fmt.Printf("Builder = %+v\n", builder)

	// build configuration
	config := builder.Build()
	fmt.Printf("Config = %+v\n", config)

	// setup core
	mrt := maroto.New(config)
	fmt.Printf("Core = %+v\n", mrt)

	// header
	headerRow1 := text.NewRow(5, "Sales")
	headerRow2 := text.NewRow(5, "Simple report", props.Text{
		Style: fontstyle.Bold,
		Align: align.Center,
		Color: &props.Color{Red: 0, Green: 117, Blue: 191},
	})
	headerRow3 := text.NewRow(5, "Simple report", props.Text{
		Style: fontstyle.Bold,
		Align: align.Right,
		Color: &props.Color{Red: 0, Green: 117, Blue: 191},
	})

	err := mrt.RegisterHeader(headerRow1, headerRow2, headerRow3)
	if err != nil {
		fmt.Printf("Cannot register document header, err = %v\n", err)
		return
	}

	// footer
	footerRow1 := text.NewRow(5, "(c) POINTER d.o.o.", props.Text{
		Size:  5,
		Align: align.Left,
		Color: &props.Color{Red: 0, Green: 117, Blue: 191},
	})

	err = mrt.RegisterFooter(footerRow1)
	if err != nil {
		fmt.Printf("Cannot register document footer, err = %v\n", err)
		return
	}

	// line
	mrt.AddRow(5, line.NewCol(12))
	mrt.AddRow(15, line.NewCol(11))
	mrt.AddRow(5, line.NewCol(12, props.Line{Style: linestyle.Dashed, Color: &props.BlueColor, Thickness: 1}))

	// create document
	document, err := mrt.Generate()
	if err != nil {
		fmt.Printf("Cannot generate document, err = %v\n", err)
		return
	}

	// save document
	fileName := "doc/test4.pdf"
	err = document.Save(fileName)
	if err != nil {
		fmt.Printf("Cannot save document, err = %v\n", err)
		return
	}
	fmt.Printf("Saved document, file name = %v\n", fileName)

}

func test5() {

	fmt.Printf("\n---test5---\n")

	// setup configuration
	pn := props.PageNumber{
		Pattern: "Page {current} of {total}",
		Place:   props.RightBottom,
	}

	builder := config.NewBuilder()
	builder.WithPageSize(pagesize.A4)
	builder.WithOrientation(orientation.Vertical)
	builder.WithPageNumber(pn)
	builder.WithAuthor("Branko", true)
	builder.WithSubject("Sales", true)
	builder.WithTitle("Simple report", true)
	builder.WithDebug(true) // debug = true creates frames
	fmt.Printf("Builder = %+v\n", builder)

	// build configuration
	config := builder.Build()
	fmt.Printf("Config = %+v\n", config)

	// setup core
	mrt := maroto.New(config)
	fmt.Printf("Core = %+v\n", mrt)

	// header
	headerRow1 := text.NewRow(5, "Sales")
	headerRow2 := text.NewRow(5, "Simple report", props.Text{
		Style: fontstyle.Bold,
		Align: align.Center,
		Color: &props.Color{Red: 0, Green: 117, Blue: 191},
	})
	headerRow3 := text.NewRow(5, "Simple report", props.Text{
		Style: fontstyle.Bold,
		Align: align.Right,
		Color: &props.Color{Red: 0, Green: 117, Blue: 191},
	})

	err := mrt.RegisterHeader(headerRow1, headerRow2, headerRow3)
	if err != nil {
		fmt.Printf("Cannot register document header, err = %v\n", err)
		return
	}

	// footer
	footerRow1 := text.NewRow(5, "(c) POINTER d.o.o.", props.Text{
		Size:  5,
		Align: align.Left,
		Color: &props.Color{Red: 0, Green: 117, Blue: 191},
	})

	err = mrt.RegisterFooter(footerRow1)
	if err != nil {
		fmt.Printf("Cannot register document footer, err = %v\n", err)
		return
	}

	// line
	mrt.AddRow(5, line.NewCol(12))
	mrt.AddRow(5, line.NewCol(12, props.Line{Style: linestyle.Dashed, Color: &props.BlueColor, Thickness: 1}))

	// text
	mrt.AddRow(5,
		text.NewCol(4, "Column size 4"),
		text.NewCol(4, "Column size 4"),
		text.NewCol(4, "Column size 4"))

	mrt.AddRow(5,
		text.NewCol(6, "Column size 6"),
		text.NewCol(3, "Column size 3"),
		text.NewCol(3, "Column size 3"))

	data1 := fmt.Sprintf("%20.20s %30.30s", "hello", "world")
	mrt.AddRow(5, text.NewCol(12, data1))

	data2 := "12345678901234567890123456789012345678901234567890123456789012345678901234567890"
	mrt.AddRow(5, text.NewCol(12, data2))

	data3 := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat..."
	mrt.AddRow(25,
		text.NewCol(6, "Column size 6"),
		text.NewCol(6, data3))

	// qr code
	mrt.AddRow(40,
		code.NewQrCol(4, "https://pointer.hr", props.Rect{
			Percent: 50,
		}),
		code.NewQrCol(4, "https://pointer.hr", props.Rect{
			Percent: 75,
		}),
		code.NewQrCol(4, "https://pointer.hr", props.Rect{
			Percent: 100,
		}),
	)
	mrt.AddRow(40,
		code.NewQrCol(4, "https://pointer.hr", props.Rect{
			Left:    5,
			Top:     5,
			Center:  true,
			Percent: 50,
		}),
		code.NewQrCol(4, "https://pointer.hr", props.Rect{
			Left:    5,
			Top:     5,
			Center:  true,
			Percent: 75,
		}),
		code.NewQrCol(4, "https://pointer.hr", props.Rect{
			Left:    5,
			Top:     5,
			Center:  true,
			Percent: 100,
		}),
	)

	// create document
	document, err := mrt.Generate()
	if err != nil {
		fmt.Printf("Cannot generate document, err = %v\n", err)
		return
	}

	// save document
	fileName := "doc/test5.pdf"
	err = document.Save(fileName)
	if err != nil {
		fmt.Printf("Cannot save document, err = %v\n", err)
		return
	}
	fmt.Printf("Saved document, file name = %v\n", fileName)

}

func test6() {

	fmt.Printf("\n---test6---\n")

	const LINE_SIZE = 5

	// setup fonts
	fontFamily := "courier-new"
	fontRepo := repository.New()
	fontRepo.AddUTF8Font(fontFamily, fontstyle.Normal, "./fonts/RobotoMono-Regular.ttf")
	fontRepo.AddUTF8Font(fontFamily, fontstyle.Italic, "./fonts/RobotoMono-Italic.ttf")
	fontRepo.AddUTF8Font(fontFamily, fontstyle.Bold, "./fonts/RobotoMono-Bold.ttf")
	fontRepo.AddUTF8Font(fontFamily, fontstyle.BoldItalic, "./fonts/RobotoMono-BoldItalic.ttf")

	fonts, err := fontRepo.Load()
	if err != nil {
		fmt.Printf("Cannot load fonts, err = %v\n", err)
		return
	}
	fmt.Printf("Fonts = %+v\n", fonts)

	// setup configuration
	pn := props.PageNumber{
		Pattern: "APP 1.2.3 | {current}/{total}",
		Place:   props.RightBottom,
	}

	builder := config.NewBuilder()
	builder.WithPageSize(pagesize.A4)
	builder.WithOrientation(orientation.Vertical)
	builder.WithPageNumber(pn)
	builder.WithAuthor("Branko", true)
	builder.WithCreator("APP 1.2.3", true)
	builder.WithSubject("Sales", true)
	builder.WithTitle("Simple report", true)
	// builder.WithDebug(true) // debug = true creates frames
	builder.WithCustomFonts(fonts)
	builder.WithDefaultFont(&props.Font{Family: fontFamily})
	fmt.Printf("Builder = %+v\n", builder)

	// build configuration
	config := builder.Build()
	fmt.Printf("Config = %+v\n", config)

	// setup core
	mrt := maroto.New(config)
	mrt = maroto.NewMetricsDecorator(mrt)
	fmt.Printf("Core = %+v\n", mrt)

	// header
	// .........1.........2.........3.........4.........5.........6.........7.........8.......89
	// 12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789
	// xxxxxxxxx1: xxxxxxxxx1xxxxxxxxx2 							         yyyy-mm-dd hh:mm:ss
	//
	headerRow1 := row.New(LINE_SIZE)
	headerRow1 = headerRow1.Add(
		text.NewCol(6, fmt.Sprintf("%s: %s ", "Sales", "Simple report")),
		text.NewCol(6, time.Now().Format("2006-01-02 15:04:05"), props.Text{
			Align: align.Right,
		}))
	headerRowDelimiter := text.NewRow(LINE_SIZE, "")
	err = mrt.RegisterHeader(headerRow1, headerRowDelimiter)
	if err != nil {
		fmt.Printf("Cannot register document header, err = %v\n", err)
		return
	}

	// footer
	footerRowDelimiter := text.NewRow(LINE_SIZE, "")
	err = mrt.RegisterFooter(footerRowDelimiter)
	if err != nil {
		fmt.Printf("Cannot register document footer, err = %v\n", err)
		return
	}

	// template row
	mrt.AddRow(LINE_SIZE, text.NewCol(12, ".........1.........2.........3.........4.........5.........6.........7.........8.......89"))
	mrt.AddRow(LINE_SIZE, text.NewCol(12, "12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789"))

	// add rows
	for i := 0; i < 500; i++ {
		mrt.AddRow(LINE_SIZE, text.NewCol(12, fmt.Sprintf("%5d ...1.........2.........3.........4.........5.........6.........7.........8.......89", i)))
	}

	// create document
	document, err := mrt.Generate()
	if err != nil {
		fmt.Printf("Cannot generate document, err = %v\n", err)
		return
	}

	// save document
	fileName := "doc/test6.pdf"
	err = document.Save(fileName)
	if err != nil {
		fmt.Printf("Cannot save document, err = %v\n", err)
		return
	}
	fmt.Printf("Saved document, file name = %v\n", fileName)

	// save report
	fileName = "doc/test6.pdf.txt"
	err = document.GetReport().Save(fileName)
	if err != nil {
		fmt.Printf("Cannot save report, err = %v\n", err)
		return
	}
	fmt.Printf("Saved report, file name = %v\n", fileName)
}

func test7() {

	fmt.Printf("\n---test6---\n")

	const LINE_SIZE = 4

	// setup fonts
	fontFamily := "courier-new"
	fontRepo := repository.New()
	// fontRepo.AddUTF8Font(fontFamily, fontstyle.Normal, "./fonts/RobotoMono-Regular.ttf")
	fontRepo.AddUTF8Font(fontFamily, fontstyle.Normal, "./fonts/RobotoMono-Light.ttf")
	fontRepo.AddUTF8Font(fontFamily, fontstyle.Italic, "./fonts/RobotoMono-Italic.ttf")
	fontRepo.AddUTF8Font(fontFamily, fontstyle.Bold, "./fonts/RobotoMono-Bold.ttf")
	fontRepo.AddUTF8Font(fontFamily, fontstyle.BoldItalic, "./fonts/RobotoMono-BoldItalic.ttf")

	fonts, err := fontRepo.Load()
	if err != nil {
		fmt.Printf("Cannot load fonts, err = %v\n", err)
		return
	}
	fmt.Printf("Fonts = %+v\n", fonts)

	// setup configuration
	pn := props.PageNumber{
		Pattern: "APP 1.2.3 | {current}/{total}",
		Place:   props.RightBottom,
	}

	builder := config.NewBuilder()
	builder.WithPageSize(pagesize.A4)
	builder.WithOrientation(orientation.Horizontal)
	builder.WithPageNumber(pn)
	builder.WithAuthor("Branko", true)
	builder.WithCreator("APP 1.2.3", true)
	builder.WithSubject("Sales", true)
	builder.WithTitle("Simple report", true)
	// builder.WithDebug(true) // debug = true creates frames
	builder.WithCustomFonts(fonts)
	builder.WithDefaultFont(&props.Font{Family: fontFamily})
	fmt.Printf("Builder = %+v\n", builder)

	// build configuration
	config := builder.Build()
	fmt.Printf("Config = %+v\n", config)

	// setup core
	mrt := maroto.New(config)
	mrt = maroto.NewMetricsDecorator(mrt)
	fmt.Printf("Core = %+v\n", mrt)

	// header
	// .........1.........2.........3.........4.........5.........6.........7.........8.........9.........0.........1.........2.........3
	// 1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890
	// xxxxxxxxx1: xxxxxxxxx1xxxxxxxxx2 							         										  yyyy-mm-dd hh:mm:ss
	//
	headerRow1 := row.New(LINE_SIZE)
	headerRow1 = headerRow1.Add(
		text.NewCol(6, fmt.Sprintf("%s: %s ", "Sales", "Simple report")),
		text.NewCol(6, time.Now().Format("2006-01-02 15:04:05"), props.Text{
			Align: align.Right,
		}))
	headerRowDelimiter := text.NewRow(LINE_SIZE, "")

	err = mrt.RegisterHeader(headerRow1, headerRowDelimiter, headerRowDelimiter)
	if err != nil {
		fmt.Printf("Cannot register document header, err = %v\n", err)
		return
	}

	// footer
	footerRowDelimiter := text.NewRow(LINE_SIZE, "")
	err = mrt.RegisterFooter(footerRowDelimiter)
	if err != nil {
		fmt.Printf("Cannot register document footer, err = %v\n", err)
		return
	}

	// template row
	mrt.AddRow(LINE_SIZE, text.NewCol(12, ".........1.........2.........3.........4.........5.........6.........7.........8.........9.........0.........1.........2.........3"))
	mrt.AddRow(LINE_SIZE, text.NewCol(12, "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"))

	// add rows
	for i := 0; i < 500; i++ {
		mrt.AddRow(LINE_SIZE, text.NewCol(12, fmt.Sprintf("%5d", i)))
	}

	// create document
	document, err := mrt.Generate()
	if err != nil {
		fmt.Printf("Cannot generate document, err = %v\n", err)
		return
	}

	// save document
	fileName := "doc/test7.pdf"
	err = document.Save(fileName)
	if err != nil {
		fmt.Printf("Cannot save document, err = %v\n", err)
		return
	}
	fmt.Printf("Saved document, file name = %v\n", fileName)

}
