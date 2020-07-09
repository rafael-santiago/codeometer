//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package estimator

import (
    "testing"
    "internal/ruler"
)

func TestWashingtonMonumentEstimatorK(t *testing.T) {
    w := WashingtonMonumentEstimator{}
    if w.K() != 169 {
        t.Error(`w.K() != 169`)
    }
}

func TestWashingtonMonumentEstimatorEstimate(t *testing.T) {
    testVector := []struct {
                    BytesTotal int64
                    ExpectedMessage string
                  }{
                    { 512, `Your code has 1.10% of the Washington Monument's height (169m).` },
                    { 512 << 10, `Washington Monument's height (169m) has 8.90% of your code extension (1898.95m).` },
                 }
    for _, test := range testVector {
        codestat := &ruler.CodeStat{}
        codestat.CalibrateCourier12px()
        codestat.Files = make(map[string]ruler.CodeFileInfo)
        codestat.Files["main.go"] = ruler.CodeFileInfo{test.BytesTotal}
        w := WashingtonMonumentEstimator{}
        estimative := w.Estimate(codestat)
        if estimative != test.ExpectedMessage {
            t.Errorf(`estimative != test.ExpectedMessage: %v != %v`, estimative, test.ExpectedMessage)
        }
    }
}
