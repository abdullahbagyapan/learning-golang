package test01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivideTwoValues(t *testing.T) {

	const value1 = 10
	const value2 = 2

	const result = value1 / value2

	testResult := DivideTwoValues(value1, value2)

	assert.EqualValues(t, result, testResult, "Two result should be the same")

}
