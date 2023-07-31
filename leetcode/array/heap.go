package leetcode_array

type MaxHeap struct {
	arr []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		arr: make([]int, 0),
	}
}

func (h *MaxHeap) insert(val int) {
	// 在堆中，通常用数组来表示，其中数组索引和二叉树节点之间存在一种特殊的关系。
	// 对于给定节点的索引i，其左子节点的索引为2i+1，右子节点的索引为2i+2，而其父节点的索引为(i-1)/2。
	h.arr = append(h.arr, val)
	i := len(h.arr) - 1

	// 从底部向上调整堆
	for i > 0 {
		// 父节点
		parent := (i - 1) / 2
		// 当前节点小于父节点，调整结束
		if h.arr[i] <= h.arr[parent] {
			break
		}
		// 当前节点大于父节点，交换位置
		h.arr[i], h.arr[parent] = h.arr[parent], h.arr[i]
		// 节点上移
		i = parent
	}
}

func (h *MaxHeap) extractMax() int {
	// 排除空堆的情况
	n := len(h.arr)
	if n == 0 {
		panic("Heap is empty")
	}

	// 大顶堆的堆顶元素为最大值
	max := h.arr[0]
	// 将最后一个元素放到堆顶
	h.arr[0] = h.arr[n-1]
	// 删除最后一个元素 max
	h.arr = h.arr[:n-1]

	// 从顶部向下调整堆 维持堆的特性
	// 从堆顶开始，如果当前节点小于子节点，则交换位置，直到当前节点大于子节点或者没有子节点
	i := 0
	for {
		largest := i
		left := 2*i + 1
		right := 2*i + 2

		if left < len(h.arr) && h.arr[left] > h.arr[largest] {
			largest = left
		}

		if right < len(h.arr) && h.arr[right] > h.arr[largest] {
			largest = right
		}

		if largest == i {
			break
		}

		h.arr[i], h.arr[largest] = h.arr[largest], h.arr[i]
		i = largest
	}

	return max
}

// 在findKthLargestByHeap函数中，首先获取输入整数切片nums的长度，然后调用buildMaxHeap函数来构建大顶堆。
// 接着，通过循环从堆中提取k次最大元素，每次提取后调整堆。
// 最后，返回堆顶元素（最大元素），即第k个最大元素的值。
func findKthLargestByHeap(nums []int, k int) int {
	max := 0
	heap := NewMaxHeap()

	for _, v := range nums {
		heap.insert(v)
	}

	for i := 0; i < k; i++ {
		max = heap.extractMax()
	}
	return max
}
