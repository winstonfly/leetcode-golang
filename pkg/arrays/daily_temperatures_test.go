package arrays

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	var temperatures []int
	var expected []int
	var result []int
	//temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	//expected := []int{1, 1, 4, 2, 1, 1, 0, 0}
	//result := dailyTemperatures(temperatures)
	//if !reflect.DeepEqual(result, expected) {
	//	t.Errorf("expected %v, got %v", expected, result)
	//}
	//
	//temperatures = []int{30, 40, 50, 60}
	//expected = []int{1, 1, 1, 0}
	//result = dailyTemperatures(temperatures)
	//if !reflect.DeepEqual(result, expected) {
	//	t.Errorf("expected %v, got %v", expected, result)
	//}
	//
	//temperatures = []int{30}
	//expected = []int{0}
	//result = dailyTemperatures(temperatures)
	//if !reflect.DeepEqual(result, expected) {
	//	t.Errorf("expected %v, got %v", expected, result)
	//}

	//temperatures = []int{99, 99, 99, 99, 99, 99, 100}
	//expected = []int{6, 5, 4, 3, 2, 1, 0}
	//result = dailyTemperatures(temperatures)
	//if !reflect.DeepEqual(result, expected) {
	//	t.Errorf("expected %v, got %v", expected, result)
	//}

	temperatures = []int{34, 80, 80, 34, 34, 80, 80, 80, 80, 34}
	result = dailyTemperatures(temperatures)
	expected = []int{1, 0, 0, 2, 1, 0, 0, 0, 0, 0}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
