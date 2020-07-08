//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.package measurer
package measurer

import (
    "testing"
    "internal/magnitudes"
    "internal/ruler"
    "fmt"
)

func TestMICalibrate(t *testing.T) {
    types := []interface{}{
        &ruler.CodeStat{},
        &MMCodeStat{},
        &MCodeStat{},
        &KMCodeStat{},
    }
    mi := &MICodeStat{}
    for _, t := range types {
        mi.Calibrate(t)
    }
}

func TestMIDistancePerLine(t *testing.T) {
    codestat := &ruler.CodeStat{}
    mi := &MICodeStat{}
    mi.Calibrate(codestat)
    k := ((float64(magnitudes.GetA4PaperWidthSizeInMM()) / 1000) / 1000) * 0.621371
    distance := mi.DistancePerLine()
    if distance != k {
        t.Errorf(`distance != k: %v != %v`, distance, k)
    }
}

func TestMIDistancePerPage(t *testing.T) {
    codestat := &ruler.CodeStat{}
    codestat.CalibrateCourier12px()
    k := (((float64(magnitudes.GetA4PaperWidthSizeInMM()) * float64(codestat.CharPerPage)) / 1000) / 1000) * 0.621371
    mi := &MICodeStat{}
    mi.Calibrate(codestat)
    if k != mi.DistancePerPage() {
        t.Errorf(`k != mi.DistancePerPage(): %v != %v`, k, mi.DistancePerPage())
    }
}

func TestMIDistancePerFile(t *testing.T) {
    files := []struct {
        Filepath string
        BytesTotal int64
        ExpectedSize string
    }{
        { "main.c", 101, "0.00"},
        { "sys.c", 727318, "1.64"},
        { "exit.c", 92388, "0.21"},
    }
    codestat := &ruler.CodeStat{}
    codestat.CalibrateCourier12px()
    codestat.Files = make(map[string]ruler.CodeFileInfo)
    for _, file := range files {
        codestat.Files[file.Filepath] = ruler.CodeFileInfo{file.BytesTotal}
    }
    mi := &MICodeStat{}
    mi.Calibrate(codestat)
    for _, file := range files {
        distance := fmt.Sprintf("%.2f", mi.DistancePerFile(file.Filepath))
        if distance != file.ExpectedSize {
            t.Errorf(`distance != file.ExpectedSize: %v != %v`, distance, file.ExpectedSize)
        }
    }
}

func TestMITotalDistance(t *testing.T) {
    files := []struct {
        Filepath string
        BytesTotal int64
    }{
        { "main.c", 101},
        { "sys.c", 727318},
        { "exit.c", 92388},
    }
    expectedTotalDistance := "1.85"
    codestat := &ruler.CodeStat{}
    codestat.CalibrateCourier12px()
    codestat.Files = make(map[string]ruler.CodeFileInfo)
    for _, file := range files {
        codestat.Files[file.Filepath] = ruler.CodeFileInfo{file.BytesTotal}
    }
    mi := &MICodeStat{}
    mi.Calibrate(codestat)
    distance := fmt.Sprintf("%.2f", mi.TotalDistance())
    if distance != expectedTotalDistance {
        t.Errorf(`distance != expectedTotalDistance: %v != %v`, distance, expectedTotalDistance)
    }
}
