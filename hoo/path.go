/**
 * Author: dengkuadong
 * Since: 2019-07-05
 */

package hoo

import (
    "os"
    "path/filepath"
)

func CreateSameNameDir(path string) (string, error) {
    dir := filepath.Join(filepath.Dir(path), filepath.Base(path)[0:len(filepath.Base(path))-len(filepath.Ext(path))])
    b, err := FileExists(dir)
    if err != nil {
        return "", err
    }
    if !b {
        err = os.Mkdir(dir, os.ModePerm)
        if err != nil {
            return "", err
        }
    }
    return dir, nil
}
