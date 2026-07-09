func groupAnagrams(strs []string) [][]string {

	store:=make(map[[26]int][]string)

	for _,str:=range strs{
		count:=[26]int{}
		for _, v:= range str{
			count[v-'a']++
		}

		store[count]=append(store[count], str)
	}

	res:= [][]string{}

	for _,group:= range store{
		res=append(res,group)
	}

	return res

}
