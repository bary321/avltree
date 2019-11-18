package avltree

type AVLTree struct {
	BinaryTree
	root Node
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

// todo: update
func (at *AVLTree) FindMin() Node {
	if at.root == nil {
		return nil
	}
	tmp := at.root
	for {
		if tmp.GetLeft() != nil {
			tmp = tmp.GetLeft()
		} else {

			return tmp
		}
	}
}

func (at *AVLTree) Insert(int64) {

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
