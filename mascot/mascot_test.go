package mascot_test

import (
	"testing"

	"example.com/ms-demo-1/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Dog" {
		t.Fatal("Wrong mascot!")
	}
}
