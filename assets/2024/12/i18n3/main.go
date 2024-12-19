// i18n - Internationalization
//
// Remarks
//   - none
//
// See i18n.txt for program notes.
//
// See CHANGELOG.txt for revision history.
package main

import (
	"flag"
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// imports

// typedefs

// common declarations

// option, has default value
var optLocale = flag.String("locale", "en-US", "Define locale")

func main() {

	// command line
	flag.Parse()

	// test 1: parse language tag
	test1(*optLocale)

	// test 2: parse language tag, init message printer
	test2(*optLocale)

	// test 3: parse language tag, init message printer, use message
	test3(*optLocale)

}

func test1(locale string) {

	fmt.Printf("\n---test1---\n")

	fmt.Printf("Locale = %+v\n", locale)

	// parse locale
	lang, err := language.Parse(locale)
	if err != nil {
		fmt.Printf("Cannot parse language, err = %+v\n", err)
		return
	}
	fmt.Printf("Language tag = %+v\n", lang)

}

func test2(locale string) {

	fmt.Printf("\n---test2---\n")

	fmt.Printf("Locale = %+v\n", locale)

	// parse locale
	lang, err := language.Parse(locale)
	if err != nil {
		fmt.Printf("Cannot parse language, err = %+v\n", err)
		return
	}
	fmt.Printf("Language tag = %+v\n", lang)

	// setup message printer for language
	prn := message.NewPrinter(lang)
	fmt.Printf("Message printer = %+v\n", prn)
}

func test3(locale string) {

	fmt.Printf("\n---test3---\n")

	fmt.Printf("Locale = %+v\n", locale)

	// parse locale
	lang, err := language.Parse(locale)
	if err != nil {
		fmt.Printf("Cannot parse language, err = %+v\n", err)
		return
	}
	fmt.Printf("Language tag = %+v\n", lang)

	// setup message printer for language
	prn := message.NewPrinter(lang)
	fmt.Printf("Message printer = %+v\n", prn)

	// print message
	prn.Printf("Welcome to i18n\n")
}
