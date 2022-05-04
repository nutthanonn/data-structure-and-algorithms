# Binary Search Algorithm

> Golang Binary Search Algorithm

- Binary Search เป็น Algorithm ไว้ค้นหาข้อมูลสำหรับ Array ที่ถูก Sort ข้อมูลมาแล้ว

1. หาค่า mid-index
2. ถ้าค่า array[mid] > target ก็เปลี่ยนไปหาตั้งแต่ช่วง index mid ถึง len(array) - 1 ถ้าค่า array[mid] < target ก็เปลี่ยนไปหาช่วง index 0 ถึง mid-1

3. ทำซำ้ 1-2 จนกว่า array[mid] = target ให้ return index mid
