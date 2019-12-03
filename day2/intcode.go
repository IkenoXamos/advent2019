package  main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"time"
)

func process(list []int) []int {
	index := 0
	for opcode := list[index]; opcode != 99; opcode = list[index] {
		switch opcode {
		case 1:
			if(list[index + 1] < 0 || list[index + 2] < 0 || list[index + 3] < 0) {
				// fmt.Println("Error: Negative position")
				return list
			}
			if(list[index + 1] > len(list) || list[index + 2] > len(list) || list[index + 3] > len(list)) {
				// fmt.Println("Error: Too large position")
				return list
			}

			list[list[index + 3]] = list[list[index + 1]] + list[list[index + 2]]
		case 2:
			if(list[index + 1] < 0 || list[index + 2] < 0 || list[index + 3] < 0) {
				// fmt.Println("Error: Negative position")
				return list
			}
			if(list[index + 1] > len(list) || list[index + 2] > len(list) || list[index + 3] > len(list)) {
				// fmt.Println("Error: Too large position")
				return list
			}

			list[list[index + 3]] = list[list[index + 1]] * list[list[index + 2]]
		case 99:
			fmt.Println("Process Complete.")
		default:
			// fmt.Println("Error: Invalid Opcode " + string(opcode))
		}

		index += 4
	}

	return list
}

func runProgram(list []int, noun int, verb int) int {
	list[1] = noun
	list[2] = verb

	return process(list)[0]
}

func main() {
	start := time.Now()

	filehandle, _ := os.Open("codes.dat")
	defer filehandle.Close()

	scanner := bufio.NewScanner(filehandle)

	scanner.Scan()

	codes := scanner.Text()

	list := strings.Split(codes, ",")

	intlist := make([]int, 0)

	for i := 0; i < len(list); i++ {
		val, _ := strconv.Atoi(list[i])
		intlist = append(intlist, val)
	}

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			cpy := make([]int, len(intlist))
			copy(cpy, intlist)
			if runProgram(cpy, i, j) == 19690720 {
				fmt.Println(100 * i + j)
			}
		}
	}

	elapsed := time.Since(start)

	fmt.Println(elapsed)
}
