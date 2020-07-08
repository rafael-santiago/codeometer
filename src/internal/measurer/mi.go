// package measurer gathers all codes related to distance evaluation.
// --
//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package measurer

import (
    "internal/ruler"
)

const k1KMPerMI = 0.621371

// Expresses CodeStat measures in miles.
type MICodeStat ruler.CodeStat

// Calibrates data from a known struct.
func (mi *MICodeStat) Calibrate(data interface{}) {
    mi.Lock()
    defer mi.Unlock()
    switch data.(type) {
        case *ruler.CodeStat:
            mi.calibrateFromCodeStat(data.(*ruler.CodeStat))
            break

        case *MCodeStat:
            mi.calibrateFromMCodeStat(data.(*MCodeStat))
            break

        case *MMCodeStat:
            mi.calibrateFromMMCodeStat(data.(*MMCodeStat))
            break

        case *KMCodeStat:
            mi.calibrateFromKMCodeStat(data.(*KMCodeStat))
            break

        default:
            panic("MICodeStat.Calibrate(): Unsupported type was passed.")
            break
    }
}

// Calibrates from a *ruler.CodeStat.
func (mi *MICodeStat) calibrateFromCodeStat(cs *ruler.CodeStat) {
    cs.Lock()
    defer cs.Unlock()
    mi.Files = make(map[string]ruler.CodeFileInfo)
    for k, v := range cs.Files {
        mi.Files[k] = v
    }
    mi.CharPerLine = cs.CharPerLine
    mi.CharPerPage = cs.CharPerPage
}

// Calibrates from a *MCodeStat.
func (mi *MICodeStat) calibrateFromMCodeStat(m *MCodeStat) {
    m.Lock()
    defer m.Unlock()
    mi.Files = make(map[string]ruler.CodeFileInfo)
    for k, v := range m.Files {
        mi.Files[k] = v
    }
    mi.CharPerLine = m.CharPerLine
    mi.CharPerPage = m.CharPerPage
}

// Calibrates from a *MMCodeStat.
func (mi *MICodeStat) calibrateFromMMCodeStat(mm *MMCodeStat) {
    mm.Lock()
    defer mm.Unlock()
    mi.Files = make(map[string]ruler.CodeFileInfo)
    for k, v := range mm.Files {
        mi.Files[k] = v
    }
    mi.CharPerLine = mm.CharPerLine
    mi.CharPerPage = mm.CharPerPage
}

// Calibrates from a *KMCodeStat.
func (mi *MICodeStat) calibrateFromKMCodeStat(km *KMCodeStat) {
    km.Lock()
    defer km.Unlock()
    mi.Files = make(map[string]ruler.CodeFileInfo)
    for k, v := range km.Files {
        mi.Files[k] = v
    }
    mi.CharPerLine = km.CharPerLine
    mi.CharPerPage = km.CharPerPage
}

// Returns in mi the width of a entire filled line.
func (mi *MICodeStat) DistancePerLine() float64 {
    km := &KMCodeStat{}
    km.Calibrate(mi)
    mi.Lock()
    defer mi.Unlock()
    return km.DistancePerLine() * k1KMPerMI
}

// Returns in mi the distance of a entire filled page.
func (mi *MICodeStat) DistancePerPage() float64 {
    km := &KMCodeStat{}
    km.Calibrate(mi)
    mi.Lock()
    defer mi.Unlock()
    return km.DistancePerPage() * k1KMPerMI
}

// Returns the total distance (in mi) of all loaded codes.
func (mi *MICodeStat) TotalDistance() float64 {
    km := &KMCodeStat{}
    km.Calibrate(mi)
    mi.Lock()
    defer mi.Unlock()
    return km.TotalDistance() * k1KMPerMI
}

// Returns the total distance (in mi) of a specific previous loaded file.
func (mi *MICodeStat) DistancePerFile(filename string) float64 {
    km := &KMCodeStat{}
    km.Calibrate(mi)
    mi.Lock()
    defer mi.Unlock()
    return km.DistancePerFile(filename) * k1KMPerMI
}
