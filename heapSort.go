package main

import (
	"fmt"
	"errors"
)

type heap []int

func (h *heap)Push(i int) {
	*h = append(*h, i)
	h.up()
}

func (h *heap)Pop() (int, error) {
	lens := len(*h)
	if lens == 0 {
		return 0, errors.New("heap is empty")
	}
	max := (*h)[0]
	h.down()
	return max, nil
}

func (h heap)up() {
	lastIdx := len(h) - 1
	for lastIdx > 0 {
		parentIdx := (lastIdx - 1) / 2
		if h[lastIdx] <= h[parentIdx] {
			return
		}
		h[lastIdx], h[parentIdx] = h[parentIdx], h[lastIdx]
		lastIdx = parentIdx
	}
}

func (h *heap)down() {
	// i 2*i + 1 2*i + 2
	H := *h
	lens := len(H)
	if lens <= 1 {
		return
	}
	H[0], H[lens - 1] = H[lens - 1], H[0]
	i, n := 0, lens - 1
	for {
		maxIdx := 2*i + 1
		if maxIdx >= n || maxIdx < 0 { // maxIdx < 0 after int overflow
			break
		}
		if j := maxIdx + 1; j < n && H[j] > H[maxIdx] {
			maxIdx = j
		}
		if H[i] > H[maxIdx] {
			break
		}
		H[i], H[maxIdx] = H[maxIdx], H[i]
		i = maxIdx
	}
	*h = H[:n]
}

func main() {
	eles := []int{78, 7, 4, 34, 67, 12, 89, 10, 7, 3}
	h := new(heap)
	for _, ele := range eles {
		h.Push(ele)
	}

	fmt.Println(h)

	res := []int{}
	for i := 0; i < 5; i++ {
		temp, _ := h.Pop()
		res = append(res, temp)
	}
	fmt.Println(h)
	fmt.Println(res)
}



