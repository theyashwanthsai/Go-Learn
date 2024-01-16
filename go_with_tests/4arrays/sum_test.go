package arrays

import (
	"testing"
	"reflect"
)

func TestSum(t *testing.T){
	t.Run("Collection of fixed size 5", func(t *testing.T){
		numbers := []int{1, 2, 3, 4, 5}

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


func TestSumAll(t *testing.T){
	numbers1 := []int{1,2}
	numbers2 := []int{0,9}

	got := SumAll(numbers1, numbers2)
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want){
		t.Errorf("got %v, want %v", got, want)
	}
}


func TestSumAllTails(t *testing.T){

	checkEqual := func(t testing.TB ,got, want  []int){
		t.Helper()
		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("Some of some slices", func(t *testing.T){
		numbers1 := []int{1, 2}
		numbers2 := []int{0, 9}

		got := SumAllTails(numbers1, numbers2)
		want := []int{2, 9}

		checkEqual(t, got, want)
	})

	t.Run("Safely return for an empty slice", func(t *testing.T){
		numbers1 := []int{1, 2}
		numbers2 := []int{}

		got := SumAllTails(numbers1, numbers2)
		want := []int{2, 0}

		checkEqual(t, got, want)
	})
	
}