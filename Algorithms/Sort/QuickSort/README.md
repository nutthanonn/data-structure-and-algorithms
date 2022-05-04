# QuickSort

> Golang QuickSort

- QuickSort มี 4 ขั้นตอน

  - 1 if len(arr) < 2 break
  - 2 เลือกตัวเลขขึ้นมาหนึ่งตัว เรียกเลขตัวนี้ว่า Pivot
  - 3 แบ่ง Array เป็น 2 ก้อน ก้อนแรกทุกตัวจะมีค่าน้อยกว่า Pivot ก้อนสองทุกตัวจะมีค่ามากกว่า Pivot
  - 4 กลับไปทำขั้นตอนแรกใหม่กับ Array1 และ Array2

- Algorithmic Paradigm: `Divide and Conquer`
