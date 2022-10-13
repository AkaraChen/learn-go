package ch

import (
	"testing"
)

func TestDown(t *testing.T) {
	t.Run("Ch", func(t *testing.T) {
		expect := 3
		got := Down()

		if expect != got {
			t.Fatal("Something Wrong")
		}
	})
}
