package v1

import (
	"testing"
)

// Test whether the generate package works.
func TestCanary(t *testing.T) {
	o := ObjectMeta{Name: "test"}

	if o.Name != "test" {
		t.Errorf("Something is very wrong.")
	}
}
