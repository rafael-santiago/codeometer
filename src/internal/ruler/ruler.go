//package ruler gathers relevant structures and interfaces for coding measurements.
//--
//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
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

