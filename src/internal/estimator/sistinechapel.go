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

const kSistineChapelInM = 21

type SistineChapelEstimator struct {}

// Returns the size of Sistine Chapel.
func (s *SistineChapelEstimator) K() float64 {
    return kSistineChapelInM
}

// Returns a string with some estimative of your code against Sistine Chapel size.
func (s *SistineChapelEstimator) Estimate(codestat *ruler.CodeStat) string {
    m := &measurer.MCodeStat{}
    m.Calibrate(codestat)
    return doEstimative(m, "Your code has %.2f%% of the Sistine Chapel's height (%.fm).",
                           "Sistine Chapel's height (%.fm) has %.2f%% of your code extension (%.2fm).", s)
}
