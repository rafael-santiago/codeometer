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

const kEmpireStateBuildingInM = 381

type EmpireStateBuilding struct {}

// Returns the size of Empire State Building.
func (e *EmpireStateBuilding) K() float64 {
    return kEmpireStateBuildingInM
}

// Returns a string with some estimative of your code against Empire State Building size.
func (e *EmpireStateBuilding) Estimate(codestat *ruler.CodeStat) string {
    m := &measurer.MCodeStat{}
    m.Calibrate(codestat)
    return doEstimative(m, "Your code has %.2f%% of the Empire State Building's height (%.fm).",
                           "Empire State Building's height (%.fm) has %.2f%% of your code extension (%.2fm).", e)
}
