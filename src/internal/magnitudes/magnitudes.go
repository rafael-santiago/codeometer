// package magnitudes gathers constant values used on all measurements stuff.
//--
//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package magnitudes

const a4PaperWidthInMM = 297
const charPerLineA4Courier12px = 82
const charPerPageA4Courier12px = 5002
const charPerLineA4Courier10px = 99
const charPerPageA4Courier10px = 7326

func GetCourier12pxParams() (charPerLineNr, charPerPageNr int64) {
    return charPerLineA4Courier12px, charPerPageA4Courier12px
}

func GetCourier10pxParams() (charPerLineNr, charPerPageNr int64) {
    return charPerLineA4Courier10px, charPerPageA4Courier10px
}

func GetA4PaperWidthSizeInMM() int {
    return a4PaperWidthInMM
}
