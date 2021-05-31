package chainlets

import (
	"reflect"
	"testing"
)

func TestEdgeString(t *testing.T) {
	expectedVal := "pkgA pkgB"
	e := edge{Pkg: "pkgA", Dependency: "pkgB"}
	val := e.String()
	if val != expectedVal {
		t.Errorf("Incorrect result; expected %v, found %v", expectedVal, val)
	}
}

func TestStrToGraph(t *testing.T) {
	expectedG := Graph([]edge{
		{"pkgA", "pkgB"},
		{"pkgA", "pkgC"},
		{"pkgC", "pkgD"},
		{"pkgD", "pkgF"},
		{"pkgD", "pkgG"},
	})

	graphStr := `pkgA pkgB
	pkgA pkgC
	pkgC pkgD
	pkgD pkgF
	pkgD pkgG`
	g := StrToGraph(graphStr)

	if !reflect.DeepEqual(expectedG, g) {
		t.Errorf("Incorrect result; expected %v, found %v", expectedG, g)
	}
}

func TestGraphExcludePkgs(t *testing.T) {
	expectedG := Graph([]edge{
		{"pkgA", "pkgB"},
		{"pkgA", "pkgC"},
	})

	graphStr := `pkgA pkgB
	pkgA pkgC
	pkgC pkgD
	pkgD pkgF
	pkgD pkgG`
	g := StrToGraph(graphStr)
	g = g.ExcludePkgs([]string{"pkgC", "pkgD"})

	if !reflect.DeepEqual(expectedG, g) {
		t.Errorf("Incorrect result; expected %v, found %v", expectedG, g)
	}
}

func TestGraphChains(t *testing.T) {
	expectedChains := []Chain{{Pkg: "pkgA", Rest: &Chain{Pkg: "pkgC", Rest: &Chain{Pkg: "pkgD", Rest: &Chain{Pkg: "pkgG", Rest: nil}}}}}
	graphStr := `pkgA pkgB
	pkgA pkgC
	pkgC pkgD
	pkgD pkgF
	pkgD pkgG`
	g := StrToGraph(graphStr)

	chains := g.Chains("pkgG")
	if !reflect.DeepEqual(expectedChains, chains) {
		t.Errorf("Incorrect result; expected %v, found %v", expectedChains, chains)
	}

	expectedChains = []Chain{}
	chains = g.Chains("pkgW")
	if !reflect.DeepEqual(expectedChains, chains) {
		t.Errorf("Incorrect result; expected %v, found %v", expectedChains, chains)
	}

}
