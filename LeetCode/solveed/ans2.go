package solveed

import (
	"fmt"
	"math/rand"

	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/models"
	"github.com/MoonyHsiao/leetCodeChallenge/LeetCode/models/sort"
)

func MaxRingProblem(nums []int) int {
	dataMap := make(map[int]int)
	traverseMap := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		dataMap[i+1] = nums[i]
		traverseMap[i+1] = false
	}
	isEnd := false
	maxCount := 0
	for !isEnd {
		index := 1
		for i := range traverseMap {
			isEnd = traverseMap[i]
			if !isEnd {
				index = i
				break
			}
		}

		if isEnd {
			break
		}

		set := GetSameSet(dataMap, dataMap[index])
		fmt.Printf("Set:%v\n", set)

		for i := range set {
			index := set[i]
			traverseMap[index] = true
		}

		if len(set) > maxCount {
			maxCount = len(set)
		}

	}
	return maxCount
}

func GetSameSet(m map[int]int, target int) map[int]int {
	res := make(map[int]int)
	res[target] = target
	temp := m[target]
	isEnd := false
	for !isEnd {
		_, ok := res[temp]
		if ok {
			isEnd = true
		} else {
			res[temp] = temp
			temp = m[temp]
		}
	}
	return res
}

func RobotSim(commands []int, obstacles [][]int) int {
	type point struct {
		x, y int
	}
	steps := [][]int{}
	setpU := []int{0, 1}
	steps = append(steps, setpU)
	setpR := []int{1, 0}
	steps = append(steps, setpR)
	setpD := []int{0, -1}
	steps = append(steps, setpD)
	setpL := []int{-1, 0}
	steps = append(steps, setpL)
	dirIndex, posX, posY, dist := 0, 0, 0, 0
	obstaclesMap := make(map[point]bool)
	for _, v := range obstacles {
		// temp := fmt.Sprintf("%v,%v", obstacles[i][0], obstacles[i][1])//這件事情會直接消耗50ms 兩次就是100ms
		// obstaclesMap[temp] = 0
		obstaclesMap[point{v[0], v[1]}] = true
	}
	for i := range commands {
		res := commands[i]
		if res == -1 {
			dirIndex = (dirIndex + 1) % 4

		} else if res == -2 {
			dirIndex = (dirIndex - 1) % 4
			if dirIndex < 0 {
				dirIndex += 4
			}
		} else {
			stepX, stepY := steps[dirIndex][0], steps[dirIndex][1]
			for i := 0; i < res; i++ {
				_posX, _posY := posX+stepX, posY+stepY
				// step := fmt.Sprintf("%v,%v", _posX, _posY)
				// _, cantGo := obstaclesMap[step]
				if obstaclesMap[point{_posX, _posY}] {
					break
				} else {
					posX, posY = _posX, _posY
				}

			}
			dist = models.Max(dist, posX*posX+posY*posY)
		}
	}
	return dist
}

func MaxCoins(piles []int) int {
	res := 0
	sorted := sort.MergeSort(piles)
	size := len(piles)
	sorted = sorted[size/3:]
	for i := len(sorted) - 1; i >= 0; i-- {
		if i%2 == 0 {
			res += sorted[i]
		}
	}
	return res
}

type Solution struct {
	list []int
}

func Constructor(nums []int) Solution {
	res := Solution{list: nums}
	return res
}

func (this *Solution) Pick(target int) int {
	count := 1
	index := -1
	getFirstNum := false
	k := 1
	for i := range this.list {
		if this.list[i] == target {
			if !getFirstNum {
				index = i
				getFirstNum = true
			} else {
				r := rand.Intn(count)
				if r < k {
					index = i
				}
			}
			count++
		}
	}
	return index
}

func NumIslands_DFS(grid [][]byte) int {

	res := 0
	if len(grid) == 0 {
		return res
	}
	m := len(grid)
	n := len(grid[0])
	//確認是否走過的陣列
	// fmt.Printf("m:%v\n", m)
	// fmt.Printf("n:%v\n", n)
	size := models.Max(m, n)
	mark := make([][]int, size, size)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mark[i] = append(mark[i], 0)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mark[i][j] == 0 && grid[i][j] == '1' {
				dfs(&mark, grid, i, j)
				res += 1
			}
		}
	}
	return res
}

func dfs(mark *[][]int, grid [][]byte, x, y int) {
	(*mark)[x][y] = 1
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	m := len(grid)
	n := len(grid[0])
	//  走上下左右的方向
	for i := 0; i < 4; i++ {
		newx := dx[i] + x
		newy := dy[i] + y
		if newx < 0 || newx >= m || newy >= n || newy < 0 {
			continue
		}
		if (*mark)[newx][newy] == 0 && grid[newx][newy] == '1' {
			dfs(mark, grid, newx, newy)
		}
	}
}

func FindDisappearedNumbers(nums []int) []int {
	res := []int{}
	size := len(nums)
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]] = nums[i]
	}
	for i := 1; i <= size; i++ {
		if m[i] != i {
			res = append(res, i)
		}
	}
	return res
}

type ProductOfNumbers struct {
	list []int
}

func ConstructorProductOfNumbers() ProductOfNumbers {
	res := ProductOfNumbers{}
	return res
}

func (this *ProductOfNumbers) Add(num int) {
	this.list = append(this.list, num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	res := 1
	size := len(this.list) - 1
	for k > 0 {
		res = res * this.list[size]
		k--
		size--
	}
	return res
}
