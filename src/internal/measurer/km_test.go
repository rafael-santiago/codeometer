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

func TestKMCalibrate(t *testing.T) {
    types := []interface{}{
        &ruler.CodeStat{},
        &MMCodeStat{},
        &MCodeStat{},
        &MICodeStat{},
    }
    km := &KMCodeStat{}
    for _, t := range types {
        km.Calibrate(t)
    }
}

func TestKMDistancePerLine(t *testing.T) {
    codestat := &ruler.CodeStat{}
    km := &KMCodeStat{}
    km.Calibrate(codestat)
    k := (float64(magnitudes.GetA4PaperWidthSizeInMM()) / 1000) / 1000
    distance := km.DistancePerLine()
    if distance != k {
        t.Errorf(`distance != k: %v != %v`, distance, k)
    }
}

func TestKMDistancePerPage(t *testing.T) {
    codestat := &ruler.CodeStat{}
    codestat.CalibrateCourier12px()
    k := ((float64(magnitudes.GetA4PaperWidthSizeInMM()) * float64(codestat.CharPerPage)) / 1000) / 1000
    km := &KMCodeStat{}
    km.Calibrate(codestat)
    if k != km.DistancePerPage() {
        t.Errorf(`k != km.DistancePerPage(): %v != %v`, k, km.DistancePerPage())
    }
}

func TestKMDistancePerFile(t *testing.T) {
    files := []struct {
        Filepath string
        BytesTotal int64
        ExpectedSize string
    }{
        { "main.c", 101, "0.00"},
        { "sys.c", 727318, "2.63"},
        { "exit.c", 92388, "0.33"},
    }
    codestat := &ruler.CodeStat{}
    codestat.CalibrateCourier12px()
    codestat.Files = make(map[string]ruler.CodeFileInfo)
    for _, file := range files {
        codestat.Files[file.Filepath] = ruler.CodeFileInfo{file.BytesTotal}
    }
    km := &KMCodeStat{}
    km.Calibrate(codestat)
    for _, file := range files {
        distance := fmt.Sprintf("%.2f", km.DistancePerFile(file.Filepath))
        if distance != file.ExpectedSize {
            t.Errorf(`distance != file.ExpectedSize: %v != %v`, distance, file.ExpectedSize)
        }
    }
}

func TestKMTotalDistance(t *testing.T) {
    files := []struct {
        Filepath string
        BytesTotal int64
    }{
        { "main.c", 101},
        { "sys.c", 727318},
        { "exit.c", 92388},
    }
    expectedTotalDistance := "2.97"
    codestat := &ruler.CodeStat{}
    codestat.CalibrateCourier12px()
    codestat.Files = make(map[string]ruler.CodeFileInfo)
    for _, file := range files {
        codestat.Files[file.Filepath] = ruler.CodeFileInfo{file.BytesTotal}
    }
    km := &KMCodeStat{}
    km.Calibrate(codestat)
    distance := fmt.Sprintf("%.2f", km.TotalDistance())
    if distance != expectedTotalDistance {
        t.Errorf(`distance != expectedTotalDistance: %v != %v`, distance, expectedTotalDistance)
    }
}
