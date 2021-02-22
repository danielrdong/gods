package stack

import "fmt"

type ArrayStack struct {
	MaxTop int          // 表示栈最大存放个数
	Top    int          // 表示栈顶
	arr    [maxSize]int // 数组模拟栈
}

func NewArrStack() Stack {
	return &ArrayStack{MaxTop: maxSize, Top: -1}
}

func (a *ArrayStack) Push(val int) error {
	if a.IsFull() {
		return stackFull
	}
	a.Top++
	a.arr[a.Top] = val
	return nil
}

func (a *ArrayStack) Pop() (int, error) {
	if a.IsEmpty() {
		return -1, stackEmpty
	}
	val := a.arr[a.Top]
	a.Top--
	return val, nil
}

func (a *ArrayStack) IsEmpty() bool {
	return a.Top == -1
}

func (a *ArrayStack) IsFull() bool {
	return a.Top == a.MaxTop-1
}

func (a *ArrayStack) Print() {
	if a.IsEmpty() {
		fmt.Println("stack empty")
		return
	}
	for i := a.Top; i >= 0; i-- {
		fmt.Printf("Array[%d]=%d\t", i, a.arr[i])
	}
	fmt.Println()
}
