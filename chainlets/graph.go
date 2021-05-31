package chainlets

import (
	"strings"
)

type edge struct {
	Pkg, Dependency string
}

func (e edge) String() string {
	return e.Pkg + " " + e.Dependency
}

// Graph represents a dependecy graph comprised of (Pkg -> Dependecy) directed pairs
type Graph []edge

// ExcludePkgs filters the graph to a version that doesn't include the passed pkg in
// any directed pair as a dependant (Pkg)
func (g Graph) ExcludePkgs(pkgs []string) Graph {
	var gg Graph
	for _, i := range pkgs {
		gg = g.excludePkg(i)
	}
	return gg
}

func (g Graph) excludePkg(pkg string) Graph {
	i := 0
	for {
		if i >= len(g) {
			break
		}
		if strings.Contains(g[i].Pkg, pkg) {
			g = append(g[:i], g[i+1:]...)
		} else {
			i++
		}
	}
	return g
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

func strToEdge(line string) *edge {
	deps := strings.Split(strings.Trim(line, " \n\t"), " ")
	c := edge{Pkg: deps[0], Dependency: deps[1]}
	return &c
}
