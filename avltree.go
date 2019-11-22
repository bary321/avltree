package avltree

// 采用懒惰删除
type AVLTree struct {
	root Node
}

func (at *AVLTree) GetData() int64 {
	return 0
}

func (at *AVLTree) GetLeft() Node      { return at.root }
func (at *AVLTree) GetRight() Node     { return at.root }
func (at *AVLTree) SetData(int64)      {}
func (at *AVLTree) SetLeft(node Node)  { at.root = node }
func (at *AVLTree) SetRight(node Node) { at.root = node }

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
	node, isleft := at.InsertLeaf(data)
	if node == nil {
		return
	}
	parent, ok := node.GetParent().(*AVLNode)
	if !ok {
		return
	}
	maxLeftDepth := node.Depth
	maxRightDepth := node.Depth
	maxdepth := node.Depth
	for {
		if parent.Left == node {
			maxRightDepth = MaxDepth(parent.Right)
			if maxRightDepth == 0 {
				maxRightDepth = parent.Depth
			}
			if maxdepth-maxRightDepth > 1 {
				ancestor := parent.Parent
				if IsNil(ancestor) {
					ancestor = at
				}
				var rt Node
				if isleft {
					rt = RightSingleRotation(parent)
				} else {
					rt = RightDoubleRotation(parent)
				}
				UpdateDepth(rt.(*AVLNode), parent.Depth)
				Replace(ancestor, parent, rt)
				return
			}
			if maxRightDepth > maxdepth {
				maxdepth = maxRightDepth
			}
		}
		if parent.Right == node {
			maxLeftDepth = MaxDepth(parent.Left)
			if maxLeftDepth == 0 {
				maxLeftDepth = parent.Depth
			}
			if maxdepth-maxLeftDepth > 1 {
				ancestor := parent.Parent
				if IsNil(ancestor) {
					ancestor = at
				}
				var rt Node
				if !isleft {
					rt = LeftSingleRotation(parent)
				} else {
					rt = LeftDoubleRotation(parent)
				}
				UpdateDepth(rt.(*AVLNode), parent.Depth)
				Replace(ancestor, parent, rt)
				return
			}
			if maxLeftDepth > maxdepth {
				maxdepth = maxLeftDepth
			}
		}
		node = parent
		if IsNil(parent.GetParent()) {
			break
		} else {
			parent = parent.GetParent().(*AVLNode)
		}
	}
}

func Replace(parent, src, dst Node) {
	if parent.GetLeft() == src {
		parent.SetLeft(dst)
	} else if parent.GetRight() == src {
		parent.SetRight(dst)
	}
}

func MaxDepth(node Node) int {
	if IsNil(node) {
		return 0
	}
	if IsNil(node.GetLeft()) && IsNil(node.GetRight()) {
		return node.(*AVLNode).Depth
	}

	maxRight := MaxDepth(node.GetRight())
	maxLeft := MaxDepth(node.GetLeft())
	if maxRight > maxLeft {
		return maxRight
	} else {
		return maxLeft
	}
}

func (at *AVLTree) MaxLeftDepth() int {
	node := at.FindMin().(*AVLNode)
	return node.Depth
}

func (at *AVLTree) MaxRightDepth() int {
	node := at.FindMax().(*AVLNode)
	return node.Depth
}

func (at *AVLTree) InsertLeaf(data int64) (*AVLNode, bool) {
	if at.root == nil {
		tmp := NewAVLNode(data)
		//tmp.Parent = at
		at.root = tmp
		at.root.(*AVLNode).SetDepth(1)
		return tmp, false
	}
	isleft := false
	parent := at.root.(*AVLNode)
	tmp := at.root.(*AVLNode)
	//ok := false
	for {
		parent = tmp
		if tmp.GetData() < data {
			tmp, _ = tmp.GetRight().(*AVLNode)
			isleft = false
			if IsNil(tmp) {
				tmp = NewAVLNode(data)
				tmp.Parent = parent
				tmp.Depth = parent.Depth + 1
				parent.Right = tmp
				return tmp, isleft
			}
			continue
		} else if tmp.GetData() > data {
			tmp, _ = tmp.GetLeft().(*AVLNode)
			isleft = true
			if IsNil(tmp) {
				tmp = NewAVLNode(data)
				tmp.Parent = parent
				tmp.Depth = parent.Depth + 1
				parent.Left = tmp
				return tmp, isleft
			}
			continue
		} else {
			if tmp.Del {
				tmp.Del = false
			}
			return nil, isleft
		}
	}
}

func DepthDiff(node *AVLNode) int {
	left, ok := node.GetLeft().(*AVLNode)
	if !ok {
		left = NewAVLNode(0)
	}
	right, ok := node.GetRight().(*AVLNode)
	if !ok {
		right = NewAVLNode(0)
	}
	return left.Depth - right.Depth
}

func UpdateDepth(node *AVLNode, depth int) {
	node.Depth = depth
	if !IsNil(node.Left) {
		UpdateDepth(node.Left.(*AVLNode), depth+1)
	}
	if !IsNil(node.Right) {
		UpdateDepth(node.Right.(*AVLNode), depth+1)
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
		node.Right = RightSingleRotation(node.Right.(*AVLNode))
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
