package stack

import "errors"

/*
栈 是一种 先进后出(FILO) 的数据结构
允许插入和删除的一端，为变化的一端，称为栈顶(Top)，另一端为固定的一端，称为栈底(Bottom)。
*/

const maxSize = 5

var (
	stackFull  = errors.New("stack full")
	stackEmpty = errors.New("stack empty")
)

type Stack interface {
	Push(val int) error // 添加元素, 入栈
	Pop() (int, error)  // 推出元素, 出栈
	IsFull() bool       // 是否为满
	IsEmpty() bool      // 是否为空
	Print()             // 打印
}
