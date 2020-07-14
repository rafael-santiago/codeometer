// package loader concentrates all code loading stuff, however all you should call is LoadCode().
// --
//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package loader

import (
    "regexp"
)

// Strips off temporary directory from a path
func getCodeKey(srcpath string) string {
    pattern := regexp.MustCompile(`^.*:\\.*\\codeometer-temp[0123456789]+\\`)
    return string(pattern.ReplaceAll([]byte(srcpath), []byte("")))
}
