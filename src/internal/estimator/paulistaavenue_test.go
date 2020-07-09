//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package estimator

import (
    "testing"
    "internal/ruler"
)

func TestPaulistaAvenueEstimatorK(t *testing.T) {
    p := PaulistaAvenueEstimator{}
    if p.K() != 2.8 {
        t.Error(`p.K() != 2.8`)
    }
}

func TestPaulistaAvenueEstimatorEstimate(t *testing.T) {
    testVector := []struct {
                    BytesTotal int64
                    ExpectedMessage string
                  }{
                    { 512, `Your code has 0.07% of the Paulista avenue extension (2.8 km).` },
                    { 512 << 11, `Paulista avenue's extension (2.8 km) has 73.73% of your code extension (3.80 km).` },
                 }
    for _, test := range testVector {
        codestat := &ruler.CodeStat{}
        codestat.CalibrateCourier12px()
        codestat.Files = make(map[string]ruler.CodeFileInfo)
        codestat.Files["main.go"] = ruler.CodeFileInfo{test.BytesTotal}
        p := PaulistaAvenueEstimator{}
        estimative := p.Estimate(codestat)
        if estimative != test.ExpectedMessage {
            t.Errorf(`estimative != test.ExpectedMessage: %v != %v`, estimative, test.ExpectedMessage)
        }
    }
}