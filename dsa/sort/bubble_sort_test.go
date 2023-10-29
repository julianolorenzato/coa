package sort

import (
	"slices"
	"testing"
)

func TestBubbleSort1(t *testing.T) {
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
			if got := BubbleSort1(test.args); slices.Compare(got, test.want) != 0 {
				t.Errorf("BubbleSort1 = %v, want = %v", got, test.want)
			}
		})
	}
}

func BenchmarkBubbleSort1(b *testing.B) {
	benchmarks := []struct {
		name      string
		sliceSize int
	}{
		{
			name:      "benchmark 100",
			sliceSize: 100,
		},
		{
			name:      "benchmark 1_000",
			sliceSize: 1_000,
		},
		{
			name:      "benchmark 10_000",
			sliceSize: 10_000,
		},
		{
			name:      "benchmark 20_000",
			sliceSize: 20_000,
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				input := randomSlice(bm.sliceSize)
				b.StartTimer()
				BubbleSort1(input)
			}
		})
	}
}

func TestBubbleSort2(t *testing.T) {
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
			if got := BubbleSort2(test.args); slices.Compare(got, test.want) != 0 {
				t.Errorf("BubbleSort2 = %v, want = %v", got, test.want)
			}
		})
	}
}

func BenchmarkBubbleSort2(b *testing.B) {
	benchmarks := []struct {
		name      string
		sliceSize int
	}{
		{
			name:      "benchmark 100",
			sliceSize: 100,
		},
		{
			name:      "benchmark 1_000",
			sliceSize: 1_000,
		},
		{
			name:      "benchmark 10_000",
			sliceSize: 10_000,
		},
		{
			name:      "benchmark 20_000",
			sliceSize: 20_000,
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				input := randomSlice(bm.sliceSize)
				b.StartTimer()
				BubbleSort2(input)
			}
		})
	}
}

func TestBubbleSort3(t *testing.T) {
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
			if got := BubbleSort3(test.args); slices.Compare(got, test.want) != 0 {
				t.Errorf("BubbleSort3 = %v, want = %v", got, test.want)
			}
		})
	}
}

func BenchmarkBubbleSort3(b *testing.B) {
	benchmarks := []struct {
		name      string
		sliceSize int
	}{
		{
			name:      "benchmark 100",
			sliceSize: 100,
		},
		{
			name:      "benchmark 1_000",
			sliceSize: 1_000,
		},
		{
			name:      "benchmark 10_000",
			sliceSize: 10_000,
		},
		{
			name:      "benchmark 20_000",
			sliceSize: 20_000,
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				input := randomSlice(bm.sliceSize)
				b.StartTimer()
				BubbleSort3(input)
			}
		})
	}
}

func TestBubbleSort4(t *testing.T) {
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
			if got := BubbleSort4(test.args); slices.Compare(got, test.want) != 0 {
				t.Errorf("BubbleSort4 = %v, want = %v", got, test.want)
			}
		})
	}
}

func BenchmarkBubbleSort4(b *testing.B) {
	benchmarks := []struct {
		name      string
		sliceSize int
	}{
		{
			name:      "benchmark 100",
			sliceSize: 100,
		},
		{
			name:      "benchmark 1_000",
			sliceSize: 1_000,
		},
		{
			name:      "benchmark 10_000",
			sliceSize: 10_000,
		},
		{
			name:      "benchmark 20_000",
			sliceSize: 20_000,
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				input := randomSlice(bm.sliceSize)
				b.StartTimer()
				BubbleSort4(input)
			}
		})
	}
}
