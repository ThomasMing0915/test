package main

type sliceStack struct {
	arr  []int //用切片模拟栈
	size int
}

func (s *sliceStack) IsEmpty() bool {
	return s.size == 0
}

func (s *sliceStack) Size() int {
	return s.size
}

func (s *sliceStack) Push(v int) {
	s.arr = append(s.arr, v)
	s.size += 1
}

func (s *sliceStack) Pop() int {
	if s.IsEmpty() {
		panic("slice stack is empty, can't pop")
	}
	topValue := s.arr[s.size-1]
	s.arr = s.arr[:s.size]
	s.size -= 1

	return topValue
}

func (s *sliceStack) Top() int {
	if s.IsEmpty() {
		panic("slice stack is empty, can't top")
	}

	topValue := s.arr[s.size-1]
	return topValue
}
