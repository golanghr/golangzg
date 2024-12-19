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

	// import translation to init catalog
	_ "i18n/res/translation"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
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

	// test 4: parse language tag, init message printer, use message with variable
	test4(*optLocale)

	// test 5: format number for language
	test5()

	// test 6: format number for custom
	test6()
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

func test4(locale string) {

	fmt.Printf("\n---test4---\n")

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

	paymentCode := "CASH"
	revenue := 12345.67
	prn.Printf("Payment = %s\n", paymentCode)
	prn.Printf("Revenue = %f\n", revenue)

}

func test5() {

	fmt.Printf("\n---test5---\n")

	i := 12345
	f := 12345.67

	fmt.Printf("Language = %v", "default")
	fmt.Printf("\t Int = %d \t Float = %.2f\n", i, f)

	tags := []language.Tag{
		language.Croatian,
		language.English,
		language.German,
		language.French,
	}

	languages := []string{
		"hr",
		"en",
		"de",
		"fr",
		"xx",
	}

	// check tags
	for _, tag := range tags {

		fmt.Printf("Language tag = %v", tag)

		prn := message.NewPrinter(tag)
		prn.Printf("\t Int = %d \t Float = %.2f\n", i, f)

	}

	// check languages
	for _, lang := range languages {

		// Make is a convenience wrapper for Parse that omits the error. In case of an error, a sensible default is returned.
		tag := language.Make(lang)
		fmt.Printf("Language tag = %v, language = %v", tag, lang)

		prn := message.NewPrinter(tag)
		prn.Printf("\t Int = %d \t Float = %.2f\n", i, f)
	}
}

func test6() {

	fmt.Printf("\n---test6---\n")

	i := 12345
	f := 12345.67
	p := 12.34

	fmt.Printf("Language = %v", "default")
	fmt.Printf("\t Int = %d \t Float = %.2f\n", i, f)

	tags := []language.Tag{
		language.Croatian,
		language.English,
		language.German,
		language.French,
	}

	// check tags
	for _, tag := range tags {

		fmt.Printf("Language tag = %v", tag)

		prn := message.NewPrinter(tag)
		prn.Printf("\t Int = %v \t Float = %v \t Percent = %v\n", number.Decimal(i), number.Decimal(f), number.Percent(p))

	}

	// options
	tag := language.Croatian
	fmt.Printf("Language tag = %v \n", tag)

	prn := message.NewPrinter(language.Croatian)
	prn.Printf("\t Decimal width = %v\n", number.Decimal(f, number.FormatWidth(15)))
	prn.Printf("\t Decimal pad   = %v\n", number.Decimal(f, number.FormatWidth(15), number.Pad('#')))
}
