// package estimator prints estimatives about your code metrics based on famous places, momuments.
// --
package estimator

import (
    "internal/ruler"
    "internal/measurer"
)

const kArcDeTriompheInM = 50

type ArcDeTriompheEstimator struct {}

// Returns the size of Arc de Triomphe.
func (a *ArcDeTriompheEstimator) K() float64 {
    return kArcDeTriompheInM
}

// Returns a string with some estimative of your code against Arc de Triomphe size.
func (a *ArcDeTriompheEstimator) Estimate(codestat ruler.CodeStat) string {
    m := &measurer.MCodeStat{}
    m.Calibrate(codestat)
    return doEstimative(m, "Your code has %.2f%% of the Arc de Triomphe's height (%.2f m).",
                           "Arc de Triomphe's height has %.2f%% of your code extension (%.2f m)", a)
}
