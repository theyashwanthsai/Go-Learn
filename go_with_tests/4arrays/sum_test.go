package arrays

import (
	"testing"
)

func TestSum(t *testing.T){
	t.Run("Collection of fixed size 5", func(t *testing.T){
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15
	
		if got != want{
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})

	t.Run("Collection of any size", func(t *testing.T){
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want{
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})
	
}