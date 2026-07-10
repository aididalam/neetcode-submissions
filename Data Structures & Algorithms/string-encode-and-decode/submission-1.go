type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	res:= ""

	for _,str:= range strs{
		res+= strconv.Itoa(len(str))+"#"+str
	}

	return res
}

func (s *Solution) Decode(encoded string) []string {
	sol := []string{}
	count := 0

	for i := 0; i < len(encoded); i++ {
		ch := encoded[i]

		if ch != '#' {
			num := int(ch - '0')
			count = count*10 + num
			continue
		}

		str := ""

		for j := 0; j < count; j++ {
			str += string(encoded[i+1+j])
		}

		sol = append(sol, str)

		
		i = i + count
		count = 0
	}

	return sol
}
