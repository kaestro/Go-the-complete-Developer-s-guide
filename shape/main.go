package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct { line float64 }
type square struct { line float64 }


func main() {

}

func (t triangle) getArea() float64 {
	return t.line * t.line / 2
}

func (s square) getArea() float64 {
	return s.line * s.line
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}