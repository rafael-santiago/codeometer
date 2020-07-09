//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package estimator

import (
    "testing"
    "internal/ruler"
)

func TestChineseGreatWallEstimatorK(t *testing.T) {
    c := ChineseGreatWallEstimator{}
    if c.K() != 21196 {
        t.Error(`c.K() != 21196`)
    }
}

func TestChineseGreatWallEstimatorEstimate(t *testing.T) {
    testVector := []struct {
                    BytesTotal int64
                    ExpectedMessage string
                  }{
                    { 10, `Your code has 0.00% of the Chinese great wall extension (21196 km).` },
                    { 12345678912, `Chinese great wall's extension (21196 km) has 47.40% of your code extension (44715.45 km).` },
                 }
    for _, test := range testVector {
        codestat := &ruler.CodeStat{}
        codestat.CalibrateCourier12px()
        codestat.Files = make(map[string]ruler.CodeFileInfo)
        codestat.Files["main.go"] = ruler.CodeFileInfo{test.BytesTotal}
        c := ChineseGreatWallEstimator{}
        estimative := c.Estimate(codestat)
        if estimative != test.ExpectedMessage {
            t.Errorf(`estimative != test.ExpectedMessage: %v != %v`, estimative, test.ExpectedMessage)
        }
    }
}

