package queue

import (
	"fmt"
)

/*
思路分析：
 head为队列头部(包含)，tail为队列尾部(不包含)
1、什么时候队列满：(Tail + 1) % MaxSize == Head
2、什么时候队列空：Tail == Head
3、队列有多少元素：(Tail + MaxSize - Head) % MaxSize
4、初始化时： Head = 0, Tail = 0
*/

type CycleQueue struct {
	Array   [maxSize]interface{}
	MaxSize int
	Head    int
	Tail    int
}

func NewCycleQueue() Queue {
	return &CycleQueue{MaxSize: maxSize, Head: 0, Tail: 0}
}

func (cq *CycleQueue) Push(value interface{}) error {
	if cq.IsFull() {
		return queueFull
	}
	cq.Array[cq.Tail] = value
	cq.Tail = (cq.Tail + 1) % cq.MaxSize
	return nil
}

func (cq *CycleQueue) Pop() (value interface{}, err error) {
	if cq.IsEmpty() {
		return -1, queueEmpty
	}
	value = cq.Array[cq.Head]
	cq.Head = (cq.Head + 1) % cq.MaxSize
	return
}

func (cq *CycleQueue) IsFull() bool {
	return (cq.Tail+1)%cq.MaxSize == cq.Head
}

func (cq *CycleQueue) IsEmpty() bool {
	return cq.Tail == cq.Head
}

func (cq *CycleQueue) Size() int {
	return (cq.Tail + cq.MaxSize - cq.Head) % cq.MaxSize
}

func (cq *CycleQueue) Print() {
	size := cq.Size()
	if size == 0 {
		fmt.Println("queue empty")
	}
	tempHead := cq.Head
	for i := 0; i < size; i++ {
		fmt.Printf("Array[%d]=%d\t", tempHead, cq.Array[tempHead])
		tempHead = (tempHead + 1) % cq.MaxSize
	}
	fmt.Println()
}
