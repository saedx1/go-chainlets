package chainlets

const _separator = " -> "

// Chain is a linked list-like representation of a dependency chain
type Chain struct {
	Pkg  string // name of the dependant Pkg
	Rest *Chain // the chain that Pkg depends on
}

// Length returns the number of dependencies in the chain
func (c Chain) Length() int {
	count := 1
	current := c.Rest
	for {
		if current == nil {
			break
		}
		count++
		current = current.Rest
	}
	return count
}

func (c Chain) String() string {
	str := ""
	str += c.Pkg
	current := c.Rest
	for {
		if current == nil {
			break
		}
		str += _separator
		str += current.Pkg
		current = current.Rest
	}
	return str
}
