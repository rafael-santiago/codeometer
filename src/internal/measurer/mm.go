// package measurer gathers all codes related to distance evaluation.
// --
package measurer

import (
    "internal/ruler"
    "internal/magnitudes"
)

type MMCodeStat ruler.CodeStat

// Returns in mm the width of a entire filled line.
func (mm *MMCodeStat) DistancePerLine() float64 {
    mm.Lock()
    defer mm.Unlock()
    return float64(magnitudes.GetA4PaperWidthSizeInMM())
}

// Returns in mm the distance of a entire filled page.
func (mm *MMCodeStat) DistancePerPage() float64 {
    mm.Lock()
    defer mm.Unlock()
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
