package graph

import "testing"

func TestDFS(t *testing.T) {
	tests := []struct {
		Input    *Node[int]
		Expected int
	}{
		{
			Input: &Node[int]{
				Value: 0,
				Children: []*Node[int]{
					{Value: 1},
					{Value: 2},
					{Value: 0},
				},
			},
			Expected: 2,
		},
	}
	for id, test := range tests {
		count := 0
		test.Input.DFS(func(number int) bool {
			if number == 0 {
				count++
			}
			return true
		})
		if count != test.Expected {
			t.Errorf("[case %2d] Expected %d but got %d", id, test.Expected, count)
		}
	}
}

func TestBFS(t *testing.T) {
	tests := []struct {
		Input    *Node[int]
		Expected int
	}{
		{
			Input: &Node[int]{
				Value: 0,
				Children: []*Node[int]{
					{Value: 1},
					{Value: 2},
					{Value: 0},
				},
			},
			Expected: 2,
		},
	}
	for id, test := range tests {
		count := 0
		test.Input.BFS(func(number int) bool {
			if number == 0 {
				count++
			}
			return true
		})
		if count != test.Expected {
			t.Errorf("[case %2d] Expected %d but got %d", id, test.Expected, count)
		}
	}
}
