package main

import (
	"fmt"
)

type element struct {
	data int
	next *element
}

func newList(data int) element {
	return element{data, nil}
}

func printList(list *element) {
	for {
		fmt.Println(list.data)
		if list.next != nil {
			list = list.next
		} else {
			break
		}
	}
}

func addEndList(list *element, data int) {
	for {
		if list.next != nil {
			list = list.next
		} else {
			list.next = &element{data: data, next: nil}
			break
		}
	}
}

func addStartList(list element, data int) element {
	return element{data: data, next: &list}
}

func addAfter(list *element, index int, data int) {
	counter := 0
	for {
		if counter != index {
			if list.next != nil {
				counter++
				list = list.next
			} else {
				fmt.Println("object not found")
				break
			}

		} else {
			newElement := element{data: data, next: list.next}
			list.next = &newElement
			break
		}
	}

}

func main() {
	list := newList(12)
	addEndList(&list, 14)
	list = addStartList(list, 0)
	addAfter(&list, 0, 69)
	printList(&list)
}
