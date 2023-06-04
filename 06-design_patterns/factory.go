package main

import "fmt"

type Service interface {
	sum(a int, b int) int
}

type AService struct {
	c int
}

type BService struct {
	c int
}

func (s *AService) sum(a int, b int) int {
	return a
}

func (s *BService) sum(a int, b int) int {
	return b
}

func main() {
	var a *Service
	var b *Service
	a = &AService{c: 1}
	b = &BService{c: 2}
	fmt.Println(a.sum(1, 2))
	fmt.Println(b.sum(1, 2))
}
