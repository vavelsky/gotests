package iterations

import (
	"fmt"
	"testing"
)

func TestIter(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkIter(b *testing.B) {
	for i := 0; i < 5; i++ {
		Repeat("a", 4)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
