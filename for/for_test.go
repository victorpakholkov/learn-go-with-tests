package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	got := Repeat("a", 4)
	wanted := "aaaa"
	if got != wanted {
		t.Errorf("wanted: '%q', got: '%q'", wanted, got)
	}

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 8)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 5)
	fmt.Println(repeated)
	// Output: bbbbb
}
