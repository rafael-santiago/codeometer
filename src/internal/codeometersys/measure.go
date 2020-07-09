// package codeometersys - gathers default values and useful application system functions.
// --
//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package codeometersys

import (
    "fmt"
    "os"
    "internal/ruler"
    "internal/options"
    "internal/loader"
    "internal/measurer"
)

// The 'measure' command handler.
func measure() int {
    src := options.GetOption("src", "")
    if len(src) == 0 {
        fmt.Fprintf(os.Stderr, "error: --src option is missing.\n")
        return 1
    }

    exts := options.GetArrayOption("exts")

    if len(exts) == 0 {
        fmt.Fprintf(os.Stderr, "error: --exts option is missing.\n")
        return 1
    }

    codestat := &ruler.CodeStat{}

    fontSize := options.GetOption("font-size", "12px")
    if fontSize == "12px" {
        codestat.CalibrateCourier12px()
    } else if fontSize == "10px" {
        codestat.CalibrateCourier12px()
    } else {
        fmt.Fprintf(os.Stderr, "error: '%s' font size is invalid. It must be '--10px' or '--12px' (default).\n", fontSize)
        return 1
    }

    err := loader.LoadCode(codestat, src, exts...)

    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %s\n", err)
        return 1
    }

    measurers := map[string]interface{}{"mi" : &measurer.MICodeStat{},
                                        "km" : &measurer.KMCodeStat{},
                                        "m"  : &measurer.MCodeStat{},
                                        "mm" : &measurer.MMCodeStat{}}

    var info string

    wantedMeasures := options.GetArrayOption("measures", "km")

    for _, wantedMeasure := range wantedMeasures {
        m, found := measurers[wantedMeasure]
        if !found {
            fmt.Fprintf(os.Stderr, "error: '%s' is a unknown measure. It must be a list containing: 'mi', 'km', 'm' or 'mm'.\n", wantedMeasure)
            return 1
        }
        var totalDistance float64
        switch m.(type) {
            case *measurer.MICodeStat:
                o := m.(*measurer.MICodeStat)
                o.Calibrate(codestat)
                totalDistance = o.TotalDistance()
                break

            case *measurer.KMCodeStat:
                o := m.(*measurer.KMCodeStat)
                o.Calibrate(codestat)
                totalDistance = o.TotalDistance()
                break

            case *measurer.MCodeStat:
                o := m.(*measurer.MCodeStat)
                o.Calibrate(codestat)
                totalDistance = o.TotalDistance()
                break

            case *measurer.MMCodeStat:
                o := m.(*measurer.MMCodeStat)
                o.Calibrate(codestat)
                totalDistance = o.TotalDistance()
                break
        }
        if len(info) > 0 {
            info += ", "
        }
        if len(wantedMeasure) > 1 {
            wantedMeasure = " " + wantedMeasure
        }
        info += fmt.Sprintf("%.2f%s", totalDistance, wantedMeasure)
    }

    info += ".\n"

    fmt.Fprintf(os.Stdout, "%s has %s", src, info)

    return 0
}

// The 'measure' command helper.
func measureHelp() int {
    fmt.Fprintf(os.Stdout, "use: codeometer measure --src=<file path | zip file path | git repo url | directory path>\n"+
                           "                        --exts=<extensions> --font-size=<font-size> --measures=<measures>.\n")
    return 0
}

