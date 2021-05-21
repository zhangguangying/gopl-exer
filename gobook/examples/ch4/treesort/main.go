package main

type tree struct {
	value int
	left, right *tree
}

func Sort(values []int)  {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, root *tree) []int {
	if root != nil {
		appendValues(values, root.left)
		values = append(values, root.value)
		appendValues(values, root.right)
	}
	return values
}

func add(t *tree, v int) *tree {
	if t == nil {
		t = new(tree)
		t.value = v
		return t
	}
	if v < t.value {
		t.left = add(t.left, v)
	} else {
		t.right = add(t.right, v)
	}
	return t
}
