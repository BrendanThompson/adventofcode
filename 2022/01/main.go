package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var inputFilePath string
	var element int
	var newLineCount int
	slicedData := [][]int{}
	elementData := []int{}
	totaledData := []int{}

	flag.StringVar(&inputFilePath, "input", "./2022/01/input", "input file path")
	flag.Parse()

	inputData, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}

	var splitData = bytes.Split(inputData, []byte("\n"))

	for _, data := range splitData {
		dataString := string(data)

		if dataString == "" {
			slicedData = append(slicedData, elementData)
			element++
			newLineCount++
			elementData = []int{}
			continue
		} else {
			dataInt, _ := strconv.Atoi(dataString)
			elementData = append(elementData, dataInt)
		}
	}

	for _, slice := range slicedData {
		sliceTotal := total(slice)
		totaledData = append(totaledData, sliceTotal)
	}

	sort.Ints(totaledData)

	topThreeCalories := totaledData[len(totaledData)-1] + totaledData[len(totaledData)-2] + totaledData[len(totaledData)-3]

	fmt.Printf("Part 1: %d\n", totaledData[len(totaledData)-1])
	fmt.Printf("Part 2: %d", topThreeCalories)
}

func total(in []int) int {
	var total int

	for _, i := range in {
		total = total + i
	}

	return total
}
