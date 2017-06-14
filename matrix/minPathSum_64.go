package matrix

import (
	"fmt"
	"strconv"
)

//Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.
//Note: You can only move either down or right at any point in time.

type Point struct {
	x int
	y int
}

func (p *Point) ToString() string {
	return strconv.Itoa(p.x) + "|" + strconv.Itoa(p.y)
}

func minPathSum(grid [][]int) int {
	ht := make(map[string]int, 0)
	p := Point{0, 0}

	ht[p.ToString()] = grid[0][0]
	explorePoint(&grid, &ht, grid[0][0], Point{0, 1})
	explorePoint(&grid, &ht, grid[0][0], Point{1, 0})
	fmt.Println(ht)
	return 10
}

func explorePoint(grid *[][]int, ht *map[string]int, origDist int, newPoint Point) {

	if newPoint.x < 0 || newPoint.y < 0 || newPoint.x >= len(*grid) || newPoint.y >= len(*grid) {
		return
	}

	if newPoint.x == len(*grid)-1 && newPoint.x == len(*grid)-1 {
		return
	}
	fmt.Println("go in")

	//get the origin from ht
	//origDist := (*ht)[origin.ToString()]
	newN := (*grid)[newPoint.x][newPoint.y] + origDist
	n := -11
	//fmt.Printf("Origx:%d y=%d distance=%d\n",origin.x,origin.y,origDist)
	fmt.Printf("NewPointx:%d y=%d distance=%d\n", newPoint.x, newPoint.y, origDist)
	if n1, ok := (*ht)[newPoint.ToString()]; !ok {
		//write in ht
		(*ht)[newPoint.ToString()] = newN
	} else {
		n = n1
	}
	if n > newN {
		n = newN
		(*ht)[newPoint.ToString()] = newN
	}

	//search for existing target

	explorePoint(grid, ht, n, Point{newPoint.x, newPoint.y + 1})
	explorePoint(grid, ht, n, Point{newPoint.x + 1, newPoint.y})
	//explorePoint(grid, ht, newPoint, Point{newPoint.x - 1, newPoint.y})
	//explorePoint(grid, ht, newPoint, Point{newPoint.x, newPoint.y - 1})
}
