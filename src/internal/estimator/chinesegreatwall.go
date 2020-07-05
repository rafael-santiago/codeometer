// package estimator prints estimatives about your code metrics based on famous places, momuments.
// --
package estimator

import (
    "internal/ruler"
    "internal/measurer"
)

const kChineseGreatWallSizeInKM = 21196

type ChineseGreatWallEstimator struct {}

// Returns the size of Chinese great wall.
func (c *ChineseGreatWallEstimator) K() float64 {
    return kChineseGreatWallSizeInKM
}

// Returns a string with some estimative of your code against Chinese great wall size.
func (c *ChineseGreatWallEstimator) Estimate(codestat ruler.CodeStat) string {
    km := &measurer.KMCodeStat{}
    km.Calibrate(codestat)
    return doEstimative(km, "Your code has %.2f%% of the Chinese great wall extension (%d km).",
                            "Chinese great wall has %.2f%% of your code extension (%.2f km)", c)
}
