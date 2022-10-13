package waitgroup

import (
	"testing"
)

func TestWaitGroup(t *testing.T) {
	expect := 3
	got := WaitGroup()

	if expect != got {
		t.Fatal("Something wrong")
	}
}
