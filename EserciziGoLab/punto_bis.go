package main

import "fmt"

type Punto struct {
	x, y float64
}

func main() {
	var x, y float64
	fmt.Scan(&x, &y)
	p1 := newPunto(x, y)
	fmt.Print(p1)
	specchiaPunto(&p1)
	fmt.Println(p1)
}

func newPunto(x, y float64) (p Punto) {
	p.x, p.y = x, y
	return
}

func specchiaPunto(p *Punto) {
	(*p).x = -p.x
}
