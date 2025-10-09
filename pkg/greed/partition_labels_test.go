package greed

import (
	"slices"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		v := partitionLabels("ababcbacadefegdehijhklij")
		want := []int{9, 7, 8}
		if !slices.Equal(want, v) {
			t.Errorf("want %v, got %v", want, v)
		}
	})
}
