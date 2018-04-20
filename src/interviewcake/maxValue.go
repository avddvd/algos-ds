package interviewcake

import "fmt"

func maxValue(weightValues [][]int, maxWeight int) int {
	result := map[int]int{}
	for currWeight := 0; currWeight <= maxWeight; currWeight++ {
		currMax := 0
		for _, weightValue := range weightValues {
			weight, value := weightValue[0], weightValue[1]
			if weight <= currWeight {
				residue := currWeight - weight
				residueValue := 0
				if _, ok := result[residue]; ok {
					residueValue = result[residue]
				}
				currValue := value + residueValue
				if currValue > currMax {
					currMax = currValue
				}
			}
		}
		result[currWeight] = currMax
	}
	fmt.Println(result)
	return result[maxWeight]
}
