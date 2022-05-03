#include <stdio.h>
#include <stdlib.h>



typedef struct Node {
    int Val;
    struct Node * Next;
} node_t;

node_t * head = NULL;


void Insert(node_t** head,int val) {
    node_t * new_node = (node_t *)malloc(sizeof(node_t));
    node_t * current = *head;


    new_node->Val  = val;
    new_node->Next = NULL;

    if(*head == NULL) {
        *head = new_node;   
        return;
    }

    while (current -> Next != NULL) {
        current = current -> Next;
    }
    current->Next = new_node;
    return;
}



void Data() {
    struct Node *current = head;
    while (current != NULL) {   
        printf("%d\n", current->Val);
        current = current->Next;
    }
}



int main(int argc, char const *argv[]) {
    Insert(&head, 1);
    Insert(&head, 2);
    Insert(&head, 3);
    Insert(&head, 4);
    Data(); 
    

    return 0;
}
