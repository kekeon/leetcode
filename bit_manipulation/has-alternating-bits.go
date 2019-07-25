package bit_manipulation

func hasAlternatingBits(n int) bool {
   n = n^n>>1
   return n&(n+1) == 0
}