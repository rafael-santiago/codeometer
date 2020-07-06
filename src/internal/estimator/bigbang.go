// package estimator prints estimatives about your code metrics based on famous places, momuments.
// --
package estimator

import (
    "internal/ruler"
    "internal/measurer"
)

const kBigBangInM = 96

type BigBangEstimator struct {}

// Returns the size of Big Bang.
func (b *BigBangEstimator) K() float64 {
    return kBigBangInM
}

// Returns a string with some estimative of your code against Big Bang size.
func (b *BigBangEstimator) Estimate(codestat ruler.CodeStat) string {
    m := &measurer.MCodeStat{}
    m.Calibrate(codestat)
    return doEstimative(m, "Your code has %.2f%% of the Big Bang's height (%.2f m).",
                           "Big Bang's height has %.2f%% of your code extension (%.2f m)", b)
}
