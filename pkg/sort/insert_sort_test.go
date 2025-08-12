package sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "test1",
			arr:  []int{5, 2, 9, 1, 5, 6},
			want: []int{1, 2, 5, 5, 6, 9},
		},
		{
			name: "test2",
			arr:  []int{3, 0, -2, 8, -1},
			want: []int{-2, -1, 0, 3, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := insertSort(tt.arr)
			fmt.Println("got: ", got)
			if !assert.Equal(t, got, tt.want) {
				t.Errorf("insertSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
