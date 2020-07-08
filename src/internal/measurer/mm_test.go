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

func TestCalibrate(t *testing.T) {
    types := []interface{}{
        &ruler.CodeStat{},
        &MCodeStat{},
        &KMCodeStat{},
        &MICodeStat{},
    }
    mm := &MMCodeStat{}
    for _, t := range types {
        mm.Calibrate(t)
    }
}

func TestDistancePerLine(t *testing.T) {
    codestat := &ruler.CodeStat{}
    mm := &MMCodeStat{}
    mm.Calibrate(codestat)
    k := float64(magnitudes.GetA4PaperWidthSizeInMM())
    distance := mm.DistancePerLine()
    if distance != k {
        t.Errorf(`distance != k: %v != %v`, distance, k)
    }
}

func TestDistancePerPage(t *testing.T) {
    codestat := &ruler.CodeStat{}
    codestat.CalibrateCourier12px()
    k := float64(magnitudes.GetA4PaperWidthSizeInMM()) * float64(codestat.CharPerPage)
    mm := &MMCodeStat{}
    mm.Calibrate(codestat)
    if k != mm.DistancePerPage() {
        t.Errorf(`k != mm.DistancePerPage(): %v != %v`, k, mm.DistancePerPage())
    }
}

func TestDistancePerFile(t *testing.T) {
    files := []struct {
        Filepath string
        BytesTotal int64
        ExpectedSize string
    }{
        { "main.c", 101, "365.82"},
        { "sys.c", 727318, "2634310.32"},
        { "exit.c", 92388, "334624.83"},
    }
    codestat := &ruler.CodeStat{}
    codestat.CalibrateCourier12px()
    codestat.Files = make(map[string]ruler.CodeFileInfo)
    for _, file := range files {
        codestat.Files[file.Filepath] = ruler.CodeFileInfo{file.BytesTotal}
    }
    mm := &MMCodeStat{}
    mm.Calibrate(codestat)
    for _, file := range files {
        distance := fmt.Sprintf("%.2f", mm.DistancePerFile(file.Filepath))
        if distance != file.ExpectedSize {
            t.Errorf(`distance != file.ExpectedSize: %v != %v`, distance, file.ExpectedSize)
        }
    }
}

func TestTotalDistance(t *testing.T) {
    files := []struct {
        Filepath string
        BytesTotal int64
        ExpectedSize float64
    }{
        { "main.c", 101, 365.82},
        { "sys.c", 727318, 2634310.32},
        { "exit.c", 92388, 334624.83},
    }
    expectedTotalDistance := "2969300.96"
    codestat := &ruler.CodeStat{}
    codestat.CalibrateCourier12px()
    codestat.Files = make(map[string]ruler.CodeFileInfo)
    for _, file := range files {
        codestat.Files[file.Filepath] = ruler.CodeFileInfo{file.BytesTotal}
    }
    mm := &MMCodeStat{}
    mm.Calibrate(codestat)
    distance := fmt.Sprintf("%.2f", mm.TotalDistance())
    if distance != expectedTotalDistance {
        t.Errorf(`distance != expectedTotalDistance: %v != %v`, distance, expectedTotalDistance)
    }
}
