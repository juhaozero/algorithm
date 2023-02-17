package sort

import "fmt"

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {

		var flag = true
		for j := 0; j < len(arr)-1-i; j++ {

			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				// 在里面我们将其设置为 false
				flag = false
			}
		}
		if flag {
			break
		}
	}
}
func selectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		// 未排序区间最小值初始化为第一个元素
		min := i
		// 从未排序区间第二个元素开始遍历，直到找到最小值
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}

	}
}
func insertionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		// 每次从未排序区间取一个数据 value
		value := arr[i]
		// 在已排序区间找到插入位置
		j := i - 1
		for ; j >= 0; j-- {
			// 如果比 value 大后移
			if arr[j] > value {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		// 插入数据 value
		arr[j+1] = value
	}
	return arr
}

// 快速排序
func qsort(data []int64, low, high int) {
	if low < high {
		fmt.Println(data)
		m := partition(data, low, high)
		qsort(data, low, m-1)
		qsort(data, m+1, high)
	}
}

func partition(data []int64, low, high int) int {
	key := data[low]
	tmpLow := low
	tmpHigh := high
	for {
		//查找大于等于key的元素
		for {
			if data[tmpHigh] > key {
				break
			}

			if tmpHigh <= tmpLow {
				break
			}
			tmpHigh--
		}
		//找到小于key的元素，该元素的位置一定是low到tmpHigh+1之间。因为array[tmpHigh+1]必定大于key
		for {
			if data[tmpLow] < key {
				break
			}

			if tmpLow >= tmpHigh {
				break
			}
			tmpLow++
		}
		if tmpLow >= tmpHigh {
			break
		}

		data[tmpLow], data[tmpHigh] = data[tmpHigh], data[tmpLow]

	}
	data[tmpLow], data[low] = data[low], data[tmpLow]
	return tmpLow
}

// Binary 二分
func Binary(arr []int, num, low, high int) int {
	// 递归终止条件
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	// 递归查找
	if num > arr[mid] && mid == len(arr)-1 {
		return mid + 1
	} else if num > arr[mid] {
		return Binary(arr, num, mid+1, high)
	} else if num < arr[mid] && num > arr[mid-1] {
		return mid
	} else if num < arr[mid] {
		return Binary(arr, num, low, mid-1)
	} else if num == arr[mid] {
		return mid + 1
	} else {
		return mid
	}
}
