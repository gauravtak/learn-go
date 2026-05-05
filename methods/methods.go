package main


type rect struct {
	width int
	height int
}

type circle struct {
	radius float32
}

func (r *rect) area() int {
	return r.width * r.height 
}

func (c *circle) area() float32 {
	return 3.14 * float32(c.radius) * float32(c.radius)
}

func main() {
	r := rect{width: 10, height: 20}
	rp := &r
	rp.area()

	c := circle{radius: 7.5}
	cp := &c
	cp.area()
}
