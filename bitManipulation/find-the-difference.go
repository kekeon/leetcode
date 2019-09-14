package bitManipulation

func findTheDifference(s string, t string) byte {
	sumS := 0
	for _, v := range s {
		sumS += int(v)
	}

	sumT := 0
	for _, v := range t {
		sumT += int(v)
	}

	return byte(sumT - sumS)

}
