### 队列

#### 介绍

队列 是一种 先进先出 的数据结构。  
由尾部插入，并从头部弹出。

#### 实现接口

```go
type Queue interface {
    Push(value interface{}) error // 添加元素, 入队
    Pop() (interface{}, error)    // 推出元素, 出队
    Size() int                    // 队列大小
    IsFull() bool                 // 是否为满
    IsEmpty() bool                // 是否为空
    Print()                       // 打印
}
```

#### 几种实现

- [由数组实现的单向队列](array_single.go)
- [由数组实现的循环队列](array_cycle.go)