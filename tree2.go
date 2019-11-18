package avltree

type Tree2 struct {
	BinaryTree
}

func (t *Tree2) PopMin() Node {
	parent := t.GetRoot()
	if parent == nil {
		return nil
	}
	next := parent.GetLeft()
	if next == nil {
		t.SetRoot(nil)
		return next
	}
	for {
		if next.GetLeft() == nil {
			parent.SetLeft(nil)
			return next
		}
		parent = next
		next = next.GetLeft()
	}
}
