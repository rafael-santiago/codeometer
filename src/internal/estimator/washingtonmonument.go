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

const kWashingtonMonumentInM = 169

type WashingtonMonument struct {}

// Returns the size of Washington Monument.
func (w *WashingtonMonument) K() float64 {
    return kWashingtonMonumentInM
}

// Returns a string with some estimative of your code against Washington Monument size.
func (w *WashingtonMonument) Estimate(codestat *ruler.CodeStat) string {
    m := &measurer.MCodeStat{}
    m.Calibrate(codestat)
    return doEstimative(m, "%s has %.2f%% of the Washington Monument's height (%.fm).",
                           "Washington Monument's height (%.fm) has %.2f%% of %s extension (%.2fm).", w)
}
