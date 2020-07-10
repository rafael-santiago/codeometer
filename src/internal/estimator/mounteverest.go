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

const kMountEverestInM = 8848

type MountEverest struct {}

// Returns the size of Mount Everest.
func (m *MountEverest) K() float64 {
    return kMountEverestInM
}

// Returns a string with some estimative of your code against Mount Everest size.
func (m *MountEverest) Estimate(codestat *ruler.CodeStat) string {
    meter := &measurer.MCodeStat{}
    meter.Calibrate(codestat)
    return doEstimative(meter, "Your code has %.2f%% of the Mount Everest's height (%.fm).",
                               "Mount Everest's height (%.fm) has %.2f%% of your code extension (%.2fm).", m)
}
