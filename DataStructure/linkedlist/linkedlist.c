#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <stdbool.h>

struct Node{   
    int data;
    struct Node *next;
};


struct Node *head = NULL;


void Insert(int data){
    //create linked
    struct Node *link = (struct Node*)malloc(sizeof(struct Node));  //เหมือนสร้าง Node ใหม่
    link -> data = data; // ให้ Node ใหม่ เก็บ data  
    link -> next = head; // ให้ Node.next เก็บ adress ของ head
    head = link; // ให้ head เก็บ Node ใหม่
}

void getLinkedList(){
    struct Node *ptr = head;
    while (ptr != NULL){
        printf("%d\n", ptr->data);
        ptr = ptr -> next;
    }
}

int main() {
    Insert(2);
    Insert(4);
    Insert(6);
    Insert(8);
    getLinkedList();
}

