package chainlets

import (
	"strings"
)

type Edge struct {
	Pkg, Dependency string
}

func (e Edge) String() string {
	return e.Pkg + " " + e.Dependency
}

func (e Edge) Reverse() Edge {
	return Edge{Pkg: e.Dependency, Dependency: e.Pkg}
}

// Graph represents a dependency graph comprised of (Pkg -> Dependency) directed pairs
type Graph []Edge

// ExcludePkgs filters the graph to a version that doesn't include the passed pkg in
// any directed pair as a dependent (Pkg)
func (g *Graph) ExcludePkgs(pkgs []string) {
	for _, i := range pkgs {
		g.excludePkg(i)
	}
}

// Chains returns a slice of chains that end with the specified pkg
func (g Graph) Chains(pkg string) []Chain {
	visited := map[string]bool{}
	unfinished := []Chain{{Pkg: pkg, Rest: nil}}
	finished := []Chain{}

	for len(unfinished) != 0 {
		chain := unfinished[0]
		unfinished = unfinished[1:]

		if chain.Rest != nil {
			key := chain.String()
			if _, ok := visited[key]; ok {
				continue
			}

			visited[key] = true
		}

		needsPkg := []Chain{}
		for i := range g {
			if g[i].Dependency == chain.Pkg {

				needsPkg = append(needsPkg, Chain{Pkg: g[i].Pkg, Rest: &chain})
			}
		}

		if len(needsPkg) != 0 {
			unfinished = append(unfinished, needsPkg...)
		} else if chain.Length() != 1 {
			finished = append(finished, chain)

		}
	}

	return finished
}

// HasCircularDep returns true if two packages depend on each other
func (g Graph) CircularDep() []Edge {
	i := 0
	visited := map[Edge]bool{}
	circular := []Edge{}
	for {
		if i >= len(g) {
			break
		}
		rev := g[i].Reverse()
		if _, ok := visited[rev]; ok {
			circular = append(circular, rev)
		}
		visited[g[i]] = true
		i++
	}
	return circular
}

func (g *Graph) excludePkg(pkg string) {
	i := 0
	for {
		if i >= len(*g) {
			break
		}
		if strings.Contains((*g)[i].Pkg, pkg) {
			*g = append((*g)[:i], (*g)[i+1:]...)
		} else {
			i++
		}
	}
}

// StrToGraph takes a `go mod graph` as a string and returns a Graph
func StrToGraph(graph string) Graph {
	edges := Graph{}
	for _, line := range strings.Split(graph, "\n") {
		if strings.Trim(line, " \n\t") == "" {
			continue
		}
		edges = append(edges, *strToEdge(line))
	}
	return edges
}

func strToEdge(line string) *Edge {
	deps := strings.Split(strings.Trim(line, " \n\t"), " ")
	c := Edge{Pkg: deps[0], Dependency: deps[1]}
	return &c
}
