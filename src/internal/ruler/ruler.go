//package ruler gathers relevant structures and interfaces for coding measurements.
//--
package ruler

import (
    "sync"
)

type CodeFileInfo struct {
    BytesTotal int64
}

type CodeStat struct {
    sync.Mutex
    Files map[string]CodeFileInfo
}

type CodingRuler interface {
    DistancePerLine() float64
    DistancePerPage() float64
    TotalDistance() float64
}

