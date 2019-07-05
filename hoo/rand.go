/**
 * Copyright (C) 1995-2019 Seasun Entertainment 西山居 版权所有;
 *
 * Author: dengkuadong <dengkuadong@kingsoft.com>
 * Since: 2019-07-05
 */

package hoo

import (
    "math/rand"
    "time"
)

var letters = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
    rand.Seed(time.Now().UnixNano())
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}
