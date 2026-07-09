func isAnagram(s string, t string) bool {
		if len(s)!=len(t){
			return false
		}
		storeS:= make(map[string]int)
		storeT:= make(map[string]int)

		for idx,strS:= range s{
			strS:= string(strS)
			strT:= string(t[idx])
			
			
			storeS[strS]++
			storeT[strT]++

		} 

		for i,v:=range storeS{
			if v != storeT[i]{
				return false
			}
		}


	return true
}
