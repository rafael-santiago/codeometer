// package codeometersys - gathers default values and useful application system functions.
// --
//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package codeometersys

import (
    "fmt"
    "os"
    "io/ioutil"
)

// The 'man' command handler.
func man() int {
    var exitCode int
    data, err := ioutil.ReadFile(manualPath())
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v.\n", err)
        exitCode = 1
    } else {
        fmt.Fprintf(os.Stdout, "%s\n", string(data))
    }
    return exitCode
}
