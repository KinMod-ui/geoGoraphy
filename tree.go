package main

import (
	"errors"

	"github.com/KinMod-ui/geoGoraphy/util"
)

type vertex struct {
	x, y int
}

func (A *vertex) subVert(B *vertex) {
	A.x -= B.x
	A.y -= B.y
}

var dx = [4]int{1, 0, -1, 0}
var dy = [4]int{0, 1, 0, -1}

type quadTree struct {
	start, end vertex
	cnt        int
	children   [4]*quadTree
}

func getChar(start vertex, limit int) rune {
	if start.x < limit {
		if start.y < limit {
			return 'A'
		} else {
			return 'D'
		}
	} else {
		if start.y < limit {
			return 'B'
		} else {
			return 'C'
		}
	}

}

func sub(start, end, subVert vertex) (vertex, vertex) {
	start.subVert(&subVert)
	end.subVert(&subVert)
	return start, end
}

func getStringGivenVertex(start, end vertex, maxi int) []rune {

	var ret []rune
	maxLimit := maxi / 2

	for {
		lastChar := getChar(start, maxLimit)
		util.Mylog.Println(start, end, lastChar)
		if end.x-start.x+1 == maxLimit {
			return append(ret, lastChar)
		} else {
			switch lastChar {
			case 'A':
				{

				}
			case 'B':
				{
					start, end = sub(start, end, vertex{maxLimit - 1, 0})
				}

			case 'C':
				{
					start, end = sub(start, end, vertex{maxLimit - 1, maxLimit - 1})
				}
			case 'D':
				{
					start, end = sub(start, end, vertex{0, maxLimit - 1})
				}
			}
			ret = append(ret, lastChar)
		}
		maxLimit /= 2
	}
}

func printTree(tree *quadTree) {
	// util.Mylog.Println(tree.start, tree.end)

	for _, c := range tree.children {
		if c != nil {
			printTree(c)
		}
	}
}

func processTree(start vertex, end vertex, maxCnt int) *quadTree {
	cnt := CountNodes(start, end)
	util.Mylog.Println(start, end, cnt)

	tree := &quadTree{
		start: start,
		end:   end,
		cnt:   cnt,
	}

	if cnt <= maxCnt {
		return tree
	}

	xDiff := (end.x - start.x) / 2
	yDiff := (end.y - start.y) / 2

	tree.children[0] = processTree(start, vertex{start.x + xDiff, start.y + yDiff}, maxCnt)
	tree.children[1] = processTree(vertex{start.x, start.y + yDiff + 1},
		vertex{start.x + xDiff, start.y + 2*yDiff + 1}, maxCnt)
	tree.children[2] = processTree(vertex{start.x + xDiff + 1, start.y + yDiff + 1},
		vertex{start.x + 2*xDiff + 1, start.y + 2*yDiff + 1}, maxCnt)
	tree.children[3] = processTree(vertex{start.x + xDiff + 1, start.y},
		vertex{start.x + 2*xDiff + 1, start.y + yDiff}, maxCnt)

	return tree
}

func isOk(v vertex) bool {
	if v.x >= 0 && v.x < maxNode && v.y >= 0 && v.y < maxNode {
		return true
	}
	return false
}

func findUntilMaxPointsReach(start, end vertex, maxCnt int) [][2]vertex {
	h := initHeap()
	visited := make(map[[2]vertex]bool)
	cnt := 0

	h.Push([2]vertex{start, end})
	visited[[2]vertex{start, end}] = true

	ret := [][2]vertex{}

	for {
		q := initHeap()
		for {
			topAny := h.Pop()
			if topAny == nil {
				break
			}
			top := topAny.([2]vertex)
			cnt += CountNodes(top[0], top[1])
			ret = append(ret, top)
			// util.Mylog.Println(top, CountNodes(top[0], top[1]))
			if cnt >= maxCnt {
				break
			}

			for _, c := range find4Closest(top[0], top[1]) {
				_, ok := visited[c]
				if !ok {
					q.Push(c)
					visited[c] = true
				}
			}
		}
		if q.Len() == 0 {
			break
		}
		h = q
	}

	return ret
}

func find4Closest(start, end vertex) [][2]vertex {
	diff := end.x - start.x + 1
	var ret [][2]vertex
	for i := range 4 {
		ns := vertex{start.x + (dx[i] * diff), start.y + (dy[i] * diff)}
		ne := vertex{end.x + (dx[i] * diff), end.y + (dy[i] * diff)}
		if isOk(ns) && isOk(ne) {
			// util.Mylog.Println(ns, ne)
			val, err := findValidArea(ns, ne)
			if err != nil {
				continue
			}
			ret = append(ret, val)
		}
	}

	return ret
}

func findValidArea(start, end vertex) ([2]vertex, error) {
	temp := root
	for {
		flag := 1
		for _, child := range temp.children {
			// util.Mylog.Println(temp, child)
			if child == nil {
				continue
			}
			if compareRanges(start, end, child.start, child.end) {
				temp = child
				flag = 0
			}
		}
		if flag == 1 {
			break
		}
	}
	if temp.children[0] == nil {
		return [2]vertex{temp.start, temp.end}, nil
	} else {
		return [2]vertex{}, errors.New("children not nil")
	}
}

// returns true if s1 -> e1 lies in range s2 , e2
func compareRanges(s1, e1, s2, e2 vertex) bool {
	if s1.x >= s2.x && s1.y >= s2.y && e1.x <= e2.x && e1.y <= e2.y {
		return true
	}
	return false
}
