//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package estimator

import (
    "testing"
    "internal/ruler"
)

func TestFrogTravelerK(t *testing.T) {
    f := FrogTraveler{}
    if f.K() != 44 {
        t.Error(`f.K() != 44`)
    }
}

func TestFrogTravelerEstimate(t *testing.T) {
    testVector := []struct {
                    BytesTotal int64
                    ExpectedMessage string
                  }{
                    { 10, `main.go has 82.32% of the Frog-Traveler's height (44 mm).` },
                    { 512, `Frog-Traveler's height (44 mm) has 2.37% of main.go extension (1854.44 mm).` },
                 }
    for _, test := range testVector {
        codestat := &ruler.CodeStat{}
        codestat.CalibrateCourier12px()
        codestat.Files = make(map[string]ruler.CodeFileInfo)
        codestat.Files["main.go"] = ruler.CodeFileInfo{test.BytesTotal}
        codestat.ProjectName = "main.go"
        f := FrogTraveler{}
        estimative := f.Estimate(codestat)
        if estimative != test.ExpectedMessage {
            t.Errorf(`estimative != test.ExpectedMessage: %v != %v`, estimative, test.ExpectedMessage)
        }
    }
}


