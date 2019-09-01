package hash_table

import (
	"testing"
	"fmt"
)

func TestLongestPalindrome(t *testing.T) {
	s := "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"
	v := longestPalindrome(s)
	fmt.Println(v)
}

func TestFindWords(t *testing.T)  {
	s := []string{"Hello", "Alaska", "Dad", "Peace"}
	v := findWords(s)
	fmt.Println(v)
}

func TestFindLHS(t *testing.T) {
	s := []int{1,1,1,1,1}
	v := findLHS(s)
	fmt.Println(v)
}

func TestMyHashMap(t *testing.T) {
	obj := ConstructorMap()

	obj.Put(100, 10)
	v := obj.Get(100)

	fmt.Println(v)
}

