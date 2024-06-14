package querier

import (
	sitter "github.com/smacker/go-tree-sitter"
)

type matchingFunc func(node *sitter.Node) bool

// Allows for searches in a node instead of a tree.
// Recursively parses every subnodes of a given sitter node to find every child matching
// a pattern defined in a func.
func GetChildrenMatching(curr *sitter.Node, matcher matchingFunc, matched []*sitter.Node) []*sitter.Node {

	for i := 0; i < int(curr.ChildCount()); i++ {
		if matcher(curr.Child(i)) {
			matched = append(matched, curr.Child(i))
		}
		matched = GetChildrenMatching(curr.Child(i), matcher, matched)
	}
	return matched
}

// Allows for searches in a node instead of a tree.
//
// Recursively parses every subnodes of a given sitter node to find every child matching
// a pattern defined in a func.
//
// Stops the search if it exceeds the maxDepht
func GetChildrenMatchingAtMaxDepth(curr *sitter.Node, matcher matchingFunc, matched []*sitter.Node, maxDepht int) []*sitter.Node {
	return getChildrenMatchingAtDepht(curr, matcher, matched, 1, maxDepht)
}

// Recursive call of getChildrenMatchingAtDepht
func getChildrenMatchingAtDepht(curr *sitter.Node, matcher matchingFunc, matched []*sitter.Node, depht, maxDepht int) []*sitter.Node {
	if depht > maxDepht {
		return matched
	}
	for i := 0; i < int(curr.ChildCount()); i++ {
		if matcher(curr.Child(i)) {
			matched = append(matched, curr.Child(i))
		}
		matched = getChildrenMatchingAtDepht(curr.Child(i), matcher, matched, depht+1, maxDepht)
	}
	return matched
}

// Allows for searches in a node instead of a tree.
// Recursively parses every subnodes of a given sitter node to find the first child matching
// a pattern defined in a func.
// If the node itself matches the function, returns itself as the first node matching the condition.
// It's a depht first search
func GetFirstMatch(curr *sitter.Node, matcher matchingFunc) *sitter.Node {

	if matcher(curr) {
		return curr
	}

	for i := 0; i < int(curr.ChildCount()); i++ {
		matched := GetFirstMatch(curr.Child(i), matcher)

		if matched != nil {
			return matched
		}
	}

	return nil
}

// Retrieves all sibling nodes of a given node.
func GetSiblings(node *sitter.Node) []*sitter.Node {
    var siblings []*sitter.Node
    parent := node.Parent()
    if parent == nil {
        return siblings
    }
    for i := 0; i < int(parent.ChildCount()); i++ {
        sibling := parent.Child(i)
        if !sibling.Equal(node) {
            siblings = append(siblings, sibling)
        }
    }
    return siblings
}

// Takes a string as a parameter, and finds all matching nodes with the given type
func QuerySelectorAll(curr *sitter.Node, query string, matched []*sitter.Node) []*sitter.Node {
    return GetChildrenMatching(curr, func(node *sitter.Node) bool {
        return node.Type() == query
    }, matched)
}

// Takes a string as a parameter, and matches it to the type of the node
// Uses depth first search
func QuerySelector(curr *sitter.Node, query string) *sitter.Node {
    return GetFirstMatch(curr, func(node *sitter.Node) bool {
        return node.Type() == query
    })
}

// Takes a string as a parameter, and matches it to the type of the node
// Uses breadth first search
func BreadthSelector(curr *sitter.Node, query string) *sitter.Node {
    return BreadthFirstMatch(curr, func(node *sitter.Node) bool {
        return node.Type() == query
    })
}

// Finds the first node matching a boolean using breadth first search
func BreadthFirstMatch(curr *sitter.Node, matcher matchingFunc) *sitter.Node {
    //mut
    root := curr
    //mut
    queue := []*sitter.Node{root}

    for len(queue) > 0 {
        //mutation
        root = queue[0]
        for i := 0; i < int(root.ChildCount()); i++ {
            n := root.Child(i)
            queue = append(queue, n)
            if matcher(n) {
                return n
            }
        }
        queue = queue[1:]
    }
    return nil
}
