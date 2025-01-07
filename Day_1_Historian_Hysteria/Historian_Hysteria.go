package main

import (
	"fmt"
	"sort"
)

type Vertex struct {
	X int
	Y int
}

func main() {

	//Compile our list
	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}
	fmt.Println("Before Sort")
	fmt.Println(list1)
	fmt.Println(list2)

	//Reorganize both arrays into a ascending order then compare
	sort.Ints(list1)
	sort.Ints(list2)
	fmt.Println("After Sort")
	fmt.Println(list1)
	fmt.Println(list2)

	//compare our lists and then pair the smallest numbers
	a := Vertex{0, 0}
	b := Vertex{0, 0}
	c := Vertex{0, 0}
	d := Vertex{0, 0}
	e := Vertex{0, 0}
	f := Vertex{0, 0}

	fmt.Println("Vertexes")
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)
	// fmt.Println(f)
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e, f)
	// a.X = list1[0]
	// a.Y = list2[0]
	// b.X = list1[1]
	// b.Y = list2[1]
	// c.X = list1[2]
	// c.Y = list2[2]
	// d.X = list1[3]
	// d.Y = list2[3]
	// e.X = list1[4]
	// e.Y = list2[4]
	// f.X = list1[5]
	// f.Y = list2[5]
	a.X, a.Y = list1[0], list2[0]
	b.X, b.Y = list1[1], list2[1]
	c.X, c.Y = list1[2], list2[2]
	d.X, d.Y = list1[3], list2[3]
	e.X, e.Y = list1[4], list2[4]
	f.X, f.Y = list1[5], list2[5]
	fmt.Printf("Filled \n")
	fmt.Printf("(%v, %v) \n", a.X, a.Y)
	fmt.Printf("(%v, %v) \n", b.X, b.Y)
	fmt.Printf("(%v, %v) \n", c.X, c.Y)
	fmt.Printf("(%v, %v) \n", d.X, d.Y)
	fmt.Printf("(%v, %v) \n", e.X, e.Y)
	fmt.Printf("(%v, %v) \n", f.X, f.Y)
	//finding the distance between our pairs using vertex
	dista := a.X - a.Y
	distb := b.X - b.Y
	distc := c.X - c.Y
	distd := d.X - d.Y
	diste := e.X - e.Y
	distf := f.X - f.Y
	if dista < 0 {
		dista = dista * (-1)
	}
	if distb < 0 {
		distb = distb * (-1)
	}
	if distc < 0 {
		distc = distc * (-1)
	}
	if distd < 0 {
		distd = distd * (-1)
	}
	if diste < 0 {
		diste = diste * (-1)
	}
	if distf < 0 {
		distf = distf * (-1)
	}
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", dista, distb, distc, distd, diste, distf)
	totaldist := 0
}
