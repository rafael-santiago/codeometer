// package loader concentrates all code loading stuff, however all you should call is LoadCode().
// --
package loader

import (
    "strings"
)

// Strips off temporary directory from a path
func getCodeKey(srcpath string) string {
    return strings.Replace(srcpath, "/codeometer-temp/", "", -1)
}
