package obj 

import (
)

type StackLocation struct {
    head *Node
    next *Node
}

type Node struct {
    data string
}

func  CreateStack() *StackLocation{

    return &StackLocation {
        head:nil,
        next:nil,
    }
}

func CreateNode(newdata string) *Node {
    return &Node {
        data:newdata,
    }
}

func (sl *StackLocation) Push(newnode *Node) *StackLocation {
    sl.next = sl.head
    sl.head = newnode

    return sl
}

func (sl *StackLocation) Pop() {
    sl.head = sl.next
}

func (sl *StackLocation) Topdata() string {
    return sl.head.data
}

func (sl *StackLocation) Wholedata() *StackLocation {
    return sl
}



