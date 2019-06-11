package array

func commonChars(A []string) []string {

	resM := map[string]int{}

	for _, v := range A[0] {
		resM[string(v)] ++
	}


	for _, s := range A {
		m := map[string]int{}
		for _, k := range s {
			m[string(k)]++
		}

		for r, v := range resM {
			if v > 0 && m[string(r)] > 0 {
				if v > m[string(r)] {
					resM[string(r)] = m[string(r)]
				}
			}else {
				resM[string(r)] = 0
			}
		}
	}

	res := []string{}
	for resk, resv := range resM {
		for i := 0; i < resv; i++ {
			res = append(res,resk)
		}
	}


	return res
}
