package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	want := 4

	if sum != want {
		t.Errorf("want %d but get %d", want, sum)
	}
}

func ExampleAdd() {
	sum := Add(2, 3)
	fmt.Println(sum)
	//Output: 5
}
