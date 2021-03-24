package main

import (
	"fmt"
	"math"
)

func main() {
	result := isRectangle([][]int{{0, 0}, {2, 0}, {2, 1}, {0, 1}})
	fmt.Println(result)
}

func isRectangle(points [][]int) bool {
	point0 := points[0]
	point1 := points[1]
	point2 := points[2]
	point3 := points[3]
	side1 := checkLengthSide(point0, point1)
	side2 := checkLengthSide(point1, point2)
	side3 := checkLengthSide(point2, point3)
	side4 := checkLengthSide(point3, point0)
	if side1 != side3 || side2 != side4 {
		return false
	}

	vec1 := getVector(point0, point1)
	vec2 := getVector(point1, point2)
	vec3 := getVector(point2, point3)
	vec4 := getVector(point3, point0)
	if checkPerpendicular(vec1, vec2) && checkPerpendicular(vec2, vec3) && checkPerpendicular(vec3, vec4) && checkPerpendicular(vec4, vec1) {
		return true
	}
	return false
}

func checkPerpendicular(vec1, vec2 []int) bool {
	return vec1[0]*vec2[0]+vec1[1]*vec2[1] == 0
}

func getVector(point, point2 []int) []int {
	return []int{point[0] - point2[0], point[1] - point2[1]}
}

func checkLengthSide(point []int, point2 []int) float64 {
	l1 := math.Pow(float64(point[0]-point2[0]), 2)
	l2 := math.Pow(float64(point[1]-point2[1]), 2)
	return math.Sqrt(l1 + l2)
}
