# TREE TRAVERSAL

> Golang TREE TRAVERSAL

### การเขียน In Pre Post เขียนเหมือนกันทุกประการ เปลี่ยนเเค่บรรทัดที่ append node.value

`Pre Order` :point_down:

```golang

func TRAVERSAL(n *Node){
    if n == nil {
        // return something
    }
    // append node value to array
    TRAVERSAL(n.Left)
    TRAVERSAL(n.Right)
}

```

`In Order` :point_down:

```golang
func TRAVERSAL(n *Node){
    if n == nil {
        // return something
    }
    TRAVERSAL(n.Left)
    // append node value to array
    TRAVERSAL(n.Right)
}
```

`Post Order` :point_down:

```golang
func TRAVERSAL(n *Node){
    if n == nil {
        // return something
    }
    TRAVERSAL(n.Left)
    TRAVERSAL(n.Right)
    // append node value to array
}
```

:point_up_2: จาก Code ด้านบน จะเห็นได้ชัดเจนเลยว่า in pre post-order มีตำเเหน่งการ append to array ที่ตามชื่อเลย

## Example TREE

<p align="center">
<img src="../../assets/in-pre-postBST.jpeg">
</p>
