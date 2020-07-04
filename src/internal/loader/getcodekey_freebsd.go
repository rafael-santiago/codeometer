// package loader concentrates all code loading stuff, however all you should call is LoadCode().
// --
//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package loader

import (
    "strings"
)

// Strips off temporary directory from a path
func getCodeKey(srcpath string) string {
    return strings.Replace(srcpath, "/codeometer-temp/", "", -1)
}
