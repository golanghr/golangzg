package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"

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
const PROGRAM_NAME string = "excelread"
const PROGRAM_VERSION string = "1.0.0"
const COPYRIGHT string = "Copyright (c) 20xx POINTER d.o.o. All rights reserved."

// option, version
var optVersion = flag.Bool("version", false, "Show program version")
var optHelp = flag.Bool("help", false, "Show help version")

// option, flag true/false
var optTest1 = flag.Bool("test1", false, "Test 1")

// argument, file name
var optFilename = ""

const MIN_COLUMNS = 7
const NUM_DECIMALS = 2

func main() {

	// command line
	flag.Parse()

	// verify options
	// verify arguments remaining after flags have been processed
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

	if fileName == "" {
		fmt.Printf("Please specify file name\n")
		return
	}

	// workbook
	file, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Printf("Cannot open workbook, ret = %v\n", err)
		return
	}
	defer file.Close()

	// workbook properties
	props, err := file.GetWorkbookProps()
	if err != nil {
		fmt.Printf("Cannot read workbook properties, ret = %v\n", err)
		return
	}
	fmt.Printf("Workbook properties = %+v\n", props)

	docprops, err := file.GetDocProps()
	if err != nil {
		fmt.Printf("Cannot read document properties, ret = %v\n", err)
		return
	}
	fmt.Printf("Document properties = %+v\n", docprops)

	// sheets
	sheets := file.GetSheetList()
	fmt.Printf("List of sheets = %v\n", sheets)

	for _, name := range sheets {
		sheetprops, err := file.GetSheetProps(name)
		if err != nil {
			fmt.Printf("Cannot read sheet properties, ret = %v\n", err)
		} else {
			fmt.Printf("Sheet = %v, properties = %+v\n", name, sheetprops)
		}
	}

	// get active sheet
	sheetName := ""
	activeSheet := file.GetActiveSheetIndex()
	if activeSheet <= 0 {
		fmt.Printf("Cannot get active sheet\n")
		sheetName = sheets[0]
	} else {
		sheetName = sheets[activeSheet]
	}
	fmt.Printf("Active sheet index = %v\n", activeSheet)
	fmt.Printf("Active sheet name = %v\n", sheetName)

	// read sheet
	products := make([]Product, 0)
	sheetName = "Sheet1"
	fmt.Printf("Selected sheet name = %v\n", sheetName)

	rows, err := file.GetRows(sheetName)
	if err != nil {
		fmt.Printf("Cannot read rows for sheet = %v, ret = %v\n", sheetName, err)
	} else {

		for i, row := range rows {
			fmt.Printf("Found row %v, type = %T, len = %v, data = %v\n", i, row, len(row), row)

			// skip header row
			if i == 0 {
				fmt.Printf("Row %v is header, skipped\n", i)
				continue
			}

			// verify empty row
			if len(row) <= 0 {
				fmt.Printf("Row %v is empty, skipped\n", i)
				continue
			}

			// verify incorrect row
			if len(row) < MIN_COLUMNS {
				fmt.Printf("Row %v has incorrect number of columns, skipped\n", i)
				continue
			}

			// convert row data
			product := Product{
				EanNum:    row[0],
				IntNum:    row[1],
				Name:      row[2],
				Warehouse: row[3],
				Price:     getNumRounded(row[4]),
				VatPerc:   getPercentRounded(row[5]),
				Uom:       row[6],
			}
			products = append(products, product)
		}
		fmt.Printf("Total records read = %v\n", len(products))
		fmt.Printf("Products = %+v\n", products)

	}
}

func getNumRounded(numStr string) float64 {

	// parse
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		fmt.Printf("Cannot parse float, incorrect format for num\n")
		num = 0.0
	}

	// round
	exp := math.Pow(10, NUM_DECIMALS)
	num = math.Round(num*exp) / exp

	return num

}

func getPercentRounded(numStr string) float64 {

	// parse
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		fmt.Printf("Cannot parse float, incorrect format for percent\n")
		num = 0.0
	}

	// round
	exp := math.Pow(10, NUM_DECIMALS)
	num = math.Round(num*exp) / exp

	// ? multiply by 100
	// num = num * exp

	return num

}
