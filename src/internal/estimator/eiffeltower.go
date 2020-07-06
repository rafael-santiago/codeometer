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

const kEiffelTowerInM = 300

type EiffelTowerEstimator struct {}

// Returns the size of Eiffel tower.
func (e *EiffelTowerEstimator) K() float64 {
    return kEiffelTowerInM
}

// Returns a string with some estimative of your code against Eifferl tower size.
func (e *EiffelTowerEstimator) Estimate(codestat ruler.CodeStat) string {
    m := &measurer.MCodeStat{}
    m.Calibrate(codestat)
    return doEstimative(m, "Your code has %.2f%% of the Eiffel tower's height (%.2f m).",
                           "Eiffel tower's height has %.2f%% of your code extension (%.2f m)", e)
}
