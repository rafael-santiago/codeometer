// package estimator prints estimatives about your code metrics based on famous places, momuments.
// --
package estimator

import (
    "internal/ruler"
    "internal/measurer"
)

const kPaulistaAvenueInKM = 2.8 // Mano do ceu... eh nois!

type PaulistaAvenueEstimator struct {}

// Returns the size of Paulista avenue, mano.
func (p *PaulistaAvenueEstimator) K() float64 {
    return kPaulistaAvenueInKM
}

// Returns a string with some estimative of your code against Paulista avenue size.
func (p *PaulistaAvenueEstimator) Estimate(codestat ruler.CodeStat) string {
    km := &measurer.KMCodeStat{}
    km.Calibrate(codestat)
    return doEstimative(km, "Your code has %.2f%% of the Paulista avenue extension (%.2f km).",
                            "Paulista avenue has %.2f%% of your code extension (%.2f km)", p)
}
