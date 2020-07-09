//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package codeometersys

import (
    "testing"
)

func Test_appVersion(t *testing.T) {
    if appVersion != "1" {
        t.Error(`appVersion != "1"`)
    }
}

func Test_commands(t *testing.T) {
    expectedHandlers := map[string]CodeometerHandlerFunc {
        "measure" : measure,
        "httpd" : httpd,
        "man" : man,
        "help" : showHelpBanner,
        "version" : showAppVersion,
        "" : showHelpBanner,
    }
    returnedHandlers := commands()
    if returnedHandlers == nil {
        t.Error(`returnedHandlers == nil`)
        t.Fail()
    }
    if len(returnedHandlers) != len(expectedHandlers) {
        t.Error(`len(returnedHandlers) != len(expectedHandlers)`)
    }
    for k, _ := range expectedHandlers {
        f, foundHandler := returnedHandlers[k]
        if !foundHandler {
            t.Errorf(`!foundHandler: %v : %v`, k, returnedHandlers)
        }
        if f == nil {
            t.Error(`f == nil : k=%v`, k)
        }
    }
}
