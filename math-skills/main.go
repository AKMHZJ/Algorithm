package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Average(numbers []float64) float64 {
	var sum float64
	for _, number := range numbers {
		sum += number
	}

	if len(numbers) == 0 {
		return 0
	}
	return math.Round(sum / float64(len(numbers)))
}

func Median(numbers []float64) float64 {
	sort.Float64s(numbers)

	n := len(numbers)
	if n == 0 {
		return 0
	}

	if n%2 != 0 {
		return math.Round(numbers[n/2])
	}
	return math.Round((numbers[(n/2)-1] + numbers[n/2]) / 2)
}

func Variance(numbers []float64) float64 {
	n := len(numbers)
	if n == 0 {
		return 0
	}

	var sum float64
	for _, number := range numbers {
		sum += number
	}
	mean := sum / float64(n)

	var sumOfSquaredDiffs float64
	for _, number := range numbers {
		diff := number - mean
		sumOfSquaredDiffs += diff * diff
	}

	return math.Round(sumOfSquaredDiffs / float64(n))
}

func StandardDeviation(numbers []float64) float64 {
	vare := Variance(numbers)
	return math.Round(math.Sqrt(vare))
}

func main() {
	arg := os.Args[1:]

	if len(arg) != 1 {
		fmt.Println("Usage: go run [File.go] [Option.txt]")
		return
	}

	input := arg[0]

	data, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error while reading file")
		return
	}

	numbers := strings.Split(string(data), "\n")

	var nm []float64
	for _, v := range numbers {
		if v == "" || (v >= "a" && v <= "z") || (v >= "A" && v <= "Z") {
			continue
		}
		t, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			return
		}
		nm = append(nm, float64(t))
	}

	fmt.Printf("Average: %d\n", int(Average(nm)))
	fmt.Printf("Median: %d\n", int(Median(nm)))
	fmt.Printf("Variance: %d\n", int(Variance(nm)))
	fmt.Printf("Standard Deviation: %d\n", int(StandardDeviation(nm)))
}
