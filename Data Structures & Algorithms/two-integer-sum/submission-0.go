func twoSum(nums []int, target int) []int {
    store:= make(map[int]int)

	for i,v:= range nums{
		j,exsit:=store[v]

		if exsit{
			return []int{j,i}
		}

		diff:= target - v
		
		store[diff]=i
	}
	return []int {0,0}
}
