package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (l *LinkedList) Insert(val int) {
	n := Node{}      // สร้าง Node ใหม่
	n.value = val    //เอา data เข้ามาเก็บใน Node
	if l.size == 0 { //ถ้าเป็นตัวเเรกใน linkedlist
		l.head = &n //head เก็บ Adress ของ node
		l.size++    //เพิ่มขนาด linkedlist
		return      //จบการทำงาน
	}

	ptr := l.head                 // สร้างตัวเเปร ptr เก็บ head ของ linkedlist
	for i := 0; i < l.size; i++ { // วนตามขนาด linkedlist
		if ptr.next == nil { //ถ้าวนถึงตัวสุดท้ายของ list ใน linked list
			ptr.next = &n // ให้ตัวต่อไปเก็บ Adress ของ Node ต่อไป
			l.size++      // เพิ่มขนาด
			return        //จบการทำงาน
		}
		ptr = ptr.next // ให้ ptr เป็น node ตัวถัดไปเรื่อยๆ
	}
}

func (l *LinkedList) DataList() {
	ptr := l.head
	for i := 0; i < l.size; i++ {
		fmt.Println(ptr.value)
		ptr = ptr.next
	}
}

func (l *LinkedList) Search(val int) int {
	ptr := l.head //ให้ ptr เก็บ head
	for i := 0; i < l.size; i++ {
		if ptr.value == val { //วนหา ถ้า value == ค่าที่ส่งเข้ามา
			return i
		}
		ptr = ptr.next // ให้ ptr เป็น node ตัวถัดไปเรื่อยๆ
	}
	return -1
}

func (l *LinkedList) PositionAtListNode(position int) *Node {
	ptr := l.head
	if position < 0 { //position น้อยกว่า 0
		return ptr
	}
	if position > l.size-1 { // position มากกว่า size of linkedlist
		return nil
	}
	for i := 0; i < position; i++ { // วนตาม position
		ptr = ptr.next
	}

	return ptr
}

func (l *LinkedList) DeleteData(position int) {
	if position < 0 {
		fmt.Println("Delete position must be positive")
	}

	if l.size == 0 {
		fmt.Println("No node in linkedlist")
	}

	prevNode := l.PositionAtListNode(position - 1)
	if prevNode == nil {
		fmt.Println("Node not found")
	}

	prevNode.next = l.PositionAtListNode(position).next
	l.size--
}

func main() {
	list := LinkedList{}

	// เพื่มข้อมูลใน linkedlist
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)

	list.DeleteData(2)                      // Delete ตำเเหน่งที่ 2
	list.DataList()                         // เเสดง data ใน linkedlist
	fmt.Println(list.PositionAtListNode(1)) // ถามว่า position 1 เก็บ node อะไร
}

// ref : https://levelup.gitconnected.com/go-singly-linked-lists-with-insertion-deletion-traversal-e9da5bb0dbe1
