package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func activityNotifications(expenditure []int32, d int32) int32 {
	numTrailDays := d
	expeditureLength := int32(len(expenditure))
	var median float32
	var notifications int32
	rangeLength := int32(201) //0-200 = 201
	countArr := make([]int32, rangeLength)

	// prime the countArr
	for i := int32(0); i < numTrailDays; i++ {
		countArr[expenditure[i]]++
	}

	// begin
	for day := int32(numTrailDays); day < expeditureLength; day++ {

		mid := (numTrailDays / 2) + 1
		midMinus1 := mid - 1
		temp := int32(0)
		found := false

		if numTrailDays%2 == 0 {
			for index, value := range countArr {
				temp += value
				if temp >= midMinus1 && temp >= mid && !found {
					median = float32(index)
					break
				} else if temp >= midMinus1 && temp < mid && !found {
					median = float32(index)
					found = true
				} else if temp >= mid {
					median = float32((median + float32(index)) / 2)
					break
				}
			}
		} else {
			for index, value := range countArr {
				temp += value
				if temp >= mid {
					median = float32(index)
					break
				}
			}
		}

		// determine if notification is warrented
		if float32(expenditure[day]) >= median*2 {
			notifications++
		}

		// adjust count array
		deleteValue := expenditure[day-numTrailDays]
		countArr[deleteValue]--
		addValue := expenditure[day]
		countArr[addValue]++
	}

	return notifications
}

func countingSort(arr []int32) []int32 {
	fmt.Println("arr:", arr)
	rangeLength := int32(10) //0-200 = 201
	countArr := make([]int32, rangeLength)
	arrLength := int32(len(arr))

	for i := int32(0); i < arrLength; i++ {
		countArr[arr[i]]++
	}
	fmt.Println("countArr:", countArr)

	return fillFromCount(countArr, arr, arrLength, rangeLength)
}

func fillFromCount(countArr []int32, arr []int32, arrLength int32, rangeLength int32) []int32 {
	// sum the countArr
	for i := int32(1); i < rangeLength; i++ {
		countArr[i] = countArr[i-1] + countArr[i]
	}
	fmt.Println("countArr:", countArr)

	// fill in sortedArr
	sortedArr := make([]int32, arrLength)

	for _, value := range arr {
		if countArr[value] > 0 {
			sortedArr[countArr[value]-1] = value
			countArr[value]--
		}
	}
	fmt.Println("sortedArr:", sortedArr)

	return nil
}

func mergeSort(arr []int32) []int32 {
	arrLength := int32(len(arr))

	if arrLength == 1 {
		return arr
	}

	arrMiddle := arrLength / 2
	var left = make([]int32, arrMiddle)
	var right = make([]int32, arrLength-arrMiddle)

	for i := int32(0); i < arrLength; i++ {
		if i < arrMiddle {
			left[i] = arr[i]
		} else {
			right[i-arrMiddle] = arr[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int32) []int32 {
	result := make([]int32, len(left)+len(right))

	i := int32(0)

	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}

func main() {
	file, err := os.Open("test-case-5")
	checkError(err)

	reader := bufio.NewReaderSize(file, 1024*1024)

	stdout, err := os.Create("test-case-5-OUTPUT")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nd := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nd[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	dTemp, err := strconv.ParseInt(nd[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	expenditureTemp := strings.Split(readLine(reader), " ")

	var expenditure []int32

	for i := 0; i < int(n); i++ {
		expenditureItemTemp, err := strconv.ParseInt(expenditureTemp[i], 10, 64)
		checkError(err)
		expenditureItem := int32(expenditureItemTemp)
		expenditure = append(expenditure, expenditureItem)
	}

	result := activityNotifications(expenditure, d)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
