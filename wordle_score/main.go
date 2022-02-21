package main

import (
	"fmt"
	"os/exec"
)

func main() {
	d := []int{0, 0, 0, 0, 0, 0}
	fmt.Println("Enter your wordle score distribution")

	for i := range d {
		score := i + 1
		fmt.Printf("%d: ", score)
		fmt.Scanln(&d[i])
	}

	games := Sum(d)
	points := Points(d)
	score := float64(points) / float64(games)
	fmt.Printf("\n\nDistribution\n------------\n")
	PrintDistribution(d)
	fmt.Printf("Total Games: %d\n", Sum(d))
	fmt.Printf("Total Points: %d\n", Points(d))
	fmt.Printf("Score: %.2f\n", score)

	exec.Command("pause")

}

func Sum(arr []int) (sum int) {
	for _, v := range arr {
		sum += v
	}
	return
}

func Points(arr []int) (points int) {
	for i, v := range arr {
		points += (i + 1) * v
	}
	return
}

func PrintDistribution(arr []int) {
	for i, v := range arr {
		fmt.Printf("%d | ", i+1)
		for i := 0; i < v; i++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
