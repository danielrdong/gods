package queue

import (
	"fmt"
)

/*
思路分析：
 head为队列头部(不包含), tail为队列尾部(包含)
1、初始化时：Head = -1, Tail = -1
2、什么时候队列满：Tail == MaxSize-1
3、什么时候队列空：Head == Tail
4、队列有多少元素：Tail - Head
*/

type SingleQueue struct {
	Array   [maxSize]interface{}
	MaxSize int
	Head    int
	Tail    int
}

func NewSingleQueue() Queue {
	return &SingleQueue{MaxSize: maxSize, Head: -1, Tail: -1}
}

func (sq *SingleQueue) Push(value interface{}) error {
	if sq.IsFull() {
		return queueFull
	}
	sq.Tail++ // Tail 后移
	sq.Array[sq.Tail] = value
	return nil
}

func (sq *SingleQueue) Pop() (interface{}, error) {
	if sq.IsEmpty() {
		return -1, queueEmpty
	}
	sq.Head++
	val := sq.Array[sq.Head]
	return val, nil
}

func (sq *SingleQueue) IsFull() bool {
	return sq.Tail == sq.MaxSize-1
}

func (sq *SingleQueue) IsEmpty() bool {
	return sq.Tail == sq.Head
}

func (sq *SingleQueue) Size() int {
	return sq.Tail - sq.Head
}

// 显示队列，找到队首，遍历到队尾
func (sq *SingleQueue) Print() {
	fmt.Println("当前队列的元素是：")
	for i := sq.Head + 1; i <= sq.Tail; i++ {
		fmt.Printf("Array[%d]=%d\t", i, sq.Array[i])
	}
	fmt.Println()
}
