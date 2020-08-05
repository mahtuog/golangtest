package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T){

	t.Run("add two numbers", func(t *testing.T) {
		got := Add(1, 2)
		want := 3

		if got != want{
			t.Errorf("got %d want %d", got, want)
		}
	})

}

func ExampleAdder(){
	sum := Add(1, 5)
	fmt.Println(sum)
}