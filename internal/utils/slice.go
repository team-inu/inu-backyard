package slice

func Subtraction(setA []string, setB []string) []string {
	results := []string{}

	existedByValueB := make(map[string]bool, len(setB))

	for _, valueB := range setB {
		existedByValueB[valueB] = true
	}

	for _, valueA := range setA {
		_, found := existedByValueB[valueA]
		if !found {
			results = append(results, valueA)
		}
	}

	return results
}

func GetDuplicateValue(values []string) []string {
	duplicatedValues := []string{}

	duplicateCountByValue := make(map[string]int)

	for _, value := range values {
		count, found := duplicateCountByValue[value]

		if !found {
			duplicateCountByValue[value] = 1
			continue
		}

		if count == 1 {
			duplicatedValues = append(duplicatedValues, value)
		}

		count++
	}

	return duplicatedValues
}
