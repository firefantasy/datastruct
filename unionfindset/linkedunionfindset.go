package unionfindset

// 组成双向链表
type linkedNode struct {
	R *linkedNode // 链表第一个节点，后续节点都指向它
	Next *linkedNode 
	Prev *linkedNode
}

func NewLinkedNodeSet(n int) *LinkedNodeSet {
	t := LinkedNodeSet{}
	for i := 0; i < n; i++ {
		n := &linkedNode{}
		n.Next = n
		n.Prev = n
		n.R = n
		t.nodes = append(t.nodes, n)
	}

	return &t
}

type LinkedNodeSet struct {
	nodes []*linkedNode
}

func(t *LinkedNodeSet) Find(i, j int) bool {
	return t.nodes[i].R == t.nodes[j].R
}

func (t *LinkedNodeSet) Union(i, j int) {
	if t.Find(i, j) {
		return 
	}
	iNode, jNode := t.nodes[i], t.nodes[j]
	iNodeR, jNodeR := iNode.R, jNode.R
	iNodeTail, jNodeTail := iNodeR.Prev, jNodeR.Prev

	iNodeTail.Next = jNodeR // i集合尾链接j集合头
	jNodeR.Prev = iNodeTail // j集合头链接i集合尾
							// i, j集合头尾链接

	jNodeTail.Next = iNodeR // j集合尾链接i集合头
	iNodeR.Prev = jNodeTail 
	// 上面链接两个双向链表操作

	// j集合的R都执行i集合的R
	p := jNodeR
	for p != jNodeTail {
		p.R = iNodeR.R
	}
	jNodeTail.R = iNodeR.R
}

