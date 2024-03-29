package main

type MyCircularDeque struct {
	head  *Node
	tail  *Node
	len   int
	count int
}

type Node struct {
	Next *Node
	Pre  *Node
	Val  int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	head := Node{
		Next: nil,
		Pre:  nil,
		Val:  -1,
	}
	tail := Node{
		Next: nil,
		Pre:  nil,
		Val:  -1,
	}
	head.Next = &tail
	tail.Pre = &head
	deque := MyCircularDeque{
		head:  &head,
		tail:  &tail,
		len:   k,
		count: 0,
	}

	return deque
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	temp := this.head.Next
	tempNode := Node{
		Next: temp,
		Pre:  this.head,
		Val:  value,
	}
	this.head.Next = &tempNode
	temp.Pre = &tempNode
	this.count++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	temp := this.tail.Pre
	tempNode := Node{
		Next: this.tail,
		Pre:  temp,
		Val:  value,
	}
	this.tail.Pre = &tempNode
	temp.Next = &tempNode
	this.count++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	deleteTemp := this.head.Next
	this.head.Next = deleteTemp.Next
	deleteTemp.Next.Pre = this.head
	deleteTemp.Next, deleteTemp.Pre = nil, nil
	this.count--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	deleteTemp := this.tail.Pre
	deleteTemp.Pre.Next = this.tail
	this.tail.Pre = deleteTemp.Pre
	deleteTemp.Next, deleteTemp.Pre = nil, nil
	this.count--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	return this.head.Next.Val
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	return this.tail.Pre.Val
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.head.Next == this.tail
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.len == this.count
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
