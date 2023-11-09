package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/xuri/excelize/v2"
)

// typedefs
type Product struct {
	EanNum    string
	IntNum    string
	Name      string
	Warehouse string
	Price     float64
	VatPerc   float64
	Uom       string
}

// common declarations
const PROGRAM_NAME string = "excelwrite"
const PROGRAM_VERSION string = "1.0.0"
const COPYRIGHT string = "Copyright (c) 20xx POINTER d.o.o. All rights reserved."

// option, version
var optVersion = flag.Bool("version", false, "Show program version")
var optHelp = flag.Bool("help", false, "Show help version")

// option, flag true/false
var optTest1 = flag.Bool("test1", false, "Test 1")

// argument, file name
var optFilename = ""

const NUM_DECIMALS = 2

func main() {

	// command line
	flag.Parse()

	// check options
	// check arguments remaining after flags have been processed
	if flag.NArg() > 0 {
		optFilename = flag.Arg(0)
	}

	// select options
	if *optVersion {
		version()
	} else if *optHelp {
		usage()
	} else if *optTest1 {
		test1(optFilename)
	}
}

func version() {
	fmt.Fprintf(os.Stdout, "%v %v", PROGRAM_NAME, PROGRAM_VERSION)
}

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s\n", os.Args[0])
	flag.PrintDefaults()
}

func test1(fileName string) {

	fmt.Printf("test 1\n")
	fmt.Printf("File name = %v\n", fileName)

	// build data
	var products []Product

	for i := 1; i <= 5; i++ {

		item := Product{
			EanNum:    fmt.Sprintf("385123456789%d", i),
			IntNum:    fmt.Sprintf("%d", i),
			Name:      fmt.Sprintf("Item %d šđžćč ŠĐŽĆČ", i),
			Warehouse: "123",
			Price:     100.0 + float64(i) + 0.34,
			VatPerc:   25.0,
			Uom:       "kom",
		}
		products = append(products, item)

	}
	fmt.Printf("Total records = %v\n", len(products))
	fmt.Printf("Products = %+v\n", products)

	writeProducts(fileName, &products)

}

func writeProducts(fileName string, products *[]Product) {

	if fileName == "" {
		fmt.Printf("Please specify file name\n")
		return
	}

	// new workbook
	file := excelize.NewFile()
	defer file.Close()

	// new sheet
	sheetName := "Sheet1"
	index, err := file.NewSheet(sheetName)
	if err != nil {
		fmt.Printf("Cannot create sheet %v, ret = %v\n", sheetName, err)
		return
	}
	fmt.Printf("Created sheet %v, index = %v\n", sheetName, index)
	file.SetActiveSheet(index)

	// setup sheet
	_ = file.SetColWidth(sheetName, "A", "X", 20)

	// add header
	rowCount := 1
	cellId := fmt.Sprintf("%v%v", "A", rowCount)

	headerRow := []string{"EanNum", "IntNum", "Name", "Warehouse", "Price", "VatPerc", "Uom"}
	err = file.SetSheetRow(sheetName, cellId, &headerRow)
	if err != nil {
		fmt.Printf("Cannot set row, sheet %v, ret = %v\n", sheetName, err)
	} else {
		fmt.Printf("Inserted row, sheet %v, cell id = %v, row = %v\n", sheetName, cellId, headerRow)
	}

	// add data
	for _, product := range *products {

		rowCount++
		cellId = fmt.Sprintf("%v%v", "A", rowCount)

		row := []interface{}{
			product.EanNum,
			product.IntNum,
			product.Name,
			product.Warehouse,
			product.Price,
			product.VatPerc,
			product.Uom,
		}

		err = file.SetSheetRow(sheetName, cellId, &row)
		if err != nil {
			fmt.Printf("Cannot set row, sheet %v, ret = %v\n", sheetName, err)
		} else {
			fmt.Printf("Inserted row, sheet %v, cell id = %v, row = %v\n", sheetName, cellId, row)
		}
	}

	// sheet name
	const REPORT_CATEGORY = "report"
	const REPORT_STATUS = "final"
	const REPORT_ID = "xlsx"
	const REPORT_SUBJECT = "Sales"
	const REPORT_TITLE = "Products"

	err = file.SetSheetName(sheetName, REPORT_TITLE)
	if err != nil {
		fmt.Printf("Cannot set sheet name, ret = %v\n", err)
	} else {
		fmt.Printf("Set sheet name, title = %v\n", REPORT_TITLE)
	}

	// document properties
	currentTime := time.Now()
	timestamp := currentTime.Format("2006-01-02T15:04:05")

	docprops := excelize.DocProperties{
		Category:       REPORT_CATEGORY,
		ContentStatus:  REPORT_STATUS,
		Created:        timestamp,
		Creator:        PROGRAM_NAME,
		Description:    fmt.Sprintf("%v, %v", PROGRAM_NAME, REPORT_CATEGORY),
		Identifier:     REPORT_ID,
		Keywords:       fmt.Sprintf("%v, %v", PROGRAM_NAME, REPORT_CATEGORY),
		LastModifiedBy: PROGRAM_NAME,
		Modified:       timestamp,
		Revision:       "123",
		Subject:        REPORT_SUBJECT,
		Title:          REPORT_TITLE,
		Language:       "",
		Version:        "",
	}

	err = file.SetDocProps(&docprops)
	if err != nil {
		fmt.Printf("Cannot set document properties, ret = %v\n", err)
	} else {
		fmt.Printf("Set document properties = %+v\n", docprops)
	}

	// save workbook
	err = file.SaveAs(fileName)
	if err != nil {
		fmt.Printf("Cannot save workbook %v, ret = %v\n", fileName, err)

	}
}
