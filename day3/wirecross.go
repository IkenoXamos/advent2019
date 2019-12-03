package  main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func signalDistance(stepgrid1, stepgrid2 [][]int, i, j int) int {
	return stepgrid1[j][i] + stepgrid2[j][i]
}

func manhattanDistance(i, j, x, y int) int {
	return abs(x - i) + abs(y - j)
}

func main() {

	SIZE := 20000
	filehandle, _ := os.Open("wires.dat")

	defer filehandle.Close()
	scanner := bufio.NewScanner(filehandle)

	scanner.Scan()

	wire1 := strings.Split(scanner.Text(), ",")
	scanner.Scan()

	wire2 := strings.Split(scanner.Text(), ",")

	grid := make([][]string, SIZE)
	for i := range grid {
		grid[i] = make([]string, SIZE)
		for j := range grid[i] {
			grid[i][j] = ""
		}
	}

	stepgrid1 := make([][]int, SIZE)
	for i := range stepgrid1 {
		stepgrid1[i] = make([]int, SIZE)
	}

	stepgrid2 := make([][]int, SIZE)
	for i := range stepgrid2 {
		stepgrid2[i] = make([]int, SIZE)
	}


	fmt.Println("Laying Wire 1")
	x, y := SIZE / 2, SIZE / 2
	steps := 0;
	for i := range wire1 {
		direction := wire1[i][0]
		distance, _ := strconv.Atoi(wire1[i][1:])

		switch direction {
		case 'R':
			fmt.Println("Moving right " + strconv.Itoa(distance))
			for i := x; i < x + distance; i++ {
				grid[y][i] = "O"
				stepgrid1[y][i] = steps
				steps++
			}
			x += distance
		case 'L':
			fmt.Println("Moving left " + strconv.Itoa(distance))
			for i := x; i > x - distance; i-- {
				grid[y][i] = "O"
				stepgrid1[y][i] = steps
				steps++
			}
			x -= distance
		case 'U':
			fmt.Println("Moving up " + strconv.Itoa(distance))
			for i := y; i > y - distance; i-- {
				grid[i][x] = "O"
				stepgrid1[i][x] = steps
				steps++
			}
			y -= distance
		case 'D':
			fmt.Println("Moving down " + strconv.Itoa(distance))
			for i := y; i < y + distance; i++ {
				grid[i][x] = "O"
				stepgrid1[i][x] = steps
				steps++
			}
			y += distance
		default:
			fmt.Println("Something went wrong")
		}
	}

	fmt.Println("Laying Wire 2")
	x, y = SIZE / 2, SIZE / 2
	steps = 0
	for i := range wire2 {
		direction := wire2[i][0]
		distance, _ := strconv.Atoi(wire2[i][1:])

		switch direction {
		case 'R':
			fmt.Println("Moving right " + strconv.Itoa(distance))
			for i := x; i < x + distance; i++ {
				if grid[y][i] == "O" {
					grid[y][i] = "X"
				} else {
					grid[y][i] = "P"
				}
				stepgrid2[y][i] = steps
				steps++
			}
			x += distance
		case 'L':
			fmt.Println("Moving left " + strconv.Itoa(distance))
			for i := x; i > x - distance; i-- {
				if grid[y][i] == "O" {
					grid[y][i] = "X"
				} else {
					grid[y][i] = "P"
				}
				stepgrid2[y][i] = steps
				steps++
			}
			x -= distance
		case 'U':
			fmt.Println("Moving up " + strconv.Itoa(distance))
			for i := y; i > y - distance; i-- {
				if grid[i][x] == "O" {
					grid[i][x] = "X"
				} else {
					grid[i][x] = "P"
				}
				stepgrid2[i][x] = steps
				steps++
			}
			y -= distance
		case 'D':
			fmt.Println("Moving down " + strconv.Itoa(distance))
			for i := y; i < y + distance; i++ {
				if grid[i][x] == "O" {
					grid[i][x] = "X"
				} else {
					grid[i][x] = "P"
				}
				stepgrid2[i][x] = steps
				steps++
			}
			y += distance
		default:
			fmt.Println("Something went wrong")
		}
	}

	min := math.MaxInt64
	closestx, closesty := SIZE / 2, SIZE / 2
	for i := range grid {
		for j := range grid[i] {
			if i == SIZE / 2 && j == SIZE / 2 {
				continue
			}
			if(grid[j][i] == "X") {
				val := signalDistance(stepgrid1, stepgrid2, i, j)
				if val < min {
					min = val
					closestx = j
					closesty = i
				}
			}
		}
	}

	fmt.Println(min)
	fmt.Println(closestx)
	fmt.Println(closesty)
}
