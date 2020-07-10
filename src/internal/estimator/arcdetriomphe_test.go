//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package estimator

import (
    "testing"
    "internal/ruler"
)

func TestArcDeTriompheK(t *testing.T) {
    a := ArcDeTriomphe{}
    if a.K() != 50 {
        t.Error(`a.K() != 50`)
    }
}

func TestArcDeTrimpheEstimate(t *testing.T) {
    testVector := []struct {
                    BytesTotal int64
                    ExpectedMessage string
                  }{
                    { 10, `main.go has 0.07% of the Arc de Triomphe's height (50m).` },
                    { 100000, `Arc de Triomphe's height (50m) has 13.80% of main.go extension (362.20m).` },
                 }
    for _, test := range testVector {
        codestat := &ruler.CodeStat{}
        codestat.CalibrateCourier12px()
        codestat.Files = make(map[string]ruler.CodeFileInfo)
        codestat.Files["main.go"] = ruler.CodeFileInfo{test.BytesTotal}
        codestat.ProjectName = "main.go"
        a := ArcDeTriomphe{}
        estimative := a.Estimate(codestat)
        if estimative != test.ExpectedMessage {
            t.Errorf(`estimative != test.ExpectedMessage: %v != %v`, estimative, test.ExpectedMessage)
        }
    }
}


