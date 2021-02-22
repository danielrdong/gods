package queue

import "errors"

/*
队列 是一种 先进先出(FIFO) 的数据结构
尾部插入，头部弹出
*/

const maxSize = 5

var (
	queueFull  = errors.New("queue full")
	queueEmpty = errors.New("queue empty")
)

type Queue interface {
	Push(value interface{}) error // 添加元素, 入队
	Pop() (interface{}, error)    // 推出元素, 出队
	Size() int                    // 队列大小
	IsFull() bool                 // 是否为满
	IsEmpty() bool                // 是否为空
	Print()                       // 打印
}
