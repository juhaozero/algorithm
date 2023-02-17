package other

import "fmt"

//删除排序数值的重复项
func removeDuplicates(nums []int) int {
	//如果是空切片，那就返回0
	if len(nums) == 0 {
		return 0
	}
	//用两个标记来比较相邻位置的值
	//当一样的话，那就不管继续
	//当不一样的时候，就把right指向的值赋值给left下一位
	left, right := 0, 1
	for ; right < len(nums); right++ {
		if nums[left] == nums[right] {
			continue
		}
		left++
		nums[left] = nums[right]
	}
	fmt.Println(nums[:left+1])
	return left + 1
}

func maxProfit(prices []int) int {
	max := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			max += prices[i+1] - prices[i]
		}
	}
	return max
}
func containsDuplicate(nums []int) bool {
	isok := false
	tmp := make(map[int]int)
	for _, v := range nums {
		_, ok := tmp[v]
		if ok {
			isok = true
			break
		} else {
			tmp[v] = 1
		}
	}
	return isok
}
func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[0]
	}
	return result
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i < 0; i-- {
		if digits[i] != 9 {
			digits[i] = digits[i] + 1
			return digits

		} else {
			digits[i] = 0
		}
	}
	newDigits := make([]int, len(digits)+1)
	newDigits[0] = 1
	newDigits = append(newDigits, digits...)
	return newDigits
}
