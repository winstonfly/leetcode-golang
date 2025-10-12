package list

import "testing"

func TestReverseKGroup(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		tests := BuildListNode([]int{1, 2, 3, 4, 5})
		v := reverseKGroup(tests, 3)
		t.Log(v)
		current := v
		var res []int
		for current != nil {
			res = append(res, current.Val)
			current = current.Next
		}

		t.Log(res)
	})

	t.Run("test2", func(t *testing.T) {
		tests := BuildListNode([]int{1, 2, 3, 4, 5})
		v := reverseKGroup(tests, 2)
		t.Log(v)
		current := v
		var res []int
		for current != nil {
			res = append(res, current.Val)
			current = current.Next
		}

		t.Log(res)
	})
}
