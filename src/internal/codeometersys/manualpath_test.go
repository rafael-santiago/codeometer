//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package codeometersys

import (
    "testing"
    "runtime"
)

func Test_manualPath(t *testing.T) {
    manualPath := manualPath()
    if runtime.GOOS == "linux" || runtime.GOOS == "freebsd" {
        if manualPath != "/usr/local/share/codeometer/doc/manual.txt" {
            t.Error(`manualPath != "/usr/local/share/codeometer/doc/manual.txt"`)
        }
    } else if runtime.GOOS == "windows" {
        if manualPath != "C:\\codeometer\\doc\\manual.txt" {
            t.Error(`manualPath != "C:\\codeometer\\doc\\manual.txt"`)
        }
    } else {
        t.Error(`Unsupported platform.`)
        t.Fail()
    }
}
