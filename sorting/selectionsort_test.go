package sorting

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	type args struct {
		source []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"3,2,6,3,4", args{[]int{3,2,6,3,4}}, []int{2,3,3,4,6}},
		{"5,4,3,2,1", args{[]int{5,4,3,2,1}}, []int{1,2,3,4,5}},
		{"1,2,3,4,5", args{[]int{1,2,3,4,5}}, []int{1,2,3,4,5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectionSort(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
