// package estimator prints estimatives about your code metrics based on famous places, momuments.
// --
package estimator

import (
    "internal/ruler"
    "internal/measurer"
)

const kNiagaraFallsInM = 51

type NiagaraFallsEstimator struct {}

// Returns the size of Niagara Falls.
func (n *NiagaraFallsEstimator) K() float64 {
    return kNiagaraFallsInM
}

// Returns a string with some estimative of your code against Niagara Falls size.
func (n *NiagaraFallsEstimator) Estimate(codestat ruler.CodeStat) string {
    m := &measurer.MCodeStat{}
    m.Calibrate(codestat)
    return doEstimative(m, "Your code has %.2f%% of the Niagara Falls' height (%.2f m).",
                           "Niagara Falls' height has %.2f%% of your code extension (%.2f m)", n)
}
