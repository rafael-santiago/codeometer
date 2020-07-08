// package measurer gathers all codes related to distance evaluation.
// --
//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package measurer

import (
    "internal/ruler"
    "internal/magnitudes"
)

// Expresses CodeStat measures in millimiters.
type MMCodeStat ruler.CodeStat

// Returns in mm the width of a entire filled line.
func (mm *MMCodeStat) DistancePerLine() float64 {
    mm.Lock()
    defer mm.Unlock()
    return float64(magnitudes.GetA4PaperWidthSizeInMM())
}

// Returns in mm the distance of a entire filled page.
func (mm *MMCodeStat) DistancePerPage() float64 {
    return mm.DistancePerLine() * float64(mm.CharPerPage)
}

// Returns the total distance (in mm) of all loaded codes.
func (mm *MMCodeStat) TotalDistance() float64 {
    var mmTotal float64
    for filename, _ := range mm.Files {
        mmTotal += mm.DistancePerFile(filename)
    }
    return  mmTotal
}

// Returns the total distance (in mm) of a specific previous loaded file.
func (mm *MMCodeStat) DistancePerFile(filename string) float64 {
    distancePerLine := mm.DistancePerLine()
    mm.Lock()
    defer mm.Unlock()
    var distance float64
    if fileInfo, ok := mm.Files[filename]; ok {
        distance = (float64(fileInfo.BytesTotal) / float64(mm.CharPerLine)) * float64(distancePerLine)
    }
    return distance
}

// Calibrates basic measurements from a known passed struct.
func (mm *MMCodeStat) Calibrate(data interface{}) {
    mm.Lock()
    defer mm.Unlock()
    switch data.(type) {
        case *ruler.CodeStat:
            mm.calibrateFromCodeStat(data.(*ruler.CodeStat))
            break

        case *MCodeStat:
            mm.calibrateFromMCodeStat(data.(*MCodeStat))
            break

        case *KMCodeStat:
            mm.calibrateFromKMCodeStat(data.(*KMCodeStat))
            break

        case *MICodeStat:
            mm.calibrateFromMICodeStat(data.(*MICodeStat))
            break

        default:
            panic("MMCodeStat.Calibrate(): Unsupported type was passed.")
            break
    }
}

// Calibrates from a *ruler.CodeStat.
func (mm *MMCodeStat) calibrateFromCodeStat(cs *ruler.CodeStat) {
    cs.Lock()
    defer cs.Unlock()
    mm.Files = make(map[string]ruler.CodeFileInfo)
    for k, v := range cs.Files {
        mm.Files[k] = v
    }
    mm.CharPerLine = cs.CharPerLine
    mm.CharPerPage = cs.CharPerPage
}

// Calibrates from a *MCodeStat.
func (mm *MMCodeStat) calibrateFromMCodeStat(m *MCodeStat) {
    m.Lock()
    defer m.Unlock()
    mm.Files = make(map[string]ruler.CodeFileInfo)
    for k, v := range m.Files {
        mm.Files[k] = v
    }
    mm.CharPerLine = m.CharPerLine
    mm.CharPerPage = m.CharPerPage
}

// Calibrates from a *KMCodeStat.
func (mm *MMCodeStat) calibrateFromKMCodeStat(km *KMCodeStat) {
    km.Lock()
    defer km.Unlock()
    mm.Files = make(map[string]ruler.CodeFileInfo)
    for k, v := range km.Files {
        mm.Files[k] = v
    }
    mm.CharPerLine = km.CharPerLine
    mm.CharPerPage = km.CharPerPage
}

// Calibrates from a *MICodeStat.
func (mm *MMCodeStat) calibrateFromMICodeStat(mi *MICodeStat) {
    mi.Lock()
    defer mi.Unlock()
    mm.Files = make(map[string]ruler.CodeFileInfo)
    for k, v := range mi.Files {
        mm.Files[k] = v
    }
    mm.CharPerLine = mi.CharPerLine
    mm.CharPerPage = mi.CharPerPage
}
