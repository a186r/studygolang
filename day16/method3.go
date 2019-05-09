// 展示内嵌结构体上的方法可以直接在外层类型的实例上调用
package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
	Point
	name string
}

func main() {
	n := &NamedPoint{Point{3, 4}, "Alina"}
	fmt.Println(n.Abs())
}

// 内嵌将一个已存在类型的字段和方法注入到了另一个类型里，：匿名字段上的方法晋升为了外层类型的方法。
// 当然类型可以有只作用于内嵌父类型上的方法

// 可以覆写方法
func (n *NamedPoint) Abs() float64 {
	return n.Point.Abs() * 100
}
