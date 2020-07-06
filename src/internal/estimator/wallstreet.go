// package estimator prints estimatives about your code metrics based on famous places, momuments.
// --
package estimator

import (
    "internal/ruler"
    "internal/measurer"
)

const kWallStreetSizeInM = 800

type WallStreetEstimator struct {}

// Returns the size of Wall street.
func (w *WallStreetEstimator) K() float64 {
    return kWallStreetSizeInM
}

// Returns a string with some estimative of your code against Wall street size.
func (w *WallStreetEstimator) Estimate(codestat ruler.CodeStat) string {
    m := &measurer.MCodeStat{}
    m.Calibrate(codestat)
    return doEstimative(m, "Your code has %.2f%% of the Wall street extension (%.2f m).",
                           "Wall street has %.2f%% of your code extension (%.2f m)", w)
}
