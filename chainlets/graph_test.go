package chainlets

import (
	"reflect"
	"testing"
)

func TestEdgeString(t *testing.T) {
	expectedVal := "pkgA pkgB"
	e := Edge{Pkg: "pkgA", Dependency: "pkgB"}
	val := e.String()
	if val != expectedVal {
		t.Errorf("Incorrect result; expected %v, found %v", expectedVal, val)
	}
}

func TestStrToGraph(t *testing.T) {
	expectedG := Graph([]Edge{
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
	expectedG := Graph([]Edge{
		{"pkgA", "pkgB"},
		{"pkgA", "pkgC"},
	})

	graphStr := `pkgA pkgB
	pkgA pkgC
	pkgC pkgD
	pkgD pkgF
	pkgD pkgG`
	g := StrToGraph(graphStr)
	g.ExcludePkgs([]string{"pkgC", "pkgD"})

	if !reflect.DeepEqual(expectedG, g) {
		t.Errorf("Incorrect result; expected %v, found %v", expectedG, g)
	}
}

func TestGraphChains(t *testing.T) {
	expectedChains := []Chain{
		{Pkg: "pkgA", Rest: &Chain{Pkg: "pkgB", Rest: &Chain{Pkg: "pkgG", Rest: nil}}},
		{Pkg: "pkgA", Rest: &Chain{Pkg: "pkgC", Rest: &Chain{Pkg: "pkgD", Rest: &Chain{Pkg: "pkgG", Rest: nil}}}},
	}
	graphStr := `pkgA pkgB
	pkgA pkgC
	pkgC pkgD
	pkgD pkgF
	
	pkgD pkgG
	pkgB pkgG`
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

func TestEdgeReverse(t *testing.T) {
	e := Edge{Pkg: "A", Dependency: "B"}.Reverse()
	expectedEdge := Edge{Pkg: "B", Dependency: "A"}
	if e != expectedEdge {
		t.Errorf("Incorrect result; expected %v, found %v", expectedEdge, e)
	}
}

func TestGraphCircularDep(t *testing.T) {
	graphStr := `pkgA pkgB
	pkgA pkgC
	pkgC pkgD
	pkgD pkgF
	pkgD pkgG`
	g := StrToGraph(graphStr)
	res := g.CircularDep()
	expectedLength := 0
	if len(res) != expectedLength {
		t.Errorf("Incorrect result length; expected length %v, found %v", expectedLength, len(res))
	}

	graphStr = `pkgA pkgB
	pkgB pkgA
	pkgC pkgD
	pkgD pkgC
	pkgD pkgG`
	g = StrToGraph(graphStr)
	res = g.CircularDep()
	expectedResult := []Edge{{"pkgA", "pkgB"}, {"pkgC", "pkgD"}}

	if len(res) != len(expectedResult) {
		t.Errorf("Incorrect result length; expected length %v, found %v", len(expectedResult), len(res))
	}

	if !reflect.DeepEqual(res, expectedResult) {
		t.Errorf("Incorrect result; expected %v, found %v", expectedResult, res)
	}

}
