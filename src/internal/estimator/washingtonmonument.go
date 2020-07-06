// package estimator prints estimatives about your code metrics based on famous places, momuments.
// --
package estimator

import (
    "internal/ruler"
    "internal/measurer"
)

const kWashingtonMonumentInM = 169

type WashingtonMonumentEstimator struct {}

// Returns the size of Washington Monument.
func (w *WashingtonMonumentEstimator) K() float64 {
    return kWashingtonMonumentInM
}

// Returns a string with some estimative of your code against Washington Monument size.
func (w *WashingtonMonumentEstimator) Estimate(codestat ruler.CodeStat) string {
    m := &measurer.MCodeStat{}
    m.Calibrate(codestat)
    return doEstimative(m, "Your code has %.2f%% of the Washington Monument's height (%.2f m).",
                           "Washington Monument's height has %.2f%% of your code extension (%.2f m)", w)
}
