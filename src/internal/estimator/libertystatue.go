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

const kLibertyStatueInM = 93

type LibertyStatue struct {}

// Returns the size of Liberty Statue.
func (l *LibertyStatue) K() float64 {
    return kLibertyStatueInM
}

// Returns a string with some estimative of your code against Liberty Statue size.
func (l *LibertyStatue) Estimate(codestat *ruler.CodeStat) string {
    m := &measurer.MCodeStat{}
    m.Calibrate(codestat)
    return doEstimative(m, "%s has %.2f%% of the Liberty Statue's height (%.fm).",
                           "Liberty Statue's height (%.fm) has %.2f%% of %s extension (%.2fm).", l)
}
