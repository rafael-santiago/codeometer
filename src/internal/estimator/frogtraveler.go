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

const kFrogTravelerInMM = 44

type FrogTraveler struct {}

// Returns the size of Frog-Traveler.
func (f *FrogTraveler) K() float64 {
    return kFrogTravelerInMM
}

// Returns a string with some estimative of your code against Frog-Traveler size.
func (f *FrogTraveler) Estimate(codestat *ruler.CodeStat) string {
    mm := &measurer.MMCodeStat{}
    mm.Calibrate(codestat)
    return doEstimative(mm, "Your code has %.2f%% of the Frog-Traveler's height (%.f mm).",
                            "Frog-Traveler's height (%.f mm) has %.2f%% of your code extension (%.2f mm).", f)
}
