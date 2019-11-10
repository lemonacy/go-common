/**
 * Author: dengkuadong
 * Date: 2019-07-05
 */

package hoo

import "regexp"

func Split(pattern, text string) []string {
    regex := regexp.MustCompile(pattern)
    return regex.Split(text, -1)
}
