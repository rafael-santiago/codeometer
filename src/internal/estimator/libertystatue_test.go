//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package estimator

import (
    "testing"
    "internal/ruler"
)

func TestLibertyStatueK(t *testing.T) {
    l := LibertyStatue{}
    if l.K() != 93 {
        t.Error(`l.K() != 93`)
    }
}

func TestLibertyStatueEstimate(t *testing.T) {
    testVector := []struct {
                    BytesTotal int64
                    ExpectedMessage string
                  }{
                    { 10, `Your code has 0.04% of the Liberty Statue's height (93m).` },
                    { 512 << 8, `Liberty Statue's height (93m) has 19.59% of your code extension (474.74m).` },
                 }
    for _, test := range testVector {
        codestat := &ruler.CodeStat{}
        codestat.CalibrateCourier12px()
        codestat.Files = make(map[string]ruler.CodeFileInfo)
        codestat.Files["main.go"] = ruler.CodeFileInfo{test.BytesTotal}
        l := LibertyStatue{}
        estimative := l.Estimate(codestat)
        if estimative != test.ExpectedMessage {
            t.Errorf(`estimative != test.ExpectedMessage: %v != %v`, estimative, test.ExpectedMessage)
        }
    }
}


