

func hasDuplicate(nums []int) bool {
	store:= make(map[int]int)
    for i,v:= range nums{
		_,exists:= store[v]
		
		if exists {
			return true
		}

		store[v]=i
	
	}

	return false
}
