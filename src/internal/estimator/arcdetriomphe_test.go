//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package estimator

import (
    "testing"
    "internal/ruler"
)

func TestArcDeTriompheEstimatorK(t *testing.T) {
    a := ArcDeTriompheEstimator{}
    if a.K() != 50 {
        t.Error(`a.K() != 50`)
    }
}

func TestArcDeTrimpheEstimatorEstimate(t *testing.T) {
    testVector := []struct {
                    BytesTotal int64
                    ExpectedMessage string
                  }{
                    { 10, `Your code has 0.07% of the Arc de Triomphe's height (50m).` },
                    { 100000, `Arc de Triomphe's height (50m) has 13.80% of your code extension (362.20m).` },
                 }
    for _, test := range testVector {
        codestat := &ruler.CodeStat{}
        codestat.CalibrateCourier12px()
        codestat.Files = make(map[string]ruler.CodeFileInfo)
        codestat.Files["main.go"] = ruler.CodeFileInfo{test.BytesTotal}
        a := ArcDeTriompheEstimator{}
        estimative := a.Estimate(codestat)
        if estimative != test.ExpectedMessage {
            t.Errorf(`estimative != test.ExpectedMessage: %v != %v`, estimative, test.ExpectedMessage)
        }
    }
}


