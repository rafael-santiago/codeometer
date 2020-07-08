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

// Expresses CodeStat measures in meters.
type KMCodeStat ruler.CodeStat

// Calibrates data from a known struct.
func (km *KMCodeStat) Calibrate(data interface{}) {
    km.Lock()
    defer km.Unlock()
    switch data.(type) {
        case *ruler.CodeStat:
            km.calibrateFromCodeStat(data.(*ruler.CodeStat))
            break

        case *MCodeStat:
            km.calibrateFromMCodeStat(data.(*MCodeStat))
            break

        case *MMCodeStat:
            km.calibrateFromMMCodeStat(data.(*MMCodeStat))
            break

        case *MICodeStat:
            km.calibrateFromMICodeStat(data.(*MICodeStat))
            break

        default:
            panic("KMCodeStat.Calibrate(): Unsupported type was passed.")
            break
    }
}

// Calibrates from a *ruler.CodeStat.
func (km *KMCodeStat) calibrateFromCodeStat(cs *ruler.CodeStat) {
    cs.Lock()
    defer cs.Unlock()
    km.Files = make(map[string]ruler.CodeFileInfo)
    for k, v := range cs.Files {
        km.Files[k] = v
    }
    km.CharPerLine = cs.CharPerLine
    km.CharPerPage = cs.CharPerPage
}

// Calibrates from a *MCodeStat.
func (km *KMCodeStat) calibrateFromMCodeStat(m *MCodeStat) {
    m.Lock()
    defer m.Unlock()
    km.Files = make(map[string]ruler.CodeFileInfo)
    for k, v := range m.Files {
        km.Files[k] = v
    }
    km.CharPerLine = m.CharPerLine
    km.CharPerPage = m.CharPerPage
}

// Calibrates from a *MMCodeStat.
func (km *KMCodeStat) calibrateFromMMCodeStat(mm *MMCodeStat) {
    mm.Lock()
    defer mm.Unlock()
    km.Files = make(map[string]ruler.CodeFileInfo)
    for k, v := range mm.Files {
        km.Files[k] = v
    }
    km.CharPerLine = mm.CharPerLine
    km.CharPerPage = mm.CharPerPage
}

// Calibrates from a *MICodeStat.
func (km *KMCodeStat) calibrateFromMICodeStat(mi *MICodeStat) {
    mi.Lock()
    defer mi.Unlock()
    km.Files = make(map[string]ruler.CodeFileInfo)
    for k, v := range mi.Files {
        km.Files[k] = v
    }
    km.CharPerLine = mi.CharPerLine
    km.CharPerPage = mi.CharPerPage
}


// Returns in km the width of a entire filled line.
func (km *KMCodeStat) DistancePerLine() float64 {
    m := &MCodeStat{}
    m.Calibrate(km)
    km.Lock()
    defer km.Unlock()
    return m.DistancePerLine() / 1000
}

// Returns in km the distance of a entire filled page.
func (km *KMCodeStat) DistancePerPage() float64 {
    m := &MCodeStat{}
    m.Calibrate(km)
    km.Lock()
    defer km.Unlock()
    return m.DistancePerPage() / 1000
}

// Returns the total distance (in km) of all loaded codes.
func (km *KMCodeStat) TotalDistance() float64 {
    m := &MCodeStat{}
    m.Calibrate(km)
    km.Lock()
    defer km.Unlock()
    return m.TotalDistance() / 1000
}

// Returns the total distance (in km) of a specific previous loaded file.
func (km *KMCodeStat) DistancePerFile(filename string) float64 {
    m := &MCodeStat{}
    m.Calibrate(km)
    km.Lock()
    defer km.Unlock()
    return m.DistancePerFile(filename) / 1000
}
