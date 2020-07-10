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

const kPaulistaAvenueInKM = 2.8 // Mano do ceu... eh nois!

type PaulistaAvenue struct {}

// Returns the size of Paulista avenue, mano.
func (p *PaulistaAvenue) K() float64 {
    return kPaulistaAvenueInKM
}

// Returns a string with some estimative of your code against Paulista avenue size.
func (p *PaulistaAvenue) Estimate(codestat *ruler.CodeStat) string {
    km := &measurer.KMCodeStat{}
    km.Calibrate(codestat)
    return doEstimative(km, "Your code has %.2f%% of the Paulista avenue extension (%.1f km).",
                            "Paulista avenue's extension (%.1f km) has %.2f%% of your code extension (%.2f km).", p)
}
