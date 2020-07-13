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
    "path/filepath"
    "sort"
)

// The 'measure' command handler.
func measure() int {
    var exitCode int
    src := options.GetOption("src", "")
    info, err := measureReport(src,
                               options.GetArrayOption("exts"),
                               options.GetOption("font-size", "12px"),
                               options.GetArrayOption("measures", "km"),
                               options.GetBoolOption("stats-per-file", false),
                               options.GetBoolOption("estimatives", false))
    if err == nil {
        fmt.Fprintf(os.Stdout, "%s has %s", filepath.Base(src), info)
    } else {
        fmt.Fprintf(os.Stderr, "%s\n", err)
        exitCode = 1
    }
    return exitCode
}

// The 'measure' command helper.
func measureHelp() int {
    fmt.Fprintf(os.Stdout, "use: codeometer measure --src=<file path | zip file path | git repo url | directory path>\n"+
                           "                        [--exts=<extensions> --font-size=<font-size> --measures=<measures>\n" +
                           "                         --stats-per-file --estimatives --async --loaders-nr=<number>]\n")
    return 0
}

// Builds up the measure report. By the way, it does the stuff of the tool.
func measureReport(src string, exts []string, fontSize string, wantedMeasures []string,
                   statsPerFile bool, estimatives bool) (string, error) {
    if len(src) == 0 {
        return "", fmt.Errorf("error: --src option is missing.")
    }

    codestat := &ruler.CodeStat{}

    if fontSize == "12px" {
        codestat.CalibrateCourier12px()
    } else if fontSize == "10px" {
        codestat.CalibrateCourier12px()
    } else {
        return "", fmt.Errorf("error: '%s' font size is invalid. It must be '--10px' or '--12px' (default).", fontSize)
    }

    err := loader.LoadCode(codestat, src, exts...)

    if err != nil {
        return "", fmt.Errorf("error: %s", err)
    }

    measurers := map[string]interface{}{"mi" : &measurer.MICodeStat{},
                                        "km" : &measurer.KMCodeStat{},
                                        "m"  : &measurer.MCodeStat{},
                                        "mm" : &measurer.MMCodeStat{}}

    var info string

    for _, wantedMeasure := range wantedMeasures {
        m, found := measurers[wantedMeasure]
        if !found {
            return "", fmt.Errorf("error: '%s' is a unknown measure. It must be a list containing: " +
                                  "'mi', 'km', 'm' or 'mm'.", wantedMeasure)
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

    if statsPerFile {
        info += "\n"

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
            info += fmt.Sprintf("%s has %s.\n", file, fileInfo)
        }

        info += "\n"
    }

    if estimatives {
        if !statsPerFile {
            info += "\n"
        }

        estimators := []interface{}{
            &estimator.ChineseGreatWall{},
            &estimator.MountEverest{},
            &estimator.PaulistaAvenue{},
            &estimator.ArcDeTriomphe{},
            &estimator.BigBang{},
            &estimator.ChristTheRedeemer{},
            &estimator.Coliseum{},
            &estimator.EiffelTower{},
            &estimator.EmpireStateBuilding{},
            &estimator.IguazuFalls{},
            &estimator.LibertyStatue{},
            &estimator.NiagaraFalls{},
            &estimator.Pantheon{},
            &estimator.SistineChapel{},
            &estimator.WallStreet{},
            &estimator.WashingtonMonument{},
            &estimator.FrogTraveler{},
        }

        for _, e := range estimators {
            switch e.(type) {
                case *estimator.ChineseGreatWall:
                    o := e.(*estimator.ChineseGreatWall)
                    info += o.Estimate(codestat)
                    break

                case *estimator.MountEverest:
                    o := e.(*estimator.MountEverest)
                    info += o.Estimate(codestat)
                    break

                case *estimator.PaulistaAvenue:
                    o := e.(*estimator.PaulistaAvenue)
                    info += o.Estimate(codestat)
                   break

                case *estimator.ArcDeTriomphe:
                    o := e.(*estimator.ArcDeTriomphe)
                    info += o.Estimate(codestat)
                    break

                case *estimator.BigBang:
                    o := e.(*estimator.BigBang)
                    info += o.Estimate(codestat)
                    break

                case *estimator.ChristTheRedeemer:
                    o := e.(*estimator.ChristTheRedeemer)
                    info += o.Estimate(codestat) 
                    break

                case *estimator.Coliseum:
                    o := e.(*estimator.Coliseum)
                    info += o.Estimate(codestat)
                    break

                case *estimator.EiffelTower:
                    o := e.(*estimator.EiffelTower)
                    info += o.Estimate(codestat)
                    break

                case *estimator.EmpireStateBuilding:
                    o := e.(*estimator.EmpireStateBuilding)
                    info += o.Estimate(codestat)
                    break

                case *estimator.IguazuFalls:
                    o := e.(*estimator.IguazuFalls)
                    info += o.Estimate(codestat)
                    break

                case *estimator.LibertyStatue:
                    o := e.(*estimator.LibertyStatue)
                    info += o.Estimate(codestat)
                    break

                case *estimator.NiagaraFalls:
                    o := e.(*estimator.NiagaraFalls)
                    info += o.Estimate(codestat)
                    break

                case *estimator.Pantheon:
                    o := e.(*estimator.Pantheon)
                    info += o.Estimate(codestat)
                    break

                case *estimator.SistineChapel:
                    o := e.(*estimator.SistineChapel)
                    info += o.Estimate(codestat)
                    break

                case *estimator.WallStreet:
                    o := e.(*estimator.WallStreet)
                    info += o.Estimate(codestat)
                    break

                case *estimator.WashingtonMonument:
                    o := e.(*estimator.WashingtonMonument)
                    info += o.Estimate(codestat)
                    break

                case *estimator.FrogTraveler:
                    o := e.(*estimator.FrogTraveler)
                    info += o.Estimate(codestat)
                    break
            }

            info += "\n"
        }
        info += "\n"
    }

    return info, nil
}
