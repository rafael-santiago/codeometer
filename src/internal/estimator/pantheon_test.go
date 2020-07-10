//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package estimator

import (
    "testing"
    "internal/ruler"
)

func TestPantheonK(t *testing.T) {
    p := Pantheon{}
    if p.K() != 14 {
        t.Error(`p.K() != 14`)
    }
}

func TestPantheonEstimate(t *testing.T) {
    testVector := []struct {
                    BytesTotal int64
                    ExpectedMessage string
                  }{
                    { 10, `main.go has 0.26% of the Pantheon's height (14m).` },
                    { 512 << 6, `Pantheon's height (14m) has 11.80% of main.go extension (118.68m).` },
                 }
    for _, test := range testVector {
        codestat := &ruler.CodeStat{}
        codestat.CalibrateCourier12px()
        codestat.Files = make(map[string]ruler.CodeFileInfo)
        codestat.Files["main.go"] = ruler.CodeFileInfo{test.BytesTotal}
        codestat.ProjectName = "main.go"
        p := Pantheon{}
        estimative := p.Estimate(codestat)
        if estimative != test.ExpectedMessage {
            t.Errorf(`estimative != test.ExpectedMessage: %v != %v`, estimative, test.ExpectedMessage)
        }
    }
}


