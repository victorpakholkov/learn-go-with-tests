package range_

import "testing"

func TestRange(t *testing.T) {
	want := "cccc"
	got := Range("c", 4)

	if got != want {
		t.Errorf("wanted: '%q', got: '%q'", want, got)
	}
}
