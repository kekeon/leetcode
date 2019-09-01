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

	vs := []int{}
	obj.Put(1, 1)
	obj.Put(2, 2)
	v := obj.Get(1)
	vs = append(vs, v)
	v = obj.Get(3)
	vs = append(vs, v)
	obj.Put(2, 1)
	v = obj.Get(2)
	vs = append(vs, v)
	obj.Remove(2)
	v = obj.Get(2)
	vs = append(vs, v)

	fmt.Println(vs)
}

