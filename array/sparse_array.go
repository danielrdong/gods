package sparse_array

import "fmt"

// 棋盘压缩 -> 二维数组压缩
const (
	chessRow, chessCol = 11, 11
	defaultVal         = 0
)

type (
	ChessMap [chessRow][chessCol]int

	SparseMap []ValNode
	ValNode   struct {
		Row int
		Col int
		Val int
	}
)

func (c ChessMap) Convert() SparseMap {
	var distMap SparseMap
	defaultNode := ValNode{Row: chessRow, Col: chessCol, Val: defaultVal}
	distMap = append(distMap, defaultNode)

	for i, v1 := range c {
		for j, v2 := range v1 {
			if v2 != defaultVal {
				node := ValNode{Row: i, Col: j, Val: v2}
				distMap = append(distMap, node)
			}
		}
	}
	return distMap
}

func (c ChessMap) Print() {
	for _, v1 := range c {
		for _, v2 := range v1 {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}

func (s SparseMap) Convert() ChessMap {
	var distMap ChessMap

	for i, v1 := range s {
		// 跳过第一行，chessMap规模信息
		if i != 0 {
			distMap[v1.Row][v1.Col] = v1.Val
		}
	}
	return distMap
}
