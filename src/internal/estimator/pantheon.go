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

const kPantheonInM = 14

type Pantheon struct {}

// Returns the size of Pantheon.
func (p *Pantheon) K() float64 {
    return kPantheonInM
}

// Returns a string with some estimative of your code against Pantheon size.
func (p *Pantheon) Estimate(codestat *ruler.CodeStat) string {
    m := &measurer.MCodeStat{}
    m.Calibrate(codestat)
    return doEstimative(m, "Your code has %.2f%% of the Pantheon's height (%.fm).",
                           "Pantheon's height (%.fm) has %.2f%% of your code extension (%.2fm).", p)
}
