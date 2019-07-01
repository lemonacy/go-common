package hoo

import (
	"regexp"
)

func Minify(html string) string {
	regex, err := regexp.Compile(`[\r\n]+`)
	PanicOnErr(err)
	html = regex.ReplaceAllString(html, "")

	regex, err = regexp.Compile(`\s+`)
	PanicOnErr(err)
	html = regex.ReplaceAllString(html, " ")

	regex, err = regexp.Compile(`([^\\])>\s+<`)
	PanicOnErr(err)
	html = regex.ReplaceAllString(html, "${1}><")

	regex, err = regexp.Compile(`\s*(>|<|{|})\s*`)
	PanicOnErr(err)
	html = regex.ReplaceAllString(html, "$1")

	return html
}
