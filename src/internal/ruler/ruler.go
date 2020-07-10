//package ruler gathers relevant structures and interfaces for coding measurements.
//--
//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package ruler

import (
    "sync"
    "internal/magnitudes"
)

type CodeFileInfo struct {
    BytesTotal int64
}

type CodeStat struct {
    sync.Mutex
    Files map[string]CodeFileInfo
    CharPerLine int64
    CharPerPage int64
    ProjectName string
}

type CodingRuler interface {
    DistancePerLine() float64
    DistancePerPage() float64
    TotalDistance() float64
    DistancePerFile(filename string) float64
}

func (c *CodeStat) CalibrateCourier12px() {
    c.Lock()
    defer c.Unlock()
    c.CharPerLine, c.CharPerPage = magnitudes.GetCourier12pxParams()
}

func (c *CodeStat) CalibrateCourier10px() {
    c.Lock()
    defer c.Unlock()
    c.CharPerLine, c.CharPerPage = magnitudes.GetCourier10pxParams()
}
