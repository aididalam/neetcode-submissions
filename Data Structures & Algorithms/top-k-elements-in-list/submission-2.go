func topKFrequent(nums []int, k int) []int {
	store:= make(map[int]int)

	for _,v:= range nums{
		store[v]++
	}

	arr:= make([][]int, len(nums)+1)

	for num,ferq:= range store{
		arr[ferq]=append(arr[ferq],num)
	}

	res := []int{}

	for i:=len(arr)-1; i>0; i-- {
		for _,value:= range arr[i] {
			res=append(res, value)

			if len(res)==k{
				return res
			}
		} 
	}
	return res
}