// package options - offers conveniences for user's command line options reading.
// --
//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package options

import (
    "strings"
    "os"
)

func GetCommand() string {
    if len(os.Args) > 1 {
        return os.Args[1]
    }
    return ""
}

func GetOption(option, defaultValue string) string {
    if len(os.Args) > 2 {
        wanted := "--" + option + "="
        for _, a := range os.Args[2:] {
            if strings.HasPrefix(a, wanted) {
                return a[len(wanted):]
            }
        }
    }
    return defaultValue
}

func GetBoolOption(option string, defaultValue bool) bool {
    if len(os.Args) > 2 {
        wanted := "--" + option
        for _, a := range os.Args[2:] {
            if wanted == a {
                return true
            }
        }
    }
    return defaultValue
}

func GetArrayOption(option string, defaultValue ...string) []string {
    if data := GetOption(option, ""); len(data) > 0 {
        return strings.Split(data, ",")
    }
    return defaultValue
}
