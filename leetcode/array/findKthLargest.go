package leetcode_array

import (
	"math/rand"
)

// 利用快速排序的思想，从数组 S 中随机找出一个元素 X，把数组分为两部分 Sa 和 Sb。Sa 中的元素大于等于 X，Sb 中元素小于 X。这时有两种情况：
// Sa 中元素的个数小于 k，则 Sb 中的第 k-|Sa| 个元素即为第k大数；
// Sa 中元素的个数大于等于 k，则返回 Sa 中的第 k 大数。时间复杂度近似为 O(n)
func quickSelect(a []int, l, r, index int) int {
	q := randomPartition(a, l, r)
	if q == index {
		return a[q]
	} else if q < index {
		// 右区间查找
		return quickSelect(a, q+1, r, index)
	}
	// 左区间查找
	return quickSelect(a, l, q-1, index)
}

func randomPartition(a []int, l, r int) int {
	// 在l和r之间随机选一个数作为主元下标
	i := rand.Int()%(r-l+1) + l
	// 将pivot与r交换 将主元挪到最后边 避免因为主元的下标对主元左边元素排列的影响
	a[i], a[r] = a[r], a[i]
	return partition(a, l, r)
}

func partition(a []int, l, r int) int {
	x := a[r]
	i := l - 1
	for l < r {
		// 小于最右边的元素的在左边排序
		// i 为排序后的下标
		if a[l] <= x {
			i++
			a[i], a[l] = a[l], a[i]
		}
		l++
	}
	// 将最右边的元素 主元 挪到排序好的位置 满足左边的元素都小于主元 右边的元素都大于主元
	a[i+1], a[r] = a[r], a[i+1]

	return i + 1
}
