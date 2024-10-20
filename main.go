package main

import (
	"fmt"
	"math/rand"

	"github.com/KinMod-ui/geoGoraphy/util"
)

const maxNode = 8

var grid [maxNode][maxNode]int

var root *quadTree

func main() {

	for i := 0; i < maxNode; i++ {
		x := rand.Intn(maxNode)
		y := rand.Intn(maxNode)
		grid[x][y] = 1
	}

	// for i := 0; i < 8; i++ {
	// 	for j := 0; j < 8; j++ {
	// 		if i >= 4 && j < 4 {
	// 			grid[i][j] = rand.Intn(2)
	// 		}
	// 	}
	// }

	for i := 0; i < maxNode; i++ {
		for j := 0; j < maxNode; j++ {
			fmt.Printf("%d ", grid[i][j])
		}
		util.Mylog.Println()
	}
	root = processTree(vertex{0, 0}, vertex{maxNode - 1, maxNode - 1}, 1)

	printTree(root)

	util.Mylog.Println(findUntilMaxPointsReach(vertex{4, 4}, vertex{5, 5}, 10))

}
