package tree

import "testing"

func TestCompareSubTrees(t *testing.T) {
	type args struct {
		root1 *node
		root2 *node
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "both nil",
			args: args{
				root1: nil,
				root2: nil,
			},
			want: true,
		},
		{
			name: "only first nil",
			args: args{
				root1: nil,
				root2: &node{Val: 45},
			},
			want: false,
		},
		{
			name: "only second nil",
			args: args{
				root1: &node{Val: 88},
				root2: nil,
			},
			want: false,
		},
		{
			name: "none nil and both equals",
			args: args{
				root1: &node{Val: 109},
				root2: &node{Val: 109},
			},
			want: true,
		},
		{
			name: "none nil and both not equals",
			args: args{
				root1: &node{Val: 71},
				root2: &node{Val: 96},
			},
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := CompareSubTrees(test.args.root1, test.args.root2)

			if got != test.want {
				t.Errorf("want = %v, got = %v\n", test.want, got)
			}
		})
	}
}
