package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// typedefs
type Item struct {
	Num   string
	Name  string
	Price float64
}

// common declarations
const PROGRAM_NAME string = "csvwrite"
const PROGRAM_VERSION string = "1.0.0"
const COPYRIGHT string = "Copyright (c) 20xx POINTER d.o.o. All rights reserved."

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

	// build data
	var items []Item

	for i := 1; i <= 5; i++ {

		item := Item{
			Num:   fmt.Sprintf("%d", 1234567890120+i),
			Name:  fmt.Sprintf("Item %d šđžćč ŠĐŽĆČ", i),
			Price: 1234.50 + float64(i),
		}
		items = append(items, item)

	}
	fmt.Printf("Total records = %v\n", len(items))
	fmt.Printf("Items = %+v\n", items)

	// create file
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot create file, ret = %+v", err)
		return
	}
	defer file.Close()

	// setup csv writer
	// default = utf8
	// setup delimiter
	var csvWriter *csv.Writer = csv.NewWriter(file)
	csvWriter.Comma = ';'
	defer csvWriter.Flush()

	// write by record
	for _, value := range items {

		item := value
		fmt.Printf("Item = %+v\n", item)

		rec := []string{
			item.Num,
			item.Name,
			fmt.Sprintf("%f", item.Price),
		}
		fmt.Printf("Rec = %+v\n", rec)

		err = csvWriter.Write(rec)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot create file, ret = %+v", err)
			return
		}
	}

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

	// build data
	var items []Item

	for i := 1; i <= 5; i++ {

		item := Item{
			Num:   fmt.Sprintf("%d", 1234567890120+i),
			Name:  fmt.Sprintf("Item %d šđžćč ŠĐŽĆČ", i),
			Price: 1234.50 + float64(i),
		}
		items = append(items, item)

	}
	fmt.Printf("Total records = %v\n", len(items))
	fmt.Printf("Items = %+v\n", items)

	// create file
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot create file, ret = %+v", err)
		return
	}
	defer file.Close()

	// setup printer
	tag := language.Make(lang)
	prn := message.NewPrinter(tag)

	// setup csv writer
	// default = utf8
	var csvWriter *csv.Writer
	if enc == "cp852" {
		encoder := charmap.CodePage852.NewEncoder()
		writer := encoder.Writer(file)
		csvWriter = csv.NewWriter(writer)
	} else if enc == "win1250" {
		encoder := charmap.Windows1250.NewEncoder()
		writer := encoder.Writer(file)
		csvWriter = csv.NewWriter(writer)
	} else {
		csvWriter = csv.NewWriter(file)
	}

	// setup delimiter
	csvWriter.Comma = ';'
	defer csvWriter.Flush()

	// write header
	rec := []string{
		"NUM",
		"NAME",
		"PRICE",
	}
	err = csvWriter.Write(rec)
	if err != nil {
		fmt.Printf("Cannot write to file, err = %+v", err)
		return
	}

	// write by record
	countRecs := 0
	for _, value := range items {

		item := value
		fmt.Printf("Item = %+v\n", item)

		rec := []string{
			item.Num,
			item.Name,
			prn.Sprintf("%.2f", item.Price),
		}
		fmt.Printf("Rec = %+v\n", rec)

		err = csvWriter.Write(rec)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot create file, ret = %+v", err)
			return
		}
		countRecs++
	}
	fmt.Printf("Number of records written = %v", countRecs)

}
