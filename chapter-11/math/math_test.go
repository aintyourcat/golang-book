// Writing tests for our code is a good way to ensure quality and improve reliability.

package math

import "testing"

/*
We could run the test with the `go test` command.
That command will look for any tests in any of the files in current folder and run them.
Tests are identified with the word `Test` and taking one argument of type `*testing.T`.
For example, since we're testing the `Average` function we name the test function `TestAverage`.

We write tests that use the code we're testing.

In `TestAverage`, we create a struct to represent the inputs and expected outputs for the `Average` function.
Then we create a list of these structs (pairs).
After that we loop through each one and check the `Average` function invocation output.

Run the following tests with `Go111MODULE=auto go test`
*/

func TestAverage(t *testing.T) {
	type testpair struct {
		values  []float64
		average float64
	}

	var tests = []testpair{
		{[]float64{1, 2}, 1.5},
		{[]float64{1, 1, 1, 1, 1, 1}, 1.0},
		{[]float64{-1, 1}, 0},
		{[]float64{}, 0},
	}

	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For pair", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

func TestMax(t *testing.T) {
	type testpair struct {
		values []float64
		max    float64
	}

	var tests = []testpair{
		{[]float64{6, 7, 8, 9, 0}, 9},
		{[]float64{1, 1, 1, 1}, 1.0},
		{[]float64{-1, 1}, 1.0},
		{[]float64{}, 0},
	}

	for _, pair := range tests {
		v := Max(pair.values)
		if v != pair.max {
			t.Error(
				"For pair", pair.values,
				"expected", pair.max,
				"got", v,
			)
		}
	}
}

func TestMin(t *testing.T) {
	type testpair struct {
		values []float64
		min    float64
	}

	var tests = []testpair{
		{[]float64{6, 7, 8, 9, 10}, 6},
		{[]float64{1, 1, 1, 1}, 1.0},
		{[]float64{-1, 1}, -1.0},
		{[]float64{}, 0},
	}

	for _, pair := range tests {
		v := Min(pair.values)
		if v != pair.min {
			t.Error(
				"For pair", pair.values,
				"expected", pair.min,
				"got", v,
			)
		}
	}
}
