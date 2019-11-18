package avltree

type AVLTree struct {
	BinaryTree
	root Node
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
