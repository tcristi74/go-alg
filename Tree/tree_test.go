package BinaryTree

import (
	"testing"
	"fmt"
)




func TestIsBalancedTree(t *testing.T){

	head :=TreeNode{Val:100}
	l1 :=TreeNode{Val:50}
	l2 :=TreeNode{Val:25}
	l3 :=TreeNode{Val:10}
	l4 :=TreeNode{Val:60}
	r1 :=TreeNode{Val:105}
	r2 :=TreeNode{Val:110}
	r3 :=TreeNode{Val:125}
	r4 :=TreeNode{Val:80}
	r5 :=TreeNode{Val:75}
	r6 :=TreeNode{Val:30}
	l5 :=TreeNode{Val:103}
	l6 :=TreeNode{Val:5}

	head.Left = &l1
	head.Right = &r1

	l1.Left = &l2
	l1.Right =&r5

	r5.Left =&l4
	r5.Right = &r4

	l2.Left = &l3
	l2.Right = &r6

	r1.Right=&r2
	r1.Left= &l5

	l3.Left =&l4

	l4.Left = &l6

	r2.Right = &r3

	res :=IsBalanced(&head)

	if !res {
		t.Errorf("tree is not balanced")

	} else {
		t.Logf("tree is balanced")
		//fmt.Println("tree is balanced")
	}
}

func TestIsBalancedTree2(t *testing.T){

	fmt.Println("start test 2")
	head :=TreeNode{Val:100}
	l1 :=TreeNode{Val:50}
	l2 :=TreeNode{Val:25}
	l4 :=TreeNode{Val:60}
	r1 :=TreeNode{Val:105}
	r4 :=TreeNode{Val:80}
	r5 :=TreeNode{Val:75}

	head.Left = &l1

	l1.Left = &l2
	l1.Right =&r5

	r5.Left =&l4
	r5.Right = &r4

	head.Right = &r1


	res :=IsBalanced(&head)

	if !res {
		t.Errorf("tree is not balanced")

	} else {
		t.Logf("tree is balanced")
		//fmt.Println("tree is balanced")
	}
}
