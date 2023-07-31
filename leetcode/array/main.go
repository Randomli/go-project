package leetcode_array

import (
	"fmt"
	"math/rand"
	"time"

	"henry.com/go-project/logger"
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

// 5. 给定有序数组 保留k个相同元素的通用解法
func removeDuplicates2(nums []int) int {
	var process func(k int) int
	process = func(k int) int {
		// 记录切片的下标
		i := 0
		for _, v := range nums {
			// 因为是有序的切片 将当前位置的前k个元素与当前元素比较
			// 不同的 赋值给当前元素 i自增, 相同的 i不自增,循环下一个元素
			// 如果i<k或者nums[i-k] != v，就将v赋值给nums[i]
			if i < k || nums[i-k] != v {
				nums[i] = v
				i++
			}
		}
		return i
	}
	return process(2)
}

// 6. 给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums
// 原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
// 我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
// 必须在不使用库内置的 sort 函数的情况下解决这个问题。
func sortColors(nums []int) {
	n := len(nums)
	i := 0
	l := 0
	r := n - 1
	for i <= r {
		fmt.Println(i, l, r, nums)
		switch nums[i] {
		case 0:
			nums[l], nums[i] = nums[i], nums[l]
			l++
			i++
		case 2:
			nums[r], nums[i] = nums[i], nums[r]
			r--
		case 1:
			i++
		}
	}
}

// 7. 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
// 快速排序的思想
// 随机选取一个元素作为主元 对主元前面的元素排序
// 如果主元的下标q刚好等于k，那么主元就是第k大的元素
// 否则，如果 q 比目标下标小，就递归右子区间，否则递归左子区间。
// 这样就可以把原来递归两个区间变成只递归一个区间，提高了时间效率。
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
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
	// 	nums := []int{0, 0, 1, 1, 1, 1, 1, 2, 3, 3}
	// 	n := removeDuplicates2(nums)
	// 	logger.Info(nums)
	// 	logger.Info(n)

	// 6
	// nums := []int{2, 0, 1}
	// sortColors(nums)
	// logger.Info(nums)

	// 7
	nums := []int{2, 4, 1, 3, 9, 9, 8}
	// 解法一：快速排序 效率更高
	// k := findKthLargest(nums, 2)
	// logger.Info(k)
	// logger.Info(nums)

	// 解法二：堆排序
	kk := findKthLargestByHeap(nums, 2)
	logger.Info(kk)
	logger.Info(nums)
}
