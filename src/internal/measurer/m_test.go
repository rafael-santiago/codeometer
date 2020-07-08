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

func TestMCalibrate(t *testing.T) {
    types := []interface{}{
        &ruler.CodeStat{},
        &MMCodeStat{},
        &KMCodeStat{},
        &MICodeStat{},
    }
    m := &MCodeStat{}
    for _, t := range types {
        m.Calibrate(t)
    }
}

func TestMDistancePerLine(t *testing.T) {
    codestat := &ruler.CodeStat{}
    m := &MCodeStat{}
    m.Calibrate(codestat)
    k := float64(magnitudes.GetA4PaperWidthSizeInMM()) / 1000
    distance := m.DistancePerLine()
    if distance != k {
        t.Errorf(`distance != k: %v != %v`, distance, k)
    }
}

func TestMDistancePerPage(t *testing.T) {
    codestat := &ruler.CodeStat{}
    codestat.CalibrateCourier12px()
    k := (float64(magnitudes.GetA4PaperWidthSizeInMM()) * float64(codestat.CharPerPage)) / 1000
    m := &MCodeStat{}
    m.Calibrate(codestat)
    if k != m.DistancePerPage() {
        t.Errorf(`k != mm.DistancePerPage(): %v != %v`, k, m.DistancePerPage())
    }
}

func TestMDistancePerFile(t *testing.T) {
    files := []struct {
        Filepath string
        BytesTotal int64
        ExpectedSize string
    }{
        { "main.c", 101, "0.37"},
        { "sys.c", 727318, "2634.31"},
        { "exit.c", 92388, "334.62"},
    }
    codestat := &ruler.CodeStat{}
    codestat.CalibrateCourier12px()
    codestat.Files = make(map[string]ruler.CodeFileInfo)
    for _, file := range files {
        codestat.Files[file.Filepath] = ruler.CodeFileInfo{file.BytesTotal}
    }
    m := &MCodeStat{}
    m.Calibrate(codestat)
    for _, file := range files {
        distance := fmt.Sprintf("%.2f", m.DistancePerFile(file.Filepath))
        if distance != file.ExpectedSize {
            t.Errorf(`distance != file.ExpectedSize: %v != %v`, distance, file.ExpectedSize)
        }
    }
}

func TestMTotalDistance(t *testing.T) {
    files := []struct {
        Filepath string
        BytesTotal int64
    }{
        { "main.c", 101},
        { "sys.c", 727318},
        { "exit.c", 92388},
    }
    expectedTotalDistance := "2969.30"
    codestat := &ruler.CodeStat{}
    codestat.CalibrateCourier12px()
    codestat.Files = make(map[string]ruler.CodeFileInfo)
    for _, file := range files {
        codestat.Files[file.Filepath] = ruler.CodeFileInfo{file.BytesTotal}
    }
    m := &MCodeStat{}
    m.Calibrate(codestat)
    distance := fmt.Sprintf("%.2f", m.TotalDistance())
    if distance != expectedTotalDistance {
        t.Errorf(`distance != expectedTotalDistance: %v != %v`, distance, expectedTotalDistance)
    }
}
