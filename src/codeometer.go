// package main - guess what about is this.
// --
//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package main

import (
    "internal/codeometersys"
    "internal/options"
    "fmt"
    "os"
    "io/ioutil"
)

type CodeometerHandlerFunc map[string]func()int

var gCodeometerCommands CodeometerHandlerFunc = CodeometerHandlerFunc {
    "measure" : measure,
    "httpd" : httpd,
    "man" : man,
    "" : showHelpBanner,
    "help" : showHelpBanner,
    "version" : showAppVersion,
}

func main() {
    command, found := gCodeometerCommands[options.GetCommand()]
    if !found {
        command = unknownCommand
    }
    os.Exit(command())
}

func measure() int {
    fmt.Fprintf(os.Stderr, "Not implemented.\n")
    return 1
}

func httpd() int {
    fmt.Fprintf(os.Stderr, "Not implemented.\n")
    return 1
}

func man() int {
    var exitCode int
    data, err := ioutil.ReadFile(codeometersys.ManualPath())
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v.\n", err)
        exitCode = 1
    } else {
        fmt.Fprintf(os.Stdout, "%s\n", string(data))
    }
    return exitCode
}

func showHelpBanner() int {
    fmt.Fprintf(os.Stdout, "%s", codeometersys.HelpBanner)
    return 0
}

func showAppVersion() int {
    fmt.Fprintf(os.Stdout, "codeometer-v%s\n", codeometersys.AppVersion)
    return 0
}

func unknownCommand() int {
    fmt.Fprintf(os.Stderr, "error: Unknown command: '%s'.\n", options.GetCommand())
    return 1
}
