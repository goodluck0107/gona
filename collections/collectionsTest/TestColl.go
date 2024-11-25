package main

import (
	"fmt"

	"gitee.com/andyxt/gona/collections"
)

type WordCount struct {
	Word  string
	Count int
}

func compareValue(old, new interface{}) bool {
	if new.(WordCount).Count >= old.(WordCount).Count {
		return true
	}
	return false
}
func equalValue(old, new interface{}) bool {
	if new.(WordCount).Word == old.(WordCount).Word {
		return true
	}
	return false
}
func main() {
	wordCounts := []WordCount{
		WordCount{"kate", 81},
		WordCount{"herry", 92},
		WordCount{"james", 81}}
	var aSortedLinkedList = collections.NewSortedLinkedList(compareValue, equalValue)
	for _, wordCount := range wordCounts {
		aSortedLinkedList.Add(wordCount)
	}
	fmt.Println("aSortedLinkedList Length:", aSortedLinkedList.Count)
	for element := aSortedLinkedList.Front; element != nil; element = element.Next {
		fmt.Println(element.Value.(WordCount))
	}

	aSortedLinkedList.Remove(WordCount{"kate", 81})
	aSortedLinkedList.Remove(WordCount{"herry", 81})
	aSortedLinkedList.Remove(WordCount{"james", 81})
	fmt.Println("afterRemove aSortedLinkedList Length:", aSortedLinkedList.Count)
	for element := aSortedLinkedList.Front; element != nil; element = element.Next {
		fmt.Println(element.Value.(WordCount))
	}

}
