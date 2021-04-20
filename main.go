package main

import "fmt"

type Person struct {
	name   string
	age    int
	status bool
}

func change(old *int, new int) {
	*old = new
	fmt.Println("melakukan eksekusi fungsi")
	fmt.Println(*old)
}

func main() {
	number := 30

	num2 := &number

	num3 := &num2

	num5 := number

	fmt.Println("sebelum eksekusi fungsi", number, *num2, **num3, num5)

	change(&number, 40)

	fmt.Println("setelah eksekusi fungsi", number, *num2, **num3, num5)

	// p1 := Person{"afista pratama", 23, true}

	// p2 := &p1
	// p3 := &p1

	// *p2 = Person{"afista", 50, false}

	// *p3 = Person{"wuku wuku", 30, false}

	// fmt.Println(p1)
	// fmt.Println(p2)
	// fmt.Println(p3)

	// pointer
	// sederhana
	// age := 23

	// a1 := &age
	// a2 := &age

	// *a1 = 30
	// *a2 = 50

	// a1 , a2 => reference dari si var age (&)

	// a1 = 30
	// a2 = 35

	// fmt.Println(age)
	// fmt.Println(*a1)
	// fmt.Println(*a2)
	// fmt.Println(a2)
}
