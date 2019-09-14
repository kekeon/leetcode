package bitManipulation

// 用 num & (num-1) 可以判断条件1，比如：100(4) & 011(3) == 0，结果为 0 说明符合条件1；
// 是否在奇数位可以用 0xaaaaaaaa 判断，16 进制的 a 是 1010，比如：0100(4) & 1010(a) == 0，结果为 0 说明最高位 1 在奇数位上；

func isPowerOfFour(num int) bool {
	return num>0 && num&(num-1)==0 && 0xaaaaaaaa&num==0
}