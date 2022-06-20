package learn

import "testing"

func TestFunc(t *testing.T) {
	got, _ := NextInt([]byte{0, 1, 1}, 0)

	if got != 3 {
		t.Errorf("got %q want %q", got, 3)
	}
}
