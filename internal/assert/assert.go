package assert

import (
	"testing"
)

func Equal[T comparable](t *testing.T, got T, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}
