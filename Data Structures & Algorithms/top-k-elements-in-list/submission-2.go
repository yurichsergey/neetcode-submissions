
type FreqHeap [][2]int // [number, frequency]

func (h FreqHeap) Len() int           { return len(h) }
func (h FreqHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h FreqHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *FreqHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *FreqHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
    freq := map[int]int{}
	for _, i := range nums {
		freq[i] += 1
	}

	h := &FreqHeap{}
	heap.Init(h)

	for num, freq := range freq {
		heap.Push(h, [2]int{num, freq})
		if len(*h) > k {
			heap.Pop(h) // Remove the smallest frequency
		}
	}

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(h).([2]int)[0]
	}
	return res
}
