// package estimator prints estimatives about your code metrics based on famous places, momuments.
// --
package estimator

import (
    "internal/ruler"
    "internal/measurer"
    "fmt"
)

// Defines basic functions that any estimator must have
type Estimator interface {
    Estimate(codestat ruler.CodeStat) string
    K() float64
}

// Does a specific estimative.
func Estimate(estimator Estimator, codestat ruler.CodeStat) string {
    return estimator.Estimate(codestat)
}

// Internal function that processes a estimative returning a string containing it.
func doEstimative(measurerHandle interface{}, codeIsLessMessage, codeIsGreaterMessage string,
                  estimator Estimator) string {
    var totalDistance float64

    switch measurerHandle.(type) {
        case *measurer.MMCodeStat:
            totalDistance = measurerHandle.(*measurer.MMCodeStat).TotalDistance()
            break

        case *measurer.MCodeStat:
            totalDistance = measurerHandle.(*measurer.MCodeStat).TotalDistance()
            break

        case *measurer.KMCodeStat:
            totalDistance = measurerHandle.(*measurer.KMCodeStat).TotalDistance()
            break

        case *measurer.MICodeStat:
            totalDistance = measurerHandle.(*measurer.MICodeStat).TotalDistance()
            break

    }
    var retval string
    k := estimator.K()
    if totalDistance < k {
        perc := (k / 100) * totalDistance
        retval = fmt.Sprintf(codeIsLessMessage, perc, k)
    } else {
        perc := (totalDistance / 100) * k
        retval = fmt.Sprintf(codeIsGreaterMessage, perc, totalDistance)
    }
    return retval
}
