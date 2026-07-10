func topKFrequent(nums []int, k int) []int {
	store := make(map[int]int)

	for _, v := range nums {
		store[v]++
	}

	// Each bucket contains multiple integers.
	bucket := make([][]int, len(nums)+1)

	for element, frequency := range store {
		index := len(nums) - frequency
		bucket[index] = append(bucket[index], element)
	}

	// Start with length 0 and capacity k.
	res := make([]int, 0, k)

	// Lower bucket indexes represent higher frequencies.
	for _, items := range bucket {
		for _, v := range items {
			res = append(res, v)

			if len(res) == k {
				return res
			}
		}
	}

	return res
}