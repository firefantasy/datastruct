package main

import "fmt"
import "strconv"

type Node struct {
	Val int
	Left *Node
	Right *Node
}

type binaryTree struct {
	root *Node
}

func newNode(val int) *Node {
	return &Node{Val: val}
}

func newTree() *binaryTree {
	return &binaryTree{}
}

func (t *binaryTree) Find(val int) []*Node {
	var ret []*Node
	root := t.root
	for root != nil {
		if root.Val == val {
			ret = append(ret, root)
			root = root.Right
		} else if root.Val < val {
			root = root.Right
		} else {
			root = root.Left
		}
	}

	return ret
}

func (t *binaryTree) Insert(val int) {
	root := t.root
	if root == nil {
		t.root = newNode(val)
		return
	}

	for {
		if root.Val <= val {
			if root.Right == nil {
				root.Right = newNode(val)
				return
			}
			root = root.Right
		} else {
			if root.Left == nil {
				root.Left = newNode(val)
				return
			}
			root = root.Left
		}
	}
}

func (t *binaryTree) Delete(val int) {
	for {
		var pp, p *Node
		p = t.root
		for (p != nil && p.Val != val) {
			pp = p
			if (p.Val > val) {
				p = p.Left
			} else {
				p = p.Right
			}
		}
		if p == nil {
			return
		}
	
		// 要删除的节点有左右节点
		// 1. 从被删除节点的右子树中寻找最小的节点
		// 2.
		if (p.Left != nil && p.Right != nil) {
			minP := p.Right
			minPP := p
			for minP.Left != nil {
				minPP = minP
				minP = minP.Left
			}
		    //  找到后替换当前节点
			p.Val = minP.Val
			p = minP
			pp = minPP
		}
		

		// 下面把要删除节点有一个子节点，没有子节点，两个子节点删除最小右子树节点合在一起写 

		// 删除的节点有子节点
		var child *Node
		if p.Left != nil {
			child = p.Left
		} else if p.Right != nil {
			child = p.Right
		}
		
		if pp == nil {
			// 删除节点恰好是根节点
			t.root = child
		} else if pp.Left == p {
			pp.Left = child
		} else {
			pp.Right = child
		}
	}
}

func inOrder(root *Node) []*Node{
	var ret []*Node
	if root == nil {
		return ret
	}

	if root.Left != nil {
		ret = inOrder(root.Left)
	}
	ret = append(ret, root)
	if root.Right != nil {
		ret = append(ret, inOrder(root.Right)...)
	}
	return ret
}

func main() {
	arr := []int{-2, 8, 8, 1, 1, 3, 1, 3,5,4,2,7}
	tree := newTree()
	for _, v := range arr {
		tree.Insert(v)
	}
	// for _, v := range tree.Find(2) {
	// 	fmt.Println(v.Val)
	// }
	for _, v := range inOrder(tree.root) {
		fmt.Print(strconv.Itoa(v.Val) + " ") 
	}
	tree.Delete(1)
	fmt.Print("\n")
	for _, v := range inOrder(tree.root) {
		fmt.Print(strconv.Itoa(v.Val) + " ")
	}
	tree.Delete(8)
	fmt.Print("\n")
	for _, v := range inOrder(tree.root) {
		fmt.Print(strconv.Itoa(v.Val) + " ")
	}
	fmt.Print("\n")
	tree.Insert(-2)
	tree.Insert(100)
	for _, v := range inOrder(tree.root) {
		fmt.Print(strconv.Itoa(v.Val) + " ")
	}
}