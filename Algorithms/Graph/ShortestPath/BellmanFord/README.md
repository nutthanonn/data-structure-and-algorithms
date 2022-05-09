# BellmanFord Algorithm

> BellmanFord Golang

### Single Source Shortest Path

- Algorithm ของ BellmanFord เกิดขึ้นมาเพื่อแก้ไขปัญหา path ที่มีค่าติดลบ ที่ Dijkstra's Algorithm ไม่สามารถแก้ได้แต่ BellmanFord มีปัญหาตรงที่ ถ้าเกิด Cycle ใน graph แล้วผลรวม Weight ของ Cycle นั้น มีค่าติดลบ จะทำให้เกิดปัญหาขึ้น แต่ถ้าผลรวมเป็นบวกก็จะไม่มีปัญหา

`หลักการ`

_1_ ให้จุดเริ่มต้นเป็น 0 และทุกจุดมี distance เป็น Inf

_2_ จับ Egde มาเป็นคู่ๆ ถ้า distance ของตัวแรก + weight ของตัวถัดไป แล้วมีค่าน้อยกว่า distance ของตัวถัดไป ให้ distance ของตัวถัดไป เก็บค่า distance ของตัวก่อนหน้่า + weight

## Example

```golang
//solution
.
.
.
if dist[u] != int(math.Inf(1)) && dist[u]+w <  dist[v] {
		dist[v] = dist[u] + w
}
```

_3_ ทำซ้ำข้อ 1-3 จนกว่าจะครบจำนวน Vertex - 1
