/**
 * Copyright (C) 1995-2019 Seasun Entertainment 西山居 版权所有;
 *
 * Author: dengkuadong <dengkuadong@kingsoft.com>
 * Since: 2019-07-05
 */

package hoo

import "regexp"

func Split(pattern, text string) []string {
    regex := regexp.MustCompile(pattern)
    return regex.Split(text, -1)
}
