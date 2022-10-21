package online_stock

import "math"

type StockSpanner struct {
	stack [][2]int
	idx   int
}

func Constructor() StockSpanner {
	return StockSpanner{[][2]int{{-1, math.MaxInt32}}, -1}
}

func (s *StockSpanner) Next(price int) int {
	s.idx++
	for price >= s.stack[len(s.stack)-1][1] {
		s.stack = s.stack[:len(s.stack)-1]
	}
	s.stack = append(s.stack, [2]int{s.idx, price})
	return s.idx - s.stack[len(s.stack)-2][0]
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
