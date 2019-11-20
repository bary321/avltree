package avltree

// 采用懒惰删除
type AVLTree struct {
	root Node
}

func (at *AVLTree) GetRoot() Node {
	return at.root
}

func (at *AVLTree) SetRoot(node Node) {
	at.root = node
}

func (at *AVLTree) Find(data int64) Node {
	if at.root == nil {
		return nil
	}
	tmp := at.root
	for {
		if tmp.GetData() == data {
			t := tmp.(*AVLNode)
			if !t.Del {
				return tmp
			}
			return nil
		} else if tmp.GetData() < data {
			if tmp.GetRight() != nil {
				tmp = tmp.GetRight()
			} else {
				return nil
			}
		} else {
			if tmp.GetLeft() != nil {
				tmp = tmp.GetLeft()
			} else {
				return nil
			}
		}
	}
}

func (at *AVLTree) FindMin() Node {
	if at.root == nil {
		return nil
	}
	tmp := at.root.(*AVLNode)
	reverse := false
	ok := true
	for {
		if reverse {
			tmp, ok = tmp.GetParent().(*AVLNode)
			if !ok {
				return nil
			}
			if tmp == nil {
				return nil
			}
			if !tmp.IsDel() {
				return tmp
			}
		} else {
			if tmp.GetLeft() != nil {
				tmp = tmp.GetLeft().(*AVLNode)
			} else {
				if tmp.IsDel() {
					reverse = true
					continue
				}
				return tmp
			}
		}
	}
}

func (at *AVLTree) FindMax() Node {
	if at.root == nil {
		return nil
	}
	tmp := at.root.(*AVLNode)
	reverse := false
	ok := true
	for {
		if reverse {
			tmp, ok = tmp.GetParent().(*AVLNode)
			if !ok {
				return nil
			}
			if tmp == nil {
				return nil
			}
			if !tmp.IsDel() {
				return tmp
			}
		} else {
			if tmp.GetRight() != nil {
				tmp = tmp.GetRight().(*AVLNode)
			} else {
				if tmp.IsDel() {
					reverse = true
					continue
				}
				return tmp
			}
		}
	}
}

func (at *AVLTree) Insert(data int64) {
	if at.root == nil {
		at.root = NewAVLNode(data)
		at.root.(*AVLNode).SetDepth(1)
		return
	}
}

func RightSingleRotation(node *AVLNode) Node {
	left := node.Left.(*AVLNode)
	var c *AVLNode
	if left.Right != nil {
		c = left.Right.(*AVLNode)
		c.Parent = node
	}
	up := node.Parent
	left.Right = node
	node.SetLeft(c)
	node.Parent = left
	left.Parent = up
	return left
}

func LeftSingleRotation(node *AVLNode) Node {
	up := node.Parent
	right := node.Right.(*AVLNode)
	node.Right = node.Right.GetLeft()
	right.SetParent(up)
	right.SetLeft(node)
	node.SetParent(right)
	if node.Right != nil {
		node.Right.(*AVLNode).SetParent(node)
	}
	return right
}

func RightDoubleRotation(node *AVLNode) Node {
	if node.Left != nil {
		node.Left = LeftSingleRotation(node.Left.(*AVLNode))
	}
	return RightSingleRotation(node)
}

func LeftDoubleRotation(node *AVLNode) Node {
	if node.Right != nil {
		node.Right = LeftSingleRotation(node.Right.(*AVLNode))
	}
	return LeftSingleRotation(node)
}

func (at *AVLTree) Delete(data int64) {
	node, ok := at.Find(data).(*AVLNode)
	if ok {
		node.Delete()
	}
}

func (at *AVLTree) PopMin() Node {
	node, ok := at.FindMin().(*AVLNode)
	if ok {
		node.Delete()
	}
	return node
}

func (at *AVLTree) PopMax() Node {
	node, ok := at.FindMax().(*AVLNode)
	if ok {
		node.Delete()
	}
	return node
}

func (at *AVLTree) Display() {
	PrintTree(at)
}
