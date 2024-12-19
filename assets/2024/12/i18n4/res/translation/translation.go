package translation

// gotext options
//
// main
//	-srclang 	the source-code language (default = en-US)
// 	dir 		default subdirectory to store translation files (default = locales)
//
// 	update
// 	-lang 		comma-separated list of languages to process (default = en-US)
// 	-out 		output file to write to (default = blank)
//
// 	i18n		(example) package

//go:generate gotext -srclang=en-US update -lang=hr-HR,en-US,de-DE -out=catalog.go i18n
