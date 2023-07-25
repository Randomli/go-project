package leetcode_array

import (
	"logger"
)

var LOGFILE string = "./leetcode_array.log"
var LOGERRFILE string = "./leetcode_array.err.log"

// 1. 将0元素移动到列表末尾
func moveZeroes(nums []int) {

	index_list_0 := []int{}
	for k, v := range nums {
		if v == 0 {
			index_list_0 = append(index_list_0, k)
		}
	}
	for index, v := range index_list_0 {
		// 每次循环都会删掉了一个元素 所以元素下标减去已经删除的元素个数
		v = v - index
		nums = append(nums[:v], nums[v+1:]...)
		// 列表末尾补0
		nums = append(nums, 0)
	}
}

// 1. 答案
func moveZeroesSolution(nums []int) {

	n := len(nums)
	i := -1
	// nums[0....i]表示非0元素的数列,初始值i=-1
	for j := 0; j <= n-1; j++ {
		if nums[j] != 0 {
			i++
			nums[i] = nums[j]
		}
	}
	for k := i + 1; k <= n-1; k++ {
		nums[k] = 0
	}
}

// 2. 给你一个数组nums和一个值val，你需要原地移除所有数值等于val的元素，并返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须仅使用O(1)额外空间并原地修改输入数组。
// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
func removeElement(nums []int, val int) int {
	n := len(nums)
	val_index_list := []int{}
	for index, v := range nums {
		if v == val {
			val_index_list = append(val_index_list, index)
		}
	}
	for index, v := range val_index_list {
		v = v - index
		nums = append(nums[:v], nums[v+1:]...)
		nums = append(nums, 0)
	}
	return n - len(val_index_list)
}

// 3. 找出字符串中第一个匹配项的下标
func strStr(haystack string, needle string) int {
	needleNums := len(needle)
	for index, i := range haystack {
		if i == rune(needle[0]) {
			if index+needleNums <= len(haystack) {
				if haystack[index:index+needleNums] == needle {
					return index
				}
			}
		}
	}
	return -1
}

// 4. 给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
// 考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
// 更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
// 返回 k 。
func removeDuplicates(nums []int) int {
	n := len(nums)
	newNumsN := n

	for i := 0; i <= n-1; i++ {
		// 比较前后两个元素的大小
		if i+1 <= len(nums)-1 && nums[i] == nums[i+1] {
			newNumsN--
			// 删掉相同的元素
			nums = append(nums[:i], nums[i+1:]...)
			i = i - 1
		}
	}

	return newNumsN
}

// 4 题解
func removeDuplicatesSolution(nums []int) int {
	n := len(nums)
	if n == 0 || n == 1 {
		return n
	}
	//  nums[0,i]为非重复数列
	l := 0
	r := l + 1
	for r <= n-1 {
		// 右边的指针走的快，左边的指针走的慢
		// 当左右指针的值不相等时，左指针的值加1，然后将右指针的值赋值给左指针
		// [0,0,1,1]
		// l: 0 r: 2
		// nums[l]: 0 nums[r]: 1
		// nums[l+1] = nums[r]
		// [0,1,1,1]
		// l: 1 r: 3
		if nums[r] != nums[l] {
			// 指向同一个元素不需要赋值
			if l+1 != r {
				nums[l+1] = nums[r]
			}
			l++
		}
		r++
	}
	return l + 1
}

// 5. 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
func removeDuplicates2(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	l := 0
	r := l + 1
	for r <= n-1 {
		// 三指针
		if nums[r] != nums[l] {
			if l+1 != r {
				nums[l+1] = nums[r]
			}
			l++
		}
		r++
	}
	return l + 1
}

func Run() {
	logger := logger.InitLogger(LOGFILE, LOGERRFILE)
	defer logger.Sync()
	logger.Info("in leetcode")
	// 1
	// nums := []int{0, 0, 0, 30, 0, 12, 0, 0, 0, 0}
	// moveZeroes(nums)
	// moveZeroesSolution(nums)
	// logger.Info(nums)

	// 2
	// nums := []int{0, 1, 1, 1, 2, 0}
	// val := 1
	// n := removeElement(nums, val)
	// logger.Info(nums)
	// logger.Info(n)

	// 3
	// n := strStr("a", "a")
	// logger.Info(n)

	// 4
	// nums := []int{0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 4}
	// // n := removeDuplicates(nums)
	// n := removeDuplicatesSolution(nums)
	// logger.Info(nums)
	// logger.Info(n)

	// 5
	nums := []int{0, 0, 1, 1, 1, 1, 1, 2, 3, 3}
	n := removeDuplicates2(nums)
	logger.Info(nums)
	logger.Info(n)
}
