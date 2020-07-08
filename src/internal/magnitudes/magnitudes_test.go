//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package magnitudes

import (
    "testing"
)

func TestGetCourier12pxParams(t *testing.T) {
    charPerLineNr, charPerPageNr := GetCourier12pxParams()
    if charPerLineNr != charPerLineA4Courier12px {
        t.Errorf(`charPerLineNr != charPerLineA4Courier12px: %v != %v`, charPerLineNr, charPerLineA4Courier12px)
    }
    if charPerPageNr != charPerPageA4Courier12px {
        t.Errorf(`charPerPageNr != charPerPageA4Courier12px: %v != %v`, charPerPageNr, charPerPageA4Courier12px)
    }
}

func TestGetCourier10pxParams(t *testing.T) {
    charPerLineNr, charPerPageNr := GetCourier10pxParams()
    if charPerLineNr != charPerLineA4Courier10px {
        t.Errorf(`charPerLineNr != charPerLineA4Courier10px: %v != %v`, charPerLineNr, charPerLineA4Courier10px)
    }
    if charPerPageNr != charPerPageA4Courier10px {
        t.Errorf(`charPerPageNr != charPerPageA4Courier10px: %v != %v`, charPerPageNr, charPerPageA4Courier10px)
    }
}

func TestGetA4PaperWidthSizeInMM(t *testing.T) {
    if size := GetA4PaperWidthSizeInMM(); size != a4PaperWidthInMM {
        t.Errorf(`GetA4PaperSizeInMM() != A4PaperWidthInMM: %v != %v`, size, a4PaperWidthInMM)
    }
}
