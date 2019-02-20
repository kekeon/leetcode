package main

import "fmt"

func judgeCircle(moves string) bool {

	point := [2]int{0,0}

	for i := 0; i < len(moves);i++  {
		switch string(moves[i]) {

		case "U":
			point[1]++
			break
		case "D":
			point[1]--
			break
		case "R":
			point[0]++
			break
		case "L":
			point[0]--
			break
		}
	}

	if point[0] == 0 && point[1] == 0 {
		return true
	}

	return false
}

func main() {
	v := judgeCircle("")

	fmt.Println(v)
}
