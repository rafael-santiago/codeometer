// package estimator prints estimatives about your code metrics based on famous places, momuments.
// --
package estimator

import (
    "internal/ruler"
    "internal/measurer"
)

const kColiseumInM = 48

type ColiseumEstimator struct {}

// Returns the size of Coliseum.
func (c *ColiseumEstimator) K() float64 {
    return kColiseumInM
}

// Returns a string with some estimative of your code against Coliseum size.
func (c *ColiseumEstimator) Estimate(codestat ruler.CodeStat) string {
    m := &measurer.MCodeStat{}
    m.Calibrate(codestat)
    return doEstimative(m, "Your code has %.2f%% of the Coliseum's height (%.2f m).",
                           "Coliseum's height has %.2f%% of your code extension (%.2f m)", c)
}
