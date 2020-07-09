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
    "internal/options"
)

// Shows the default help banner.
func showHelpBanner() int {
    fmt.Fprintf(os.Stdout, "%s", helpBanner)
    return 0
}

// Error handler for unknown commands.
func unknownCommand() int {
    fmt.Fprintf(os.Stderr, "error: Unknown command: '%s'.\n", options.GetCommand())
    return 1
}

// Runs codeometer based on passed user's command line.
func Run() int {
    command, found := commands()[options.GetCommand()]
    if !found {
        command = CodeometerCommandHandler{unknownCommand, nil}
    }
    return command.Runner()
}
