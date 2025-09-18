package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		tests := []struct {
			nums []int
			want []int
		}{
			{nums: []int{1, 2, 3, 4}, want: []int{24, 12, 8, 6}},
			{nums: []int{-1, 1, 0, -3, 3}, want: []int{0, 0, 9, 0, 0}},
			{nums: []int{4, 3, 2, 1, 2}, want: []int{12, 16, 24, 48, 24}},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("nums=%v", tt.nums), func(t *testing.T) {
				if got := productExceptSelf(tt.nums); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
				}
			})
		}
	})

}
