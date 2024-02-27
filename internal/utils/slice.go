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
	isDuplicateByValue := make(map[string]bool)

	for _, value := range values {
		_, found := isDuplicateByValue[value]

		if !found {
			isDuplicateByValue[value] = true
			continue
		}

		duplicatedValues = append(duplicatedValues, value)
	}

	return duplicatedValues
}

func DeduplicateValues(values []string) []string {
	deduplicateValues := []string{}
	isDuplicateByValue := make(map[string]bool)

	for _, value := range values {
		_, found := isDuplicateByValue[value]

		if !found {
			isDuplicateByValue[value] = true
			deduplicateValues = append(deduplicateValues, value)
			continue
		}
	}

	return deduplicateValues
}
