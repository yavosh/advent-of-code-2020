package algo

import "fmt"

// https://www.techiedelight.com/subset-sum-problem/

// SubsetSum Subset Sum Problem
//           We exclude current item from subset and recur for remaining items.
//           Finally, we return true if we get subset by including or excluding current
//           item else we return false. The base case of the recursion would be when
//           no items are left or sum becomes negative.
//           We return true when sum becomes 0 i.e. subset is found.
func subsetSumRecursive(data []int, n int, sum int, acc []int) (bool, []int) {

	fmt.Printf("n:%d\n", n)

	if sum == 0 {
		return true, acc
	}

	if n < 0 || sum < 0 {
		return false, nil
	}

	// Case 1. include current item in the subset (A[n]) and recur
	// for remaining items (n - 1) with remaining sum (sum - A[n])
	include, acc1 := subsetSumRecursive(data, n-1, sum-data[n], append(acc, data[n]))

	// Case 2. exclude current item n from subset and recur for
	// remaining items (n - 1)
	exclude, acc2 := subsetSumRecursive(data, n-1, sum, acc)

	// return true if we can get subset by including or excluding the
	// current item
	if include {
		return include, acc1
	}

	if exclude {
		return exclude, acc2
	}

	return false, nil
}

func SubsetSum(data []int, sum int) (bool, []int) {
	return subsetSumRecursive(data, len(data)-1, sum, make([]int, 0))
}
