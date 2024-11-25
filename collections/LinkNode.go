package collections

type LinkNode struct {
	Pre *LinkNode
	Next *LinkNode
	Value interface{}
}

func NewLinkNode(Pre *LinkNode, Value interface{}, Next *LinkNode) (this *LinkNode) {
	this = new(LinkNode)
	this.Pre = Pre
	this.Value = Value
	this.Next = Next
	return
}
