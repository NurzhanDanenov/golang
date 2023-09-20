package main

import "testing"

func Test_equalsWithoutOrder(t *testing.T) {
	type args struct {
		slc1 []int
		slc2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalsWithoutOrder(tt.args.slc1, tt.args.slc2); got != tt.want {
				t.Errorf("equalsWithoutOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
