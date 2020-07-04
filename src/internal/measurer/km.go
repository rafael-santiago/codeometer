// package measurer gathers all codes related to distance evaluation.
// --
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

// Returns in m the width of a entire filled line.
func (km *KMCodeStat) DistancePerLine() float64 {
    km.Lock()
    defer km.Unlock()
    m := &MCodeStat{}
    m.Calibrate(km)
    return m.DistancePerLine() / 1000
}

// Returns in m the distance of a entire filled page.
func (km *KMCodeStat) DistancePerPage() float64 {
    km.Lock()
    defer km.Unlock()
    m := &MCodeStat{}
    m.Calibrate(km)
    return m.DistancePerPage() / 1000
}

// Returns the total distance (in m) of all loaded codes.
func (km *KMCodeStat) TotalDistance() float64 {
    km.Lock()
    defer km.Unlock()
    m := &MCodeStat{}
    m.Calibrate(km)
    return m.TotalDistance() / 1000
}

// Returns the total distance (in m) of a specific previous loaded file.
func (km *KMCodeStat) DistancePerFile(filename string) float64 {
    km.Lock()
    defer km.Unlock()
    m := &MCodeStat{}
    m.Calibrate(km)
    return m.DistancePerFile(filename) / 1000
}
