// package codeometersys - gathers default values and useful application system functions.
// --
//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package codeometersys

import (
    "os"
)

// The 'help' command handler.
func help() int {
    var topic string
    if len(os.Args) > 1 {
        topic = os.Args[2]
    } else {
        topic = "help"
    }
    command, found := commands()[topic]
    if !found {
        command = CodeometerCommandHandler{nil, unknownCommand}
    }
    return command.Helper()
}
