package graph

import "testing"

func TestPop(t *testing.T) {
	tests := []struct {
		Input          *Queue[int]
		ExpectedValue  int
		ExpectedLength int
	}{
		{
			Input:          NewQueue(1, 2, 3),
			ExpectedValue:  1,
			ExpectedLength: 2,
		},
	}

	for id, test := range tests {
		if value := test.Input.Pop(); value != test.ExpectedValue {
			t.Errorf("[case %2d] Expected %d but got %d", id, test.ExpectedValue, value)
		}
		if length := test.Input.Len(); length != test.ExpectedLength {
			t.Errorf("[case %2d] Expected length %d but got %d", id, test.ExpectedLength, length)
		}
	}
}
