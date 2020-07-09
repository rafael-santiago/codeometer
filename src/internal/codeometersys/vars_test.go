//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package codeometersys

import (
    "testing"
)

func TestAppVersion(t *testing.T) {
    if AppVersion != "1" {
        t.Error(`AppVersion != "1"`)
    }
}
