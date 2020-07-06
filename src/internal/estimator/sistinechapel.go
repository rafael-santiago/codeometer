// package estimator prints estimatives about your code metrics based on famous places, momuments.
// --
package estimator

import (
    "internal/ruler"
    "internal/measurer"
)

const kSistineChapelInM = 21

type SistineChapelEstimator struct {}

// Returns the size of Sistine Chapel.
func (s *SistineChapelEstimator) K() float64 {
    return kSistineChapelInM
}

// Returns a string with some estimative of your code against Sistine Chapel size.
func (s *SistineChapelEstimator) Estimate(codestat ruler.CodeStat) string {
    m := &measurer.MCodeStat{}
    m.Calibrate(codestat)
    return doEstimative(m, "Your code has %.2f%% of the Sistine Chapel's height (%.2f m).",
                           "Sistine Chapel's height has %.2f%% of your code extension (%.2f m)", s)
}
