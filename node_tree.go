package avltree

type Node interface {
	GetData() int64
	GetLeft() Node
	GetRight() Node
	SetData(int64)
	SetLeft(Node)
	SetRight(Node)
}

type Tree interface {
	GetRoot() Node
	SetRoot(Node)
	Find(int64) Node
	FindMin() Node
	FindMax() Node
	Insert(int64)
	Delete(int64)
	PopMin() Node
	PopMax() Node
	Display()
}
