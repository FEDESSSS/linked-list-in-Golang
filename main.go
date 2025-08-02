package main

import (
	"fmt"
)

type element struct {
	data int
	next *element
}

func NewList(data int) element {
	return element{data, nil}
}

func PrintList(list *element) {
	for {
		fmt.Println(list.data)
		if list.next != nil {
			list = list.next
		} else {
			break
		}
	}
}

func AddEndList(list *element, data int) {
	for {
		if list.next != nil {
			list = list.next
		} else {
			list.next = &element{data: data, next: nil}
			break
		}
	}
}

func AddStartList(list element, data int) element {
	return element{data: data, next: &list}
}

func AddAfter(list *element, index int, data int) {
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

func DeleteEndList(list *element) {
	for {
		if list.next != nil {
			list = list.next
			if list.next.next == nil {
				list.next = nil
				break
			}
		}

	}
}

func DeleteStartList(list *element) element {
	if list.next != nil {
		return *list.next
	} else {
		fmt.Println("object not found")
	}
	return *list
}

func DeleteAfterList(list *element, index int) {
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
			list.next = list.next.next
			break
		}
	}
}

func main() {
	list := NewList(12)
	AddEndList(&list, 14)
	list = AddStartList(list, 0)
	AddAfter(&list, 0, 69)
	PrintList(&list)
	fmt.Println("=====================")
	DeleteEndList(&list)
	PrintList(&list)
	fmt.Println("=====================")
	list = DeleteStartList(&list)
	PrintList(&list)
	fmt.Println("=====================")
	DeleteAfterList(&list, 0)
	PrintList(&list)
}
