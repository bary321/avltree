package avltree

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

}

func (at *AVLTree) Delete(data int64) {
	if at.GetRoot() == nil {
		return
	}

	//tmp := at.Delete

}

func (at *AVLTree) PopMin() Node {
	return nil
}

func (at *AVLTree) PopMax() Node {
	return nil
}

func (at *AVLTree) Display() {

}
