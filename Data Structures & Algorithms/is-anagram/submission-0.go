func isAnagram(s string, t string) bool {
		if len(s)!=len(t){
			return false
		}
		storeS:= make(map[string]int)
		storeT:= make(map[string]int)

		for idx,strS:= range s{
			strS:= string(strS)
			strT:= string(t[idx])
			
			countS,existS:= storeS[strS]
			if(existS){
				storeS[strS]=countS+1
			}else{
				storeS[strS]=1
			}


			countT,existT:= storeT[strT]
			if(existT){
				storeT[strT]=countT+1
			}else{
				storeT[strT]=1
			}

		} 

		for i,v:=range storeS{
			if v != storeT[i]{
				return false
			}
		}


	return true
}
