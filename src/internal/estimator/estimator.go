// package estimator prints estimatives about your code metrics based on famous places, momuments.
// --
//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package estimator

import (
    "internal/ruler"
    "internal/measurer"
    "fmt"
)

// Defines basic functions that any estimator must have
type Estimator interface {
    Estimate(codestat *ruler.CodeStat) string
    K() float64
}

// Does a specific estimative.
func Estimate(estimator Estimator, codestat *ruler.CodeStat) string {
    return estimator.Estimate(codestat)
}

// Internal function that processes an estimative returning a string containing it.
func doEstimative(measurerHandle interface{}, codeIsLessMessage, codeIsGreaterMessage string,
                  estimator Estimator) string {
    var totalDistance float64
    var projectName string

    switch measurerHandle.(type) {
        case *measurer.MMCodeStat:
            o := measurerHandle.(*measurer.MMCodeStat)
            totalDistance = o.TotalDistance()
            projectName = o.ProjectName
            break

        case *measurer.MCodeStat:
            o := measurerHandle.(*measurer.MCodeStat)
            totalDistance = o.TotalDistance()
            projectName = o.ProjectName
            break

        case *measurer.KMCodeStat:
            o := measurerHandle.(*measurer.KMCodeStat)
            totalDistance = o.TotalDistance()
            projectName = o.ProjectName
            break

        case *measurer.MICodeStat:
            o := measurerHandle.(*measurer.MICodeStat)
            totalDistance = o.TotalDistance()
            projectName = o.ProjectName
            break

        default:
            panic("doEstimative(): Unexpected measurerHandle type.")
            break

    }
    var retval string
    k := estimator.K()
    if totalDistance < k {
        perc := (totalDistance / k) * 100
        retval = fmt.Sprintf(codeIsLessMessage, projectName, perc, k)
    } else {
        perc := (k / totalDistance) * 100
        retval = fmt.Sprintf(codeIsGreaterMessage, k, perc, projectName, totalDistance)
    }
    return retval
}
