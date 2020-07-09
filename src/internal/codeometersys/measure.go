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
)

// The 'measure' command handler.
func measure() int {
    fmt.Fprintf(os.Stderr, "Not implemented.\n")
    return 1
}

// The 'measure' command helper.
func measureHelp() int {
    fmt.Fprintf(os.Stdout, "use: codeometer measure --src=<file path | zip file path | git repo url | directory path> --exts=<extensions>.\n")
    return 0
}

