// package estimator prints estimatives about your code metrics based on famous places, momuments.
// --
package estimator

import (
    "internal/ruler"
    "internal/measurer"
)

const kFrogTravelerInMM = 44

type FrogTravelerEstimator struct {}

// Returns the size of Frog-Traveler.
func (f *FrogTravelerEstimator) K() float64 {
    return kFrogTravelerInMM
}

// Returns a string with some estimative of your code against Frog-Traveler size.
func (f *FrogTravelerEstimator) Estimate(codestat ruler.CodeStat) string {
    mm := &measurer.MMCodeStat{}
    mm.Calibrate(codestat)
    return doEstimative(mm, "Your code has %.2f%% of the Frog-Traveler's height (%.2f m).",
                            "Frog-Traveler's height has %.2f%% of your code extension (%.2f m)", f)
}
