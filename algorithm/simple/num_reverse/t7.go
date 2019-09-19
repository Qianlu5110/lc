/*
#
# @Time : 2019/9/18 15:00
# @Author : Qian Lu
*/

package num_reverse

/**
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:
输入: 123
输出: 321

示例 2:
输入: -123
输出: -321

示例 3:
输入: 120
输出: 21

注意:
假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231−1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Reverse(x int) int {
	flag := false
	if x < 0 {
		flag = true
	}
	var result int
	if flag {
		res, _ := reverseCalculation(0, -x/1%10, -x/10)
		result = res * -1
	} else {
		res, _ := reverseCalculation(0, x/1%10, x/10)
		result = res
	}

	if result > (1<<31)-1 || result < -(1<<31) {
		return 0
	}
	return result
}

func reverseCalculation(sum int, positionNum int, remainder int) (endx int, nextPositionNum int) {
	if remainder >= 10 {
		sum = sum*10 + positionNum*1
		end, _ := reverseCalculation(sum, remainder/1%10, remainder/10)
		endx = end
	} else if 0 < remainder && remainder < 10 {
		sum = sum*100 + positionNum*10 + remainder
		endx = sum
		return endx, sum
	} else if remainder == 0 {
		sum = positionNum * 1
		endx = sum
		return endx, sum
	}

	return endx, sum
}
