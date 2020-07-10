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

const kNiagaraFallsInM = 51

type NiagaraFalls struct {}

// Returns the size of Niagara Falls.
func (n *NiagaraFalls) K() float64 {
    return kNiagaraFallsInM
}

// Returns a string with some estimative of your code against Niagara Falls size.
func (n *NiagaraFalls) Estimate(codestat *ruler.CodeStat) string {
    m := &measurer.MCodeStat{}
    m.Calibrate(codestat)
    return doEstimative(m, "Your code has %.2f%% of the Niagara Falls' height (%.fm).",
                           "Niagara Falls' height (%.fm) has %.2f%% of your code extension (%.2fm).", n)
}
