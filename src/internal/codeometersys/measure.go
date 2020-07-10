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
    "internal/estimator"
    "sort"
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

    if options.GetBoolOption("stats-per-file", false) {
        var files []string

        for k, _ := range codestat.Files {
            files = append(files, k)
        }

        sort.Strings(files)

        for _, file := range files {
            fileInfo := ""
            for _, wantedMeasure := range wantedMeasures {
                m, _ := measurers[wantedMeasure]
                var totalDistance float64
                switch m.(type) {
                    case *measurer.MICodeStat:
                        o := m.(*measurer.MICodeStat)
                        o.Calibrate(codestat)
                        totalDistance = o.DistancePerFile(file)
                        break

                    case *measurer.KMCodeStat:
                        o := m.(*measurer.KMCodeStat)
                        o.Calibrate(codestat)
                        totalDistance = o.DistancePerFile(file)
                        break

                    case *measurer.MCodeStat:
                        o := m.(*measurer.MCodeStat)
                        o.Calibrate(codestat)
                        totalDistance = o.DistancePerFile(file)
                        break

                    case *measurer.MMCodeStat:
                        o := m.(*measurer.MMCodeStat)
                        o.Calibrate(codestat)
                        totalDistance = o.DistancePerFile(file)
                        break
                }
                if len(fileInfo) > 0 {
                    fileInfo += ", "
                }
                if len(wantedMeasure) > 1 {
                    wantedMeasure = " " + wantedMeasure
                }
                fileInfo += fmt.Sprintf("%.2f%s", totalDistance, wantedMeasure)
            }
            info += fmt.Sprintf(" %s has %s.\n", file, fileInfo)
        }

        info += "\n"
    }

    if options.GetBoolOption("estimatives", false) {
        info += "\n"

        estimators := []interface{}{
            &estimator.ChineseGreatWallEstimator{},
            &estimator.PaulistaAvenueEstimator{},
            &estimator.ArcDeTriompheEstimator{},
            &estimator.BigBangEstimator{},
            &estimator.ChristTheRedeemerEstimator{},
            &estimator.ColiseumEstimator{},
            &estimator.EiffelTowerEstimator{},
            &estimator.EmpireStateBuildingEstimator{},
            &estimator.IguazuFallsEstimator{},
            &estimator.LibertyStatueEstimator{},
            &estimator.NiagaraFallsEstimator{},
            &estimator.PantheonEstimator{},
            &estimator.SistineChapelEstimator{},
            &estimator.WallStreetEstimator{},
            &estimator.WashingtonMonumentEstimator{},
            &estimator.FrogTravelerEstimator{},
        }

        for _, e := range estimators {
            switch e.(type) {
                case *estimator.ChineseGreatWallEstimator:
                    o := e.(*estimator.ChineseGreatWallEstimator)
                    info += o.Estimate(codestat)
                    break

                case *estimator.PaulistaAvenueEstimator:
                    o := e.(*estimator.PaulistaAvenueEstimator)
                    info += o.Estimate(codestat)
                   break

                case *estimator.ArcDeTriompheEstimator:
                    o := e.(*estimator.ArcDeTriompheEstimator)
                    info += o.Estimate(codestat)
                    break

                case *estimator.BigBangEstimator:
                    o := e.(*estimator.BigBangEstimator)
                    info += o.Estimate(codestat)
                    break

                case *estimator.ChristTheRedeemerEstimator:
                    o := e.(*estimator.ChristTheRedeemerEstimator)
                    info += o.Estimate(codestat) 
                    break

                case *estimator.ColiseumEstimator:
                    o := e.(*estimator.ColiseumEstimator)
                    info += o.Estimate(codestat)
                    break

                case *estimator.EiffelTowerEstimator:
                    o := e.(*estimator.EiffelTowerEstimator)
                    info += o.Estimate(codestat)
                    break

                case *estimator.EmpireStateBuildingEstimator:
                    o := e.(*estimator.EmpireStateBuildingEstimator)
                    info += o.Estimate(codestat)
                    break

                case *estimator.IguazuFallsEstimator:
                    o := e.(*estimator.IguazuFallsEstimator)
                    info += o.Estimate(codestat)
                    break

                case *estimator.LibertyStatueEstimator:
                    o := e.(*estimator.LibertyStatueEstimator)
                    info += o.Estimate(codestat)
                    break

                case *estimator.NiagaraFallsEstimator:
                    o := e.(*estimator.NiagaraFallsEstimator)
                    info += o.Estimate(codestat)
                    break

                case *estimator.PantheonEstimator:
                    o := e.(*estimator.PantheonEstimator)
                    info += o.Estimate(codestat)
                    break

                case *estimator.SistineChapelEstimator:
                    o := e.(*estimator.SistineChapelEstimator)
                    info += o.Estimate(codestat)
                    break

                case *estimator.WallStreetEstimator:
                    o := e.(*estimator.WallStreetEstimator)
                    info += o.Estimate(codestat)
                    break

                case *estimator.WashingtonMonumentEstimator:
                    o := e.(*estimator.WashingtonMonumentEstimator)
                    info += o.Estimate(codestat)
                    break

                case *estimator.FrogTravelerEstimator:
                    o := e.(*estimator.FrogTravelerEstimator)
                    info += o.Estimate(codestat)
                    break
            }

            info += "\n"
        }
        info += "\n"
    }

    fmt.Fprintf(os.Stdout, "%s has %s", src, info)

    return 0
}

// The 'measure' command helper.
func measureHelp() int {
    fmt.Fprintf(os.Stdout, "use: codeometer measure --src=<file path | zip file path | git repo url | directory path>\n"+
                           "                        --exts=<extensions> [--font-size=<font-size> --measures=<measures>\n" +
                           "                                             --stats-per-file --estimatives]\n")
    return 0
}

