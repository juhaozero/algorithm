package search

import (
	"fmt"
	"math"
)

// 定义一个点
type Point struct {
	x int
	y int
}

// 定义一个路径
type Path struct {
	points []Point
}

// 定义一个地图
type Map struct {
	width     int
	height    int
	start     Point
	end       Point
	obstacles []Point
}

// 定义一个节点
type Node struct {
	point  Point
	parent *Node
	g      int
	h      int
	f      int
}

// 判断点是否在地图内
func (m Map) isPointInMap(point Point) bool {
	if point.x >= 0 && point.x < m.width && point.y >= 0 && point.y < m.height {
		return true
	}
	return false
}

// 判断点是否是障碍物
func (m Map) isPointObstacle(point Point) bool {
	for _, obstacle := range m.obstacles {
		if obstacle.x == point.x && obstacle.y == point.y {
			return true
		}
	}
	return false
}

// 计算曼哈顿距离
func (m Map) manhattanDistance(point1 Point, point2 Point) int {
	return int(math.Abs(float64(point1.x-point2.x)) + math.Abs(float64(point1.y-point2.y)))
}

// 计算折线距离
func (m Map) euclideanDistance(point1 Point, point2 Point) int {
	return int(math.Sqrt(math.Pow(float64(point1.x-point2.x), 2) + math.Pow(float64(point1.y-point2.y), 2)))
}

// 计算某个点的邻居
func (m Map) getNeighbors(point Point) []Point {
	neighbors := []Point{
		Point{point.x - 1, point.y},
		Point{point.x + 1, point.y},
		Point{point.x, point.y - 1},
		Point{point.x, point.y + 1},
	}
	return neighbors
}

// 判断点是否在openList中
func isPointInOpenList(openList []Node, point Point) bool {
	for _, node := range openList {
		if node.point.x == point.x && node.point.y == point.y {
			return true
		}
	}
	return false
}

// 判断点是否在closeList中
func isPointInCloseList(closeList []Node, point Point) bool {
	for _, node := range closeList {
		if node.point.x == point.x && node.point.y == point.y {
			return true
		}
	}
	return false
}

// 寻路算法
func (m Map) findPath() Path {
	openList := []Node{}
	closeList := []Node{}
	// 将起点加入openList
	startNode := Node{
		point:  m.start,
		parent: nil,
		g:      0,
		h:      m.manhattanDistance(m.start, m.end),
		f:      m.manhattanDistance(m.start, m.end),
	}
	openList = append(openList, startNode)
	// 循环查找
	for len(openList) > 0 {
		// 从openList中取出f值最小的节点
		minFNode := openList[0]
		minFIndex := 0
		for i, node := range openList {
			if node.f < minFNode.f {
				minFNode = node
				minFIndex = i
			}
		}
		// 将取出的节点从openList中移除，加入closeList
		openList = append(openList[:minFIndex], openList[minFIndex+1:]...)
		closeList = append(closeList, minFNode)
		// 如果取出的节点是终点，则结束查找
		if minFNode.point.x == m.end.x && minFNode.point.y == m.end.y {
			break
		}
		// 获取取出节点的邻居
		neighbors := m.getNeighbors(minFNode.point)
		for _, neighbor := range neighbors {
			// 判断邻居是否在地图内，是否是障碍物，是否在closeList中
			if m.isPointInMap(neighbor) && !m.isPointObstacle(neighbor) && !isPointInCloseList(closeList, neighbor) {
				// 计算邻居的g、h、f值
				g := minFNode.g + m.manhattanDistance(minFNode.point, neighbor)
				h := m.manhattanDistance(neighbor, m.end)
				f := g + h
				// 将邻居加入openList
				neighborNode := Node{
					point:  neighbor,
					parent: &minFNode,
					g:      g,
					h:      h,
					f:      f,
				}
				if !isPointInOpenList(openList, neighbor) {
					openList = append(openList, neighborNode)
				}
			}
		}
	}
	// 根据closeList中的节点，构建路径
	path := Path{}
	node := closeList[len(closeList)-1]
	for node.parent != nil {
		path.points = append(path.points, node.point)
		node = *node.parent
	}
	path.points = append(path.points, m.start)
	// 将路径反转
	for i := 0; i < len(path.points)/2; i++ {
		j := len(path.points) - i - 1
		path.points[i], path.points[j] = path.points[j], path.points[i]
	}
	return path
}
func FindPath() {
	// 初始化地图
	maps := Map{
		width:     2560,
		height:    2560,
		start:     Point{2324, 608},
		end:       Point{2489, 443},
		obstacles: []Point{},
	}
	// 寻路

	path := maps.findPath()

	// 打印路径
	fmt.Println("path", path.points)
	//fmt.Println("path", len(path.points))
}
