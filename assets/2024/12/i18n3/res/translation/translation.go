package translation

// go generate
// 	gotext
//	-srclang 	BCP 47 tag for base language
// 	update 		gotext command
// 	-out 		path to catalog file
// 	-lang 		list of the BCP 47 tags of languages to create translations for
// 	i18n		(example) module path for the package

//go:generate gotext -srclang=en-US update -lang=hr-HR,en-US,de-DE -out=catalog.go i18n
