//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.package ruler
package ruler

import (
    "testing"
    "internal/magnitudes"
)

func TestCalibrateCourier12px(t *testing.T) {
    codestat := &CodeStat{}
    codestat.CalibrateCourier12px()
    charPerLine, charPerPage := magnitudes.GetCourier12pxParams()
    if codestat.CharPerLine != charPerLine {
        t.Errorf(`charPerLine != codestat.CharPerLine: %v != %v`, charPerLine, codestat.CharPerLine)
    }
    if codestat.CharPerPage != charPerPage {
        t.Errorf(`codestat.CharPerPage != charPerPage: %v != %v`, codestat.CharPerPage, charPerPage)
    }
}

func TestCalibrateCourier10x(t *testing.T) {
    codestat := &CodeStat{}
    codestat.CalibrateCourier10px()
    charPerLine, charPerPage := magnitudes.GetCourier10pxParams()
    if codestat.CharPerLine != charPerLine {
        t.Errorf(`charPerLine != codestat.CharPerLine: %v != %v`, charPerLine, codestat.CharPerLine)
    }
    if codestat.CharPerPage != charPerPage {
        t.Errorf(`codestat.CharPerPage != charPerPage: %v != %v`, codestat.CharPerPage, charPerPage)
    }
}
