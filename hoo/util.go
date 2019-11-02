/**
 * Author: dengkuadong <dengkuadong@kingsoft.com>
 * Since: 2019-06-28
 */

package hoo

import (
	"os"
	"regexp"
	"strings"
	"time"
)

func Delay(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func Contains(list []string, ele string) bool {
	for _, v := range list {
		if ele == v {
			return true
		}
	}
	return false
}

func FileExists(filepath string) (bool, error) {
	_, err := os.Stat(filepath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Trim(s string) string {
	s = strings.TrimSpace(s)
	rgx, err := regexp.Compile(` +`)
	PanicOnErr(err)
	s = rgx.ReplaceAllString(s, " ")
	return s
}

func SubmatchOne(s, regex string) string {
	rgx, err := regexp.Compile(regex)
	PanicOnErr(err)
	substr := rgx.FindStringSubmatch(s)
	return substr[1]
}
