package integers

import (
	"testing"
	"fmt"
)

func TestAdder(t *testing.T)  {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d', but got '%d'",expected, sum)
	}
}

func ExampleAdder()  {
	sum := Add(3, 5)
	fmt.Println(sum)
	// Output: 8
}

