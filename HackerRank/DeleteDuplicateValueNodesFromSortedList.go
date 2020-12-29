// https://www.hackerrank.com/challenges/delete-duplicate-value-nodes-from-a-sorted-linked-list/problem
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

type SinglyLinkedListNode struct {
    data int32
    next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
    head *SinglyLinkedListNode
    tail *SinglyLinkedListNode
}

func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
    node := &SinglyLinkedListNode {
        next: nil,
        data: nodeData,
    }

    if singlyLinkedList.head == nil {
        singlyLinkedList.head = node
    } else {
        singlyLinkedList.tail.next = node
    }

    singlyLinkedList.tail = node
}

func printSinglyLinkedList(node *SinglyLinkedListNode, sep string, writer *bufio.Writer) {
    for node != nil {
        fmt.Fprintf(writer, "%d", node.data)

        node = node.next

        if node != nil {
            fmt.Fprintf(writer, sep)
        }
    }
}

// Complete the removeDuplicates function below.

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */
func removeDuplicates(head *SinglyLinkedListNode) *SinglyLinkedListNode {
    set := make(map[int32]bool)
    list := []*SinglyLinkedListNode{}
    for head != nil {
        if _, containsCurr := set[head.data]; !containsCurr {
            list = append(list, head)
            set[head.data] = true
        }
        head = head.next
    }
    for i, curr := range list {
        if i != len(list) - 1 {
            curr.next = list[i + 1]
        } else {
            curr.next = nil
        }
    }
    return list[0]
}

func main() {
