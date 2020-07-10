//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package estimator

import (
    "testing"
    "internal/ruler"
)

func TestNiagaraFallsK(t *testing.T) {
    n := NiagaraFalls{}
    if n.K() != 51 {
        t.Error(`n.K() != 51`)
    }
}

func TestNiagaraFallsEstimate(t *testing.T) {
    testVector := []struct {
                    BytesTotal int64
                    ExpectedMessage string
                  }{
                    { 10, `main.go has 0.07% of the Niagara Falls' height (51m).` },
                    { 512 << 6, `Niagara Falls' height (51m) has 42.97% of main.go extension (118.68m).` },
                 }
    for _, test := range testVector {
        codestat := &ruler.CodeStat{}
        codestat.CalibrateCourier12px()
        codestat.Files = make(map[string]ruler.CodeFileInfo)
        codestat.Files["main.go"] = ruler.CodeFileInfo{test.BytesTotal}
        codestat.ProjectName = "main.go"
        n := NiagaraFalls{}
        estimative := n.Estimate(codestat)
        if estimative != test.ExpectedMessage {
            t.Errorf(`estimative != test.ExpectedMessage: %v != %v`, estimative, test.ExpectedMessage)
        }
    }
}


