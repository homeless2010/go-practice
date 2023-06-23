package main

import (
	"fmt"
)

type LinkNode struct {
	Data     int64
	NextNode *LinkNode
}

func main() {
	// var l  = LinkNode{
	//	Data:     1,
	//	NextNode: nil,
	//}
	var node *LinkNode = new(LinkNode)
	node.Data = 2
	var node1 *LinkNode = new(LinkNode)
	node1.Data = 3
	node.NextNode = node1

	node2 := new(LinkNode)
	node2.Data = 4
	node1.NextNode = node2

	node3 := new(LinkNode)
	node3.Data = 5
	node2.NextNode = node3

	//nowNode := node
	//for {
	//	if nowNode != nil {
	//		fmt.Println(nowNode.Data)
	//		nowNode = nowNode.NextNode
	//	} else {
	//		break
	//	}
	//}

	linkNode := reverse3(node)
	nowNode2 := linkNode
	for {
		if nowNode2 != nil {
			fmt.Println(nowNode2.Data)
			nowNode2 = nowNode2.NextNode
		} else {
			break
		}
	}

}

/**
递归
*/
func reverse(node *LinkNode) *LinkNode {
	if node.NextNode == nil {
		return node
	}
	last := reverse(node.NextNode)
	node.NextNode.NextNode = node
	node.NextNode = nil
	return last
}

func reverse2(node *LinkNode) *LinkNode {
	if node.NextNode == nil {
		return node
	}
	last := reverse(node.NextNode)
	node.NextNode.NextNode = node
	node.NextNode = nil
	return last
}

/**
栈
*/
func reverse3(head *LinkNode) *LinkNode {
	defer func() {
		fmt.Println("dddddddddddddddddddddddddddddd")
	}()
	stack := initStack()
	nowNode := head
	for {
		if nowNode != nil {
			//fmt.Println(nowNode.Data)
			stack.push(nowNode)
			nowNode = nowNode.NextNode
		} else {
			break
		}
	}
	p, _ := stack.pop()
	node := p.(*LinkNode)
	last := node
	for {
		n, _ := stack.pop()
		if n != 0 {
			node.NextNode = n.(*LinkNode)
			//fmt.Println(node.NextNode.Data)
			node = node.NextNode
		} else {
			break
		}
	}
	node.NextNode = nil
	return last
}
