package main

import (
	"container/heap"
	"fmt"
	"math"
)

type GiftsHeap []int64

func (gh GiftsHeap) Len() int           { return len(gh) }
func (gh GiftsHeap) Less(i, j int) bool { return gh[i] > gh[j] }
func (gh GiftsHeap) Swap(i, j int)      { gh[i], gh[j] = gh[j], gh[i] }

func (gh *GiftsHeap) Push(x interface{}) {
	*gh = append(*gh, x.(int64))
}

func (gh *GiftsHeap) Pop() interface{} {
	old := *gh
	n := len(old)
	x := old[n-1]
	*gh = old[:n-1]
	return x
}

func pickGifts(gifts []int, k int) int64 {
	gh := &GiftsHeap{}
	heap.Init(gh)
	var sum int64
	for _, gift := range gifts {
		heap.Push(gh, int64(gift))
		sum += int64(gift)
	}

	for i := 1; i <= k; i++ {
		gift := heap.Pop(gh).(int64)
		sqrtGift := int64(math.Floor(math.Sqrt(float64(gift))))
		sum -= (gift - sqrtGift)
		heap.Push(gh, sqrtGift)
	}

	return sum
}

type Input struct {
	gifts []int
	k     int
}

func main() {
	inputs := []Input{
		{
			gifts: []int{25, 64, 9, 4, 100},
			k:     4,
		},
		{
			gifts: []int{1, 1, 1, 1},
			k:     4,
		},
	}
	for _, input := range inputs {
		fmt.Println(pickGifts(input.gifts, input.k))
	}
}
