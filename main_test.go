package main

import (
	"testing"
)

var testCases = []struct {
	name       string
	a          float64
	b          float64
	wantResult float64
	wantErr    bool
}{
	{"normal division #1", 5, 1, 5.0, false},
	{"normal division #2", 10, 5, 2.0, false},
	{"normal division #3", 40, 10, 4.0, false},
	{"normal division #4", 15, 5, 3.0, false},
	{"division by zero", 7, 0, 0, true},
	{"division with negative nums", -10, 2, -5.0, false},
}

func TestDivide(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := divide(testCase.a, testCase.b)

			if (err != nil) != testCase.wantErr {
				t.Errorf("Ожидаемая ошибка %v, получена: %v", testCase.wantErr, err)
			}

			if !testCase.wantErr && res != testCase.wantResult {
				t.Errorf("Ожидаемый результат %v, получен: %v", testCase.wantResult, res)
			}
		})
	}
}
