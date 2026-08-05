// Test for the mascot package. Verifies that BestMascot returns the
// expected mascot name.
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
