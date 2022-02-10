package main

import (
	"fmt"
)

type Shape interface{
	Area() float64
	Parimeter() float64
}

type React struct{
	width float64
	height float64
}

func(r React) Area() float64{
	return r.width*r.height
}

func (r React) Parimeter() float64{
	return 2*(r.width+r.height)
}

func main(){
	var s Shape
	s = React{5.0, 4.0}
	r := React{5.0,4.0}
	fmt.Printf("type of s is %T\n",s)
	fmt.Printf("value of s is %v\n",s)
	fmt.Println("area of rectang s", s.Area())
	fmt.Println("s==r is",s == r)
}