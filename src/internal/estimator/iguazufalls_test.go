//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package estimator

import (
    "testing"
    "internal/ruler"
)

func TestIguazuFallsK(t *testing.T) {
    i := IguazuFalls{}
    if i.K() != 82 {
        t.Error(`i.K() != 82`)
    }
}

func TestIguazuFallsEstimate(t *testing.T) {
    testVector := []struct {
                    BytesTotal int64
                    ExpectedMessage string
                  }{
                    { 10, `Your code has 0.04% of the Iguazu Falls' height (82m).` },
                    { 512 << 8, `Iguazu Falls' height (82m) has 17.27% of your code extension (474.74m).` },
                 }
    for _, test := range testVector {
        codestat := &ruler.CodeStat{}
        codestat.CalibrateCourier12px()
        codestat.Files = make(map[string]ruler.CodeFileInfo)
        codestat.Files["main.go"] = ruler.CodeFileInfo{test.BytesTotal}
        i := IguazuFalls{}
        estimative := i.Estimate(codestat)
        if estimative != test.ExpectedMessage {
            t.Errorf(`estimative != test.ExpectedMessage: %v != %v`, estimative, test.ExpectedMessage)
        }
    }
}


