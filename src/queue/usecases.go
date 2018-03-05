package queue

func RadixSort(items []int) []int {
	var result []int
	var digitBins []*Queue
	place := 1
	mainBin := new(Queue)
	for i := 0; i < 10; i++ {
		digitBins = append(digitBins, new(Queue))
	}
	digitCount := countDigits(items[0])
	for _, item := range items {
		mainBin.Enqueue(item)
	}
	for place <= digitCount {
		for !mainBin.IsEmpty() {
			item := mainBin.Dequeue().GetData().(int)
			digit := getNthPlaceDigit(item, place)
			digitBin := digitBins[digit]
			digitBin.Enqueue(item)
		}
		for _, bin := range digitBins {
			for !bin.IsEmpty() {
				mainBin.Enqueue(bin.Dequeue().GetData().(int))
			}
		}
		place = place * 10
	}
	for !mainBin.IsEmpty() {
		result = append(result, mainBin.Dequeue().GetData().(int))
	}
	return result
}

func countDigits(input int) int {
	count := 1
	for input/10 > 0 {
		input = input / 10
		count = count * 10
	}
	return count
}

func getNthPlaceDigit(input, place int) int {
	count := 1
	var digit int
	for count <= place && input >= 0 {
		digit = input % 10
		input = input / 10
		count = count * 10
	}
	return digit
}
