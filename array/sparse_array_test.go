package sparse_array

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	var c1Map ChessMap
	c1Map[1][2] = 4
	c1Map[4][6] = 9
	c1Map.Print()
	fmt.Println("---------------")
	sMap := c1Map.Convert()
	fmt.Println(sMap)
	fmt.Println("---------------")
	c2Map := sMap.Convert()
	c2Map.Print()

}