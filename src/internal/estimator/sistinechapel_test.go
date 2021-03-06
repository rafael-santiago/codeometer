//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package estimator

import (
    "testing"
    "internal/ruler"
)

func TestSistineChapelK(t *testing.T) {
    s := SistineChapel{}
    if s.K() != 21 {
        t.Error(`s.K() != 21`)
    }
}

func TestSistineChapelEstimate(t *testing.T) {
    testVector := []struct {
                    BytesTotal int64
                    ExpectedMessage string
                  }{
                    { 512, `main.go has 8.83% of the Sistine Chapel's height (21m).` },
                    { 32939, `Sistine Chapel's height (21m) has 17.60% of main.go extension (119.30m).` },
                 }
    for _, test := range testVector {
        codestat := &ruler.CodeStat{}
        codestat.CalibrateCourier12px()
        codestat.Files = make(map[string]ruler.CodeFileInfo)
        codestat.Files["main.go"] = ruler.CodeFileInfo{test.BytesTotal}
        codestat.ProjectName = "main.go"
        s := SistineChapel{}
        estimative := s.Estimate(codestat)
        if estimative != test.ExpectedMessage {
            t.Errorf(`estimative != test.ExpectedMessage: %v != %v`, estimative, test.ExpectedMessage)
        }
    }
}


