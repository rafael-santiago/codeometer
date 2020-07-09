//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package estimator

import (
    "testing"
    "internal/ruler"
)

func TestChristTheRedeemerEstimatorK(t *testing.T) {
    c := ChristTheRedeemerEstimator{}
    if c.K() != 38 {
        t.Error(`c.K() != 38`)
    }
}

func TestChristTheRedeemerEstimatorEstimate(t *testing.T) {
    testVector := []struct {
                    BytesTotal int64
                    ExpectedMessage string
                  }{
                    { 10, `Your code has 0.10% of the Christ the Redeemer's height (38m).` },
                    { 13300, `Christ the Redeemer's height (38m) has 78.88% of your code extension (48.17m).` },
                 }
    for _, test := range testVector {
        codestat := &ruler.CodeStat{}
        codestat.CalibrateCourier12px()
        codestat.Files = make(map[string]ruler.CodeFileInfo)
        codestat.Files["main.go"] = ruler.CodeFileInfo{test.BytesTotal}
        c := ChristTheRedeemerEstimator{}
        estimative := c.Estimate(codestat)
        if estimative != test.ExpectedMessage {
            t.Errorf(`estimative != test.ExpectedMessage: %v != %v`, estimative, test.ExpectedMessage)
        }
    }
}


