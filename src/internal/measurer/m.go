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
type MCodeStat ruler.CodeStat

// Calibrates data from a known struct.
func (m *MCodeStat) Calibrate(data interface{}) {
    m.Lock()
    defer m.Unlock()
    switch data.(type) {
        case *ruler.CodeStat:
            m.calibrateFromCodeStat(data.(*ruler.CodeStat))
            break

        case *MMCodeStat:
            m.calibrateFromMMCodeStat(data.(*MMCodeStat))
            break

        case *KMCodeStat:
            m.calibrateFromKMCodeStat(data.(*KMCodeStat))
            break

        case *MICodeStat:
            m.calibrateFromMICodeStat(data.(*MICodeStat))
            break

        default:
            panic("MCodeStat.Calibrate(): Unsupported type was passed.")
            break
    }
}

// Calibrates from a *ruler.CodeStat.
func (m *MCodeStat) calibrateFromCodeStat(cs *ruler.CodeStat) {
    cs.Lock()
    defer cs.Unlock()
    m.Files = make(map[string]ruler.CodeFileInfo)
    for k, v := range cs.Files {
        m.Files[k] = v
    }
    m.CharPerLine = cs.CharPerLine
    m.CharPerPage = cs.CharPerPage
    m.ProjectName = cs.ProjectName
}

// Calibrates from a *MCodeStat.
func (m *MCodeStat) calibrateFromMMCodeStat(mm *MMCodeStat) {
    mm.Lock()
    defer mm.Unlock()
    m.Files = make(map[string]ruler.CodeFileInfo)
    for k, v := range mm.Files {
        m.Files[k] = v
    }
    m.CharPerLine = mm.CharPerLine
    m.CharPerPage = mm.CharPerPage
    m.ProjectName = mm.ProjectName
}

// Calibrates from a *KMCodeStat.
func (m *MCodeStat) calibrateFromKMCodeStat(km *KMCodeStat) {
    km.Lock()
    defer km.Unlock()
    m.Files = make(map[string]ruler.CodeFileInfo)
    for k, v := range km.Files {
        m.Files[k] = v
    }
    m.CharPerLine = km.CharPerLine
    m.CharPerPage = km.CharPerPage
    m.ProjectName = km.ProjectName
}

// Calibrates from a *MICodeStat.
func (m *MCodeStat) calibrateFromMICodeStat(mi *MICodeStat) {
    mi.Lock()
    defer mi.Unlock()
    m.Files = make(map[string]ruler.CodeFileInfo)
    for k, v := range mi.Files {
        m.Files[k] = v
    }
    m.CharPerLine = mi.CharPerLine
    m.CharPerPage = mi.CharPerPage
    m.ProjectName = mi.ProjectName
}

// Returns in m the width of a entire filled line.
func (m *MCodeStat) DistancePerLine() float64 {
    mm := &MMCodeStat{}
    mm.Calibrate(m)
    m.Lock()
    defer m.Unlock()
    return mm.DistancePerLine() / 1000
}

// Returns in m the distance of a entire filled page.
func (m *MCodeStat) DistancePerPage() float64 {
    mm := &MMCodeStat{}
    mm.Calibrate(m)
    m.Lock()
    defer m.Unlock()
    return mm.DistancePerPage() / 1000
}

// Returns the total distance (in m) of all loaded codes.
func (m *MCodeStat) TotalDistance() float64 {
    mm := &MMCodeStat{}
    mm.Calibrate(m)
    m.Lock()
    defer m.Unlock()
    return mm.TotalDistance() / 1000
}

// Returns the total distance (in m) of a specific previous loaded file.
func (m *MCodeStat) DistancePerFile(filename string) float64 {
    mm := &MMCodeStat{}
    mm.Calibrate(m)
    m.Lock()
    defer m.Unlock()
    return mm.DistancePerFile(filename) / 1000
}
