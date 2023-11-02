package sort

import (
	"slices"
	"testing"
)

func TestInsertionSort1(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "first test",
			args: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "second test",
			args: []int{10, 43, 25, -65, 3, 98, 154, -52, 25, -12},
			want: []int{-65, -52, -12, 3, 10, 25, 25, 43, 98, 154},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := InsertionSort1(test.args); slices.Compare(got, test.want) != 0 {
				t.Errorf("InsertionSort1 = %v, want = %v", got, test.want)
			}
		})
	}
}
