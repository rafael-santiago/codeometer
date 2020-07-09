// package codeometersys - gathers default values and useful application system functions.
// --
//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package codeometersys

import (
    "os"
    "fmt"
)

// Shows the application's version.
func version() int {
    fmt.Fprintf(os.Stdout, "codeometer-v%s\n", appVersion)
    return 0
}

func versionHelp() int {
    fmt.Fprintf(os.Stdout, "use: codeometer version\n")
    return 0
}
