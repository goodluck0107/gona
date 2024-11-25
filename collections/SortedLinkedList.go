package collections

import "fmt"

type SortedLinkedList struct {
	Front *LinkNode
	Count int32
	compareFunc func(old, new interface{}) bool
	equalFunc func(old, new interface{}) bool
}

func NewSortedLinkedList(compare func(old, new interface{}) bool, equal func(old, new interface{}) bool) (this *SortedLinkedList) {
	this = new(SortedLinkedList)
	this.compareFunc = compare
	this.equalFunc = equal
	return
}

func (this *SortedLinkedList) Add(value interface{}) {
	newNode := NewLinkNode(nil, value, nil)
	if this.Front == nil {
		fmt.Println("头部赋值")
		this.Front = newNode
		this.Count = this.Count+1
		return
	}
	for element := this.Front; element != nil; element = element.Next {
		tempValue := element.Value
		if element.Pre == nil {//头部
			if this.compareFunc(tempValue, value) {//前面插入
				fmt.Println("头部 前面插入")
				this.Front = newNode
				newNode.Next = element
				element.Pre = newNode
				this.Count = this.Count+1
				break
			}else {
				if element.Next == nil {//尾部
					fmt.Println("头部 后面插入")
					element.Next = newNode
					newNode.Pre = element
					this.Count = this.Count+1
					break
				}
			}
		}else {
			if this.compareFunc(tempValue, value) {//前面插入
				fmt.Println("非头部 前面插入")
				element.Pre.Next = newNode
				newNode.Next = element
				element.Pre = newNode
				this.Count = this.Count+1
				break
			}else {
				if element.Next == nil {//尾部
					fmt.Println("非头部 后面插入")
					element.Next = newNode
					newNode.Pre = element
					this.Count = this.Count+1
					break
				}
			}
		}
	}
}

func (this *SortedLinkedList) Remove(value interface{}) (removeItem interface{} ) {
	if this.Front == nil {
		return
	}
	for element := this.Front; element != nil; element = element.Next {
		tempValue := element.Value
		if element.Pre == nil {//头部
			if this.equalFunc(tempValue, value) {//移除头部
				this.Count = this.Count-1
				if element.Next == nil {//没有其他元素了
					this.Front = nil
				}else {//还有其他元素
					this.Front = element.Next
					element.Next.Pre = nil
				}
				element.Next = nil
				element.Pre = nil
				removeItem = tempValue
				break
			}else {
				if element.Next == nil {
					break
				}
			}
		}else {
			if this.compareFunc(tempValue, value) {//不是移除头部
				this.Count = this.Count-1
				if element.Next == nil {//没有其他元素了
					element.Pre.Next = nil
				}else {//还有其他元素
					element.Pre.Next = element.Next
					element.Next.Pre = element.Pre
				}
				element.Next = nil
				element.Pre = nil
				removeItem = tempValue
				break
			}else {
				if element.Next == nil {
					break
				}
			}
		}
	}
	return
}

func (this *SortedLinkedList) Contains(value interface{}) (exists bool ) {
	if this.Front == nil {
		return
	}
	for element := this.Front; element != nil; element = element.Next {
		tempValue := element.Value
		if this.equalFunc(tempValue, value) {
			exists = true
			break
		}
	}
	return
}

func (this *SortedLinkedList) GetAll() (exists []interface{} ) {
	if this.Front == nil {
		return []interface{}{}
	}
	exists = make([]interface{}, this.Count, this.Count)
	index := 0
	for element := this.Front; element != nil; element = element.Next {
		tempValue := element.Value
		exists[index] = tempValue
		index = index+1
	}
	return
}
