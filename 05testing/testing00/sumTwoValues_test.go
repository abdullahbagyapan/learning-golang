package test00

import (
	"testing"
)

func TestSumTwoValues(t *testing.T) {

	const value1 = 10
	const value2 = 20

	const result = value1 + value2

	testResult := SumTwoValues(value1, value2)

	if result != testResult {
		t.Errorf("sumTwoValues(%d,%d) failed. Expected %d , got %d", value1, value2, result, testResult)
	} else {
		t.Logf("sumTwoValues(%d,%d) passed. Expected %d , got %d", value1, value2, result, testResult)
	}

}
