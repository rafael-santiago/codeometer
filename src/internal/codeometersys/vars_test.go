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
    expectedHandlers :=  map[string]CodeometerCommandHandler {
        "measure" : CodeometerCommandHandler{measure, measureHelp},
        "httpd" : CodeometerCommandHandler{httpd, httpdHelp},
        "man" : CodeometerCommandHandler{man, manHelp},
        "version" : CodeometerCommandHandler{version, versionHelp},
        "help" : CodeometerCommandHandler{showHelpBanner, showHelpBanner},
        "" : CodeometerCommandHandler{showHelpBanner, showHelpBanner},
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
        if f.Runner == nil {
            t.Error(`f.Runner == nil : k=%v`, k)
        }
        if f.Helper == nil {
            t.Error(`f.Helper == nil : k=%v`, k)
        }
    }
}
