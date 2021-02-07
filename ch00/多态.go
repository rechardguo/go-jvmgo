package main

type Shape struct {
}

func (self *Shape) hasArea(t string) bool {
	return t != "line"
}

type Square struct {
	shape  Shape
	length int
}
