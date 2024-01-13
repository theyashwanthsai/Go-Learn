package integers

import(
	"testing"
	"fmt"
)

func TestAdd(t *testing.T){
	got := add(2,2)
	want := 4
	if got != want{
		t.Errorf("got %d but want %d", got, want)
	}
} 


func ExampleAdd(){
	sum := add(2,5)
	fmt.Println(sum)
	// Output: 7
}