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

	// test 2: make language tag
	test11(*optLocale)
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

func test11(locale string) {

	fmt.Printf("\n---test11---\n")

	fmt.Printf("Locale = %+v\n", locale)

	// parse locale
	lang := language.Make(locale)
	fmt.Printf("Language tag = %+v\n", lang)
}
