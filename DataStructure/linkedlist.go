package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func (l *LinkedList) Insert(val int) {
	n := Node{}     // สร้าง Node ใหม่
	n.value = val   //เอา data เข้ามาเก็บใน Node
	if l.len == 0 { //ถ้าเป็นตัวเเรกใน linkedlist
		l.head = &n //head เก็บ Adress ของ node
		l.len++     //เพิ่มขนาด linkedlist
		return      //จบการทำงาน
	}

	ptr := l.head                // สร้างตัวเเปร ptr เก็บ head ของ linkedlist
	for i := 0; i < l.len; i++ { // วนตามขนาด linkedlist
		if ptr.next == nil { //ถ้าวนถึงตัวสุดท้ายของ list ใน linked list
			ptr.next = &n // ให้ตัวต่อไปเก็บ Adress ของ Node ต่อไป
			l.len++       // เพิ่มขนาด
			return        //จบการทำงาน
		}
		ptr = ptr.next // ทำให้มันยังเชื่อมต่อ list กันเหมือนเดิม
	}
}

func (l *LinkedList) Search(val int) int {
	ptr := l.head //ให้ ptr เก็บ head
	for i := 0; i < l.len; i++ {
		if ptr.value == val { //วนหา ถ้า value == ค่าที่ส่งเข้ามา
			return i
		}
		ptr = ptr.next // ทำให้มันยังเชื่อมต่อ list กันเหมือนเดิม
	}
	return -1
}

func main() {
	list := LinkedList{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	fmt.Println(list.Search(3))
}
