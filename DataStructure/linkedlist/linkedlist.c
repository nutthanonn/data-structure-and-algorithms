#include <stdio.h>
#include <stdlib.h>


typedef struct node
{
    int val;
    struct node *next;
} node_t;


int getSize(node_t *head)
{
    node_t *curr = head;
    int size = 0;
    while (curr != NULL)
    {
        size++;
        curr = curr->next;
    }
    return size;
}


void push(node_t *head, int val)
{
    node_t *curr = head;
    while (curr->next != NULL)
        curr = curr->next;

    node_t *new_node = (node_t*)malloc(sizeof(node_t));
    new_node->val = val;
    new_node->next = NULL;
    curr->next = new_node;
}


void getLinkedlist(node_t *head)
{
    node_t *curr = head;
    while (curr != NULL)
    {
        printf("%d\n", curr->val);
        curr = curr->next;
    }
}


void InsertIndex(node_t *head, int val, int pos)
{
    if(pos < 0 || pos > getSize(head))
        return;

    node_t *curr = head;
    for(int i=0;i<pos-1;i++)
        curr = curr->next;

    node_t *new_node = (node_t*)malloc(sizeof(node_t));
    new_node->val = val;
    new_node->next = curr->next;
    curr->next = new_node;
}

void delIndex(node_t *head, int pos)
{
    node_t *curr = head;
    for(int i=0;i<pos-1;i++)
        curr = curr->next;

    curr->next = curr->next->next;
}



int main()
{
    int number[5] = {1,2,3,4,5};
    node_t *head = NULL;
    head = (node_t*)malloc(sizeof(node_t));
    head->val = 0;

    for(int i=0;i<5;i++)
        push(head, number[i]);


    getLinkedlist(head);
    printf("------\n");
    InsertIndex(head, 99, 7);
    getLinkedlist(head);
    printf("------\n");


    delIndex(head, 1);
    getLinkedlist(head);


    return 0;
}
