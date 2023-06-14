package main

import "fmt"

func averageMark(gradebook []float64) float64 {
	var sum float64

	for i := 0; i < len(gradebook); i++ {
		sum += gradebook[i]

	}
	averageMark := sum / 5

	return averageMark

}

func main() {
	gradebook := make([]float64, 0)

	fmt.Println("Enter your marks")

	var mark float64 = 0

	for i := 0; i < 5; i++ {
		fmt.Scan(&mark)
		gradebook = append(gradebook, mark)
	}
	fmt.Println("Your grades:", gradebook)

	fmt.Println("Your average mark is: ", averageMark(gradebook))

}
