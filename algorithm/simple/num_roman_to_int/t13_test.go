/*
#
# @Time : 2019/9/20 11:15
# @Author : Qian Lu
*/

package num_roman_to_int

import (
	"fmt"
	"testing"
)

func Test_RomanToInt(t *testing.T) {

	// 3
	t1 := RomanToInt("III")
	fmt.Println(t1)

	// 4
	t2 := RomanToInt("IV")
	fmt.Println(t2)

	// 9
	t3 := RomanToInt("IX")
	fmt.Println(t3)

	// 58
	t4 := RomanToInt("LVIII")
	fmt.Println(t4)

	// 1994
	t5 := RomanToInt("MCMXCIV")
	fmt.Println(t5)

	// 650
	t6 := RomanToInt("DCXXI")
	fmt.Println(t6)

}

func Test_RomanSubToInt(t *testing.T) {
	v1 := RomanSubStrToInt("IV")
	fmt.Println(v1)

	v2 := RomanSubStrToInt("III")
	fmt.Println(v2)

	v3 := RomanSubStrToInt("XXI")
	fmt.Println(v3)
}
