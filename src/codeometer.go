// package main - guess what about is this.
// --
//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package main

import (
    "internal/ruler"
    "internal/measurer"
    "fmt"
)

func main() {
    c := &ruler.CodeStat{}
    c.CharPerPage = 10
    c.CharPerLine = 1
    c.Files = make(map[string]ruler.CodeFileInfo)
    c.Files["CodeStat"] = ruler.CodeFileInfo{101}
    mm := &measurer.MMCodeStat{}
    mm.CharPerPage = 20
    mm.CharPerLine = 2
    mm.Files = make(map[string]ruler.CodeFileInfo)
    mm.Files["MMCodeStat"] = ruler.CodeFileInfo{202}
    m := &measurer.MCodeStat{}
    fmt.Println("c = ", c)
    m.Calibrate(c)
    fmt.Println("m = ", m)
    m.Calibrate(mm)
    fmt.Println("mm = ", mm)
    m.Files["Meter"] = ruler.CodeFileInfo{303}
    fmt.Println("m = ", m)
    fmt.Println("mm = ", mm)
}
