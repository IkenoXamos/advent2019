package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func calculateFuel(mass int) int {
	return (mass / 3 - 2)
}

func calculateFuel2(mass int) int {
	fuel := mass / 3 - 2
	if fuel <= 0 {
		return 0
	}

	return fuel + calculateFuel2(fuel)
}

func main() {
	f, err := os.Open("masses.dat")

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)
	sum := 0
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			fmt.Println("Error!")
		}
		sum += calculateFuel2(i)
	}

	fmt.Println(sum)

	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
