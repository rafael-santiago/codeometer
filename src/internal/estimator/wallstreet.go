// package estimator prints estimatives about your code metrics based on famous places, momuments.
// --
//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package estimator

import (
    "internal/ruler"
    "internal/measurer"
)

const kWallStreetSizeInM = 800

type WallStreetEstimator struct {}

// Returns the size of Wall street.
func (w *WallStreetEstimator) K() float64 {
    return kWallStreetSizeInM
}

// Returns a string with some estimative of your code against Wall street size.
func (w *WallStreetEstimator) Estimate(codestat *ruler.CodeStat) string {
    m := &measurer.MCodeStat{}
    m.Calibrate(codestat)
    return doEstimative(m, "Your code has %.2f%% of the Wall street extension (%.fm).",
                           "Wall street's extension (%.fm) has %.2f%% of your code extension (%.2fm).", w)
}
