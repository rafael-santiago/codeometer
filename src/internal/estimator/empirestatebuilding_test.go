//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package estimator

import (
    "testing"
    "internal/ruler"
)

func TestEmpireStateBuildingEstimatorK(t *testing.T) {
    e := EmpireStateBuildingEstimator{}
    if e.K() != 381 {
        t.Error(`e.K() != 381`)
    }
}

func TestEmpireStateBuildingEstimatorEstimate(t *testing.T) {
    testVector := []struct {
                    BytesTotal int64
                    ExpectedMessage string
                  }{
                    { 10, `Your code has 0.01% of the Empire State Building's height (381m).` },
                    { 213399, `Empire State Building's height (381m) has 49.29% of your code extension (772.92m).` },
                 }
    for _, test := range testVector {
        codestat := &ruler.CodeStat{}
        codestat.CalibrateCourier12px()
        codestat.Files = make(map[string]ruler.CodeFileInfo)
        codestat.Files["main.go"] = ruler.CodeFileInfo{test.BytesTotal}
        e := EmpireStateBuildingEstimator{}
        estimative := e.Estimate(codestat)
        if estimative != test.ExpectedMessage {
            t.Errorf(`estimative != test.ExpectedMessage: %v != %v`, estimative, test.ExpectedMessage)
        }
    }
}


