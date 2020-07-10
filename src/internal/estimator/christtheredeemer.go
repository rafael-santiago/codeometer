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

type ChristTheRedeemer struct {}

// Returns the size of Christ the Redeemer.
func (c *ChristTheRedeemer) K() float64 {
    return kChristTheRedmeerInM
}

// Returns a string with some estimative of your code against Christ the Redmeer size.
func (c *ChristTheRedeemer) Estimate(codestat *ruler.CodeStat) string {
    m := &measurer.MCodeStat{}
    m.Calibrate(codestat)
    return doEstimative(m, "%s has %.2f%% of the Christ the Redeemer's height (%.fm).",
                           "Christ the Redeemer's height (%.fm) has %.2f%% of %s extension (%.2fm).", c)
}
