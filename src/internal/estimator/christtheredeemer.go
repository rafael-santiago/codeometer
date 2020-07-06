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

const kChristTheRedmeerInM = 38

type ChristTheRedeemerEstimator struct {}

// Returns the size of Christ the Redeemer.
func (c *ChristTheRedeemerEstimator) K() float64 {
    return kChristTheRedmeerInM
}

// Returns a string with some estimative of your code against Christ the Redmeer size.
func (c *ChristTheRedeemerEstimator) Estimate(codestat ruler.CodeStat) string {
    m := &measurer.MCodeStat{}
    m.Calibrate(codestat)
    return doEstimative(m, "Your code has %.2f%% of the Christ the Redeemer's height (%.2f m).",
                           "Christ the Redeemer's height has %.2f%% of your code extension (%.2f m)", c)
}
