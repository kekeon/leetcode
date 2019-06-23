package array

func isOneBitCharacter(bits []int) bool {

	l := len(bits)

	label := false

	for i := 0; i < l; {
		if bits[i] == 1 {
			i+=2
			label =  false
		}else if bits[i]==0{
			i+=1
			label =  true
		}
	}

	return label
}
