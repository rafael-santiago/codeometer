// package estimator prints estimatives about your code metrics based on famous places, momuments.
// --
package estimator

import (
    "internal/ruler"
    "internal/measurer"
)

const kPantheonInM = 14

type PantheonEstimator struct {}

// Returns the size of Pantheon.
func (p *PantheonEstimator) K() float64 {
    return kPantheonInM
}

// Returns a string with some estimative of your code against Pantheon size.
func (p *PantheonEstimator) Estimate(codestat ruler.CodeStat) string {
    m := &measurer.MCodeStat{}
    m.Calibrate(codestat)
    return doEstimative(m, "Your code has %.2f%% of the Pantheon's height (%.2f m).",
                           "Pantheon's height has %.2f%% of your code extension (%.2f m)", p)
}
