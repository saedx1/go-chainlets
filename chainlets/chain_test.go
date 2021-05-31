package chainlets_test

import (
	"testing"

	"github.com/saedx1/go-chainlets/chainlets"
)

func TestLength(t *testing.T) {
	c := chainlets.Chain{Pkg: "pkgA", Rest: &chainlets.Chain{Pkg: "pkgB", Rest: nil}}
	expectedLength := 2
	if c.Length() != expectedLength {
		t.Errorf("Incorrect chain length; expected %d, found %d", expectedLength, c.Length())
	}

	c.Rest.Rest = &chainlets.Chain{Pkg: "pkgC", Rest: nil}
	expectedLength = 3
	if c.Length() != expectedLength {
		t.Errorf("Incorrect chain length; expected %d, found %d", expectedLength, c.Length())
	}

	c = chainlets.Chain{}
	expectedLength = 1
	if c.Length() != expectedLength {
		t.Errorf("Incorrect chain length; expected %d, found %d", expectedLength, c.Length())
	}
}

func TestString(t *testing.T) {
	c := chainlets.Chain{Pkg: "pkgA", Rest: &chainlets.Chain{Pkg: "pkgB", Rest: nil}}
	expectedString := "pkgA -> pkgB"
	if c.String() != expectedString {
		t.Errorf("Incorrect chain string; expected %s, found %s", expectedString, c.String())
	}

	c.Rest.Rest = &chainlets.Chain{Pkg: "pkgC", Rest: nil}
	expectedString = "pkgA -> pkgB -> pkgC"
	if c.String() != expectedString {
		t.Errorf("Incorrect chain string; expected %s, found %s", expectedString, c.String())
	}

	c = chainlets.Chain{}
	expectedString = ""
	if c.String() != expectedString {
		t.Errorf("Incorrect chain string; expected %s, found %s", expectedString, c.String())
	}
}
