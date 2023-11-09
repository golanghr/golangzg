package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

// typedefs
type Item struct {
	Num   string
	Name  string
	Price float64
}

// common declarations
const PROGRAM_NAME string = "csvread"
const PROGRAM_VERSION string = "1.0.0"
const COPYRIGHT string = "Copyright (c) 20xx POINTER d.o.o. All rights reserved."

const MIN_COLUMNS = 3

// option, version
var optVersion = flag.Bool("version", false, "Show program version")
var optHelp = flag.Bool("help", false, "Show help version")

// option, flag true/false
var optTest1 = flag.Bool("test1", false, "Test 1")
var optTest2 = flag.Bool("test2", false, "Test 2 w/ lang, encoding")

// arguments
var optFilename = ""
var optLang = ""
var optEncoding = ""

func main() {

	// command line
	flag.Parse()

	// check options
	// check arguments remaining after flags have been processed
	if flag.NArg() > 0 {
		optFilename = flag.Arg(0)
		optLang = flag.Arg(1)
		optEncoding = flag.Arg(2)
	}

	// select options
	if *optVersion {
		version()
	} else if *optHelp {
		usage()
	} else if *optTest1 {
		test1(optFilename)
	} else if *optTest2 {
		test2(optFilename, optLang, optEncoding)
	}
}

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %v [OPTION] [FILE] [LANG] [ENCODING]\n", os.Args[0])
	flag.PrintDefaults()
}

func version() {
	fmt.Printf("%v %v", PROGRAM_NAME, PROGRAM_VERSION)
}

func test1(fileName string) {

	fmt.Printf("test 1\n")
	fmt.Printf("File name = %v\n", fileName)

	if fileName == "" {
		fmt.Printf("Please specify file name\n")
		return
	}

	// read file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open file, err = %+v", err)
		return
	}
	defer file.Close()

	// setup csv reader
	// setup delimiter
	// allow variable number of fields
	// trim leading space
	// setup handling of quotes
	var csvReader *csv.Reader = csv.NewReader(file)
	csvReader.Comma = ';'
	csvReader.FieldsPerRecord = -1
	csvReader.TrimLeadingSpace = false
	// csvReader.LazyQuotes = true
	csvReader.LazyQuotes = false

	var items []Item
	countRecs := 0
	countValidRecs := 0

	// read by record
	for {

		rec, err := csvReader.Read()

		// check eof
		if err == io.EOF {
			break
		}

		// check rec
		if err != nil {
			fmt.Printf("err = %+v\n", err)
			continue
		}

		fmt.Printf("type = %T, len = %v, rec = %+v\n", rec, len(rec), rec)
		countRecs++

		// fields in record
		if len(rec) < MIN_COLUMNS {
			fmt.Printf("\t Incorrect number of fields in record = %+v\n", countRecs)
			continue
		}
		countValidRecs++

		// record structure
		// 0 = item number
		// 1 = item name
		// 2 = price
		num := rec[0]
		name := rec[1]
		price := rec[2]
		fmt.Printf("\t Num = %+v, Name = %+v, Price = %+v\n", num, name, price)

		// convert to num
		priceVal, err := getNum(price)
		fmt.Printf("\t Price value = %f, err = %+v\n", priceVal, err)

		// build struct
		item := Item{
			Num:   num,
			Name:  name,
			Price: priceVal,
		}

		// build list of items
		items = append(items, item)

	}
	fmt.Printf("Total records read = %+v\n", countRecs)
	fmt.Printf("Total valid records = %+v\n", countValidRecs)
	fmt.Printf("Items = %+v\n", items)

}

func test2(fileName string, lang string, enc string) {

	fmt.Printf("test 2\n")
	fmt.Printf("File name = %v\n", fileName)
	fmt.Printf("Language = %v\n", lang)
	fmt.Printf("Encoding = %v\n", enc)

	if fileName == "" {
		fmt.Printf("Please specify file name\n")
		return
	}

	// read file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open file, err = %+v", err)
		return
	}
	defer file.Close()

	// setup csv reader
	var csvReader *csv.Reader
	if enc == "cp852" {
		decoder := charmap.CodePage852.NewDecoder()
		reader := decoder.Reader(file)
		csvReader = csv.NewReader(reader)
	} else if enc == "win1250" {
		decoder := charmap.Windows1250.NewDecoder()
		reader := decoder.Reader(file)
		csvReader = csv.NewReader(reader)
	} else {
		// default = utf8
		csvReader = csv.NewReader(file)
	}

	// setup delimiter
	// allow variable number of fields
	// trim leading space
	// setup handling of quotes
	csvReader.Comma = ';'
	csvReader.FieldsPerRecord = -1
	csvReader.TrimLeadingSpace = true
	csvReader.LazyQuotes = true

	var items []Item
	countRecs := 0
	countValidRecs := 0

	// read by record
	for {

		rec, err := csvReader.Read()

		// check eof
		if err == io.EOF {
			break
		}

		// check rec
		if err != nil {
			fmt.Printf("err = %+v\n", err)
			continue
		}

		fmt.Printf("type = %T, len = %v, rec = %+v\n", rec, len(rec), rec)
		countRecs++

		// fields in record
		if len(rec) < MIN_COLUMNS {
			fmt.Printf("\t Incorrect number of fields in record = %+v\n", countRecs)
			continue
		}

		countValidRecs++

		// record structure
		// 0 = item number
		// 1 = item name
		// 2 = price
		num := rec[0]
		name := rec[1]
		price := rec[2]
		fmt.Printf("\t Num = %+v, Name = %+v, Price = %+v\n", num, name, price)

		// convert to num
		priceVal, err := getNumForLang(price, lang)
		fmt.Printf("\t Price value = %f, err = %+v\n", priceVal, err)

		// build struct
		item := Item{
			Num:   num,
			Name:  name,
			Price: priceVal,
		}

		// build list of items
		items = append(items, item)

	}
	fmt.Printf("Total records read = %+v\n", countRecs)
	fmt.Printf("Total valid records = %+v\n", countValidRecs)
	fmt.Printf("Items = %+v\n", items)

}

func getNum(numStr string) (float64, error) {

	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return 0.0, errors.New("incorrect format")
	}

	return num, nil
}

func getNumForLang(numStr string, lang string) (float64, error) {

	// fix decimal number by locale
	// hr 		1.234,56 -> 1234.56
	// other	1,234.56 -> 1234.56

	if lang == "hr" {
		// replace . with (blank)
		// replace comma with point
		numStr = strings.Replace(numStr, ".", "", -1)
		numStr = strings.Replace(numStr, ",", ".", -1)
	} else if lang == "en" {
		// replace , with (blank)
		numStr = strings.Replace(numStr, ",", "", -1)
	}

	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return 0.0, errors.New("Cannot parse float, incorrect format for num")
	}

	return num, nil
}
