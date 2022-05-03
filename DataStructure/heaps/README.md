# Heaps

> Golang Heaps

- parent คือ ตำแหน่งหลักของ Node โดย index จะห่างกันอยู่ที่ _(index-1) / 2_

- left child คือ ตำแหน่งของ index ที่อยู่ทางด้านช้ายของ Heaps โดย index จะห่างกันอยู่ที่ _(2 x index) + 1_

- right child ตำแหน่งของ index ที่อยู่ทางด้านขวาของ Heaps โดย index จะห่างกันอยู่ที่ _(2 x index) + 2_

- heaps maxHeapiflyDown เป็นการ pop root node ออก จากนั้นหา node ที่มากที่สุดขึ้นมาเป็น root

- maxHeapifly เป็น function ที่ไว้ตรวจดูว่า node ที่เราเพิ่งจะ insert เข้าไปนั้น มันมีค่ามากกว่า parent node ของมันหรือไม่ ถ้ามีค่ามากกว่าก็จะเรียกใช้ function swap

- Extract เป็นการนำ root ของ Heaps ออกโดยจะนำ last index ของ array ขึ้นมาเป็น Root จากนั้นก็ทำการตรวจดูว่า child ตัวไหนมีค่ามากกว่าก็จะนำ node ตัวนั้นขึ้นมาเป็น root แทน
