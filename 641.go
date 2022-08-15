package Question_by_day

type node struct {
	pre  *node
	next *node
	val  int
}

type MyCircularDeque struct {
	capacity int
	list     *node
}

func Constructor(k int) MyCircularDeque {
	list := new(node)
	list.pre = list
	list.next = list
	return MyCircularDeque{k, list}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.capacity == 0 {
		return false
	}
	this.capacity -= 1
	newnode := new(node)
	newnode.val = value
	newnode.next = this.list.next
	this.list.next.pre = newnode
	newnode.pre = this.list
	this.list.next = newnode

	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.capacity == 0 {
		return false
	}
	this.capacity -= 1
	newnode := new(node)
	newnode.val = value
	newnode.next = this.list
	newnode.pre = this.list.pre
	this.list.pre.next = newnode
	this.list.pre = newnode

	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.list.next == this.list {
		return false
	}
	this.list.next = this.list.next.next
	this.list.next.pre = this.list
	this.capacity += 1
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.list.pre == this.list {
		return false
	}
	this.list.pre = this.list.pre.pre
	this.list.pre.next = this.list
	this.capacity += 1
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.list.next == this.list {
		return -1
	} else {
		return this.list.next.val
	}
}

func (this *MyCircularDeque) GetRear() int {
	if this.list.pre == this.list {
		return -1
	} else {
		return this.list.pre.val
	}
}

func (this *MyCircularDeque) IsEmpty() bool {
	if this.list.next == this.list {
		return true
	} else {
		return false
	}
}

func (this *MyCircularDeque) IsFull() bool {
	if this.capacity == 0 {
		return true
	} else {
		return false
	}
}
