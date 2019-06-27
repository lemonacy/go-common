package htmlutil

import (
	"github.com/lemonacy/go-common/lang"
	"regexp"
)

func Minify(html string) string {
	regex, err := regexp.Compile(`(([^\\]))>\s+<`)
	lang.PanicOnErr(err)
	html = regex.ReplaceAllString(html, "${1}><")
	return html
}