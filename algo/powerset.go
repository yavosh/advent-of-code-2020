package algo

// https://www.techiedelight.com/generate-power-set-given-set/

type combinationsResultType struct {
	combinations [][]int
}

func powerSetRecursive(data []int, n int, set []int, result *combinationsResultType) {

	if n == 0 {
		// workaround for not being able to append to a slice if we
		// want to consider it by reference
		result.combinations = append(result.combinations, set)
		return
	}

	withLast := append(set, data[n-1])
	powerSetRecursive(data, n-1, withLast, result)

	withoutLast := make([]int, len(set))
	copy(withoutLast, set)
	powerSetRecursive(data, n-1, withoutLast, result)

}

func PowerSet(data []int) [][]int {
	result := combinationsResultType{
		combinations: make([][]int, 0),
	}

	powerSetRecursive(data, len(data), []int{}, &result)
	return result.combinations
}
