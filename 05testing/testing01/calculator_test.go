package calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const value1 = 35
const value2 = 15

var calculator = Calculate{}

func TestAdd(t *testing.T) {

	const expected = value1 + value2

	result := calculator.Add(value1, value2)

	if result != expected {
		t.Errorf("TestAdd failed expected %d , Got %d", expected, result)
	} else {
		t.Logf("TestAdd passed expected %d , Got %d", expected, result)
	}

}

func TestDivide(t *testing.T) {

	const expected = float64(value1) / float64(value2)

	result := calculator.Divide(float64(value1), float64(value2))

	if result != expected {
		t.Errorf("TestDivide failed expected %f , Got %f", expected, result)
	} else {
		t.Logf("TestDivide passed expected %f , Got %f", expected, result)
	}

}

func TestMultiply(t *testing.T) {

	tables := []struct { // our test data
		value1   int
		value2   int
		expected int
	}{
		{40, 20, 800},
		{50, 4, 200},
	}

	for _, v := range tables {
		result := calculator.Multiply(v.value1, v.value2)

		if result != v.expected {
			assert.EqualValues(t, v.expected, result, "TestMultiply failed expected %d , Got %d", v.expected, result)
			break
		}
	}

}

func TestSubtract(t *testing.T) {

	const expected = value1 - value2

	result := calculator.Subtract(value1, value2)

	if result != expected {
		t.Errorf("TestSubtract failed expected %d , Got %d", expected, result)
	} else {
		t.Logf("TestSubtract passed expected %d , Got %d", expected, result)
	}

}
