# Data Structure and Algorithm

> Goal

- [DataStructure](https://github.com/nutthanonn/data-structure-and-algorithm/tree/main/DataStructure)

  - Pointer
  - [Stack](https://github.com/nutthanonn/data-structure-and-algorithm/tree/main/DataStructure/stack)
  - [Queue](https://github.com/nutthanonn/data-structure-and-algorithm/tree/main/DataStructure/queue)
  - [LinkedList](https://github.com/nutthanonn/data-structure-and-algorithm/tree/main/DataStructure/linkedlist)
  - [Heap](https://github.com/nutthanonn/data-structure-and-algorithm/tree/main/DataStructure/heaps)
  - [Binary Search Tree](https://github.com/nutthanonn/data-structure-and-algorithm/tree/main/DataStructure/Tree%20%26%20Traversal)
    - [In-order](https://github.com/nutthanonn/data-structure-and-algorithm/tree/main/DataStructure/Tree%20%26%20Traversal/InOrderTree)
    - [Pre-order](https://github.com/nutthanonn/data-structure-and-algorithm/tree/main/DataStructure/Tree%20%26%20Traversal/PreOrderTree)
    - [Post-order](https://github.com/nutthanonn/data-structure-and-algorithm/tree/main/DataStructure/Tree%20%26%20Traversal/PostOrderTree)
    - [Level-order](https://github.com/nutthanonn/data-structure-and-algorithm/tree/main/DataStructure/Tree%20%26%20Traversal/LevelOrderTree)

- [Algorithm](https://github.com/nutthanonn/data-structure-and-algorithm/tree/main/Algorithms)
  - [Sort](https://github.com/nutthanonn/data-structure-and-algorithm/tree/main/Algorithms/Sort)
    - [Bubble Sort](https://github.com/nutthanonn/data-structure-and-algorithm/tree/main/Algorithms/Sort/BubbleSort) **O(n^2)**
    - [Insertion Sort](https://github.com/nutthanonn/data-structure-and-algorithm/tree/main/Algorithms/Sort/InsertionSort) **O(n^2)**
    - [Selection Sort](https://github.com/nutthanonn/data-structure-and-algorithm/tree/main/Algorithms/Sort/SelectionSort) **O(n^2)**
    - [Merge Sort](https://github.com/nutthanonn/data-structure-and-algorithm/tree/main/Algorithms/Sort/MergeSort) **O(nlogn)**
  - [Searching](https://github.com/nutthanonn/data-structure-and-algorithm/tree/main/Algorithms/Searching)
    - [Binary Search](https://github.com/nutthanonn/data-structure-and-algorithm/tree/main/Algorithms/Searching/BinarySearch)
    - Jump Search
    - [Linear Search](https://github.com/nutthanonn/data-structure-and-algorithm/tree/main/Algorithms/Searching/LinearSearch)
  - [Dynamic Programming](https://github.com/nutthanonn/data-structure-and-algorithm/tree/main/Algorithms/DynamicProgramming)
    - fibonacci problem
  - Tree & Graph
    - [Depth-First-Search](https://github.com/nutthanonn/data-structure-and-algorithm/tree/main/Algorithms/algo/DFS)
  - Reverse Array
    - [inplace](https://github.com/nutthanonn/data-structure-and-algorithm/tree/main/Algorithms/algo/inplace)

### NUTTHANON THONGCHAROEN PREORDER TREE TRAVERSAL

```golang
package main

import "fmt"

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Stack struct {
	s []string
}

func (this *Stack) PreOrder(n *Node) {
	if n == nil {
		return
	}

	this.s = append(this.s, n.Val)
	this.PreOrder(n.Left)
	this.PreOrder(n.Right)
}

func main() {
	T := &Node{Val: "N"}
	T.Left = &Node{Val: "U"}
	T.Left.Left = &Node{Val: "T"}
	T.Left.Right = &Node{Val: "N"}
	T.Left.Left.Left = &Node{Val: "T"}
	T.Left.Left.Right = &Node{Val: "O"}
	T.Left.Left.Left.Left = &Node{Val: "H"}
	T.Left.Left.Left.Right = &Node{Val: "N"}
	T.Left.Left.Left.Left.Left = &Node{Val: "A"}

	T.Right = &Node{Val: "T"}
	T.Right.Right = &Node{Val: "H"}
	T.Right.Left = &Node{Val: "H"}
	T.Right.Left.Right = &Node{Val: "C"}
	T.Right.Left.Left = &Node{Val: "O"}
	T.Right.Left.Left.Left = &Node{Val: "N"}
	T.Right.Left.Left.Right = &Node{Val: "G"}

	T.Right.Right.Right = &Node{Val: "R"}
	T.Right.Right.Left = &Node{Val: "A"}
	T.Right.Right.Right.Right = &Node{Val: "E"}
	T.Right.Right.Right.Left = &Node{Val: "O"}
	T.Right.Right.Right.Right.Right = &Node{Val: "N"}

	my_name := &Stack{}
	my_name.PreOrder(T)
	fmt.Println(my_name.s)
}

```
