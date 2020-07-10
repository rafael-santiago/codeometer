// package estimator prints estimatives about your code metrics based on famous places, momuments.
// --
//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package estimator

import (
    "internal/ruler"
    "internal/measurer"
)

const kChineseGreatWallSizeInKM = 21196

type ChineseGreatWall struct {}

// Returns the size of Chinese great wall.
func (c *ChineseGreatWall) K() float64 {
    return kChineseGreatWallSizeInKM
}

// Returns a string with some estimative of your code against Chinese great wall size.
func (c *ChineseGreatWall) Estimate(codestat *ruler.CodeStat) string {
    km := &measurer.KMCodeStat{}
    km.Calibrate(codestat)
    return doEstimative(km, "%s has %.2f%% of the Chinese great wall extension (%.f km).",
                            "Chinese great wall's extension (%.f km) has %.2f%% of %s extension (%.2f km).", c)
}
