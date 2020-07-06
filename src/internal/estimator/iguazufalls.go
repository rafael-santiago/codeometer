// package estimator prints estimatives about your code metrics based on famous places, momuments.
// --
package estimator

import (
    "internal/ruler"
    "internal/measurer"
)

const kIguazuFallsInM = 82

type IguazuFallsEstimator struct {}

// Returns the size of Iguazu Falls.
func (i *IguazuFallsEstimator) K() float64 {
    return kIguazuFallsInM
}

// Returns a string with some estimative of your code against Iguazu Falls size.
func (i *IguazuFallsEstimator) Estimate(codestat ruler.CodeStat) string {
    m := &measurer.MCodeStat{}
    m.Calibrate(codestat)
    return doEstimative(m, "Your code has %.2f%% of the Iguazu Falls' height (%.2f m).",
                           "Iguazu Falls' height has %.2f%% of your code extension (%.2f m)", i)
}
