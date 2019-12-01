package main

import "fmt"

func activityNotifications(expenditure []int32, d int32) int32 {
	numTrailDays := d
	expeditureLength := int32(len(expenditure))
	var median float32
	var notifications int32

	for day := int32(numTrailDays); day < expeditureLength; day++ {
		begin := day - numTrailDays
		fmt.Println("unsortedTrailingDays:", expenditure[begin:day])
		sortedTrailingDays := mergeSort(expenditure[begin:day])

		fmt.Println("sortedTrailingDays:", sortedTrailingDays)

		// calculate median spending
		if numTrailDays%2 == 0 {
			median = (float32(sortedTrailingDays[numTrailDays/2-1] + sortedTrailingDays[numTrailDays/2])) / 2.0
		} else {
			median = float32(sortedTrailingDays[numTrailDays/2])
		}
		fmt.Println("median:", median, "median*2", median*2)
		fmt.Printf("day is %d, amount is %d\n", day, expenditure[day])

		// determine if notification is warrented
		if float32(expenditure[day]) >= median*2 {
			notifications++
			fmt.Println("notification sent")
		}
		fmt.Println("")
	}

	return notifications
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

func main() {}
