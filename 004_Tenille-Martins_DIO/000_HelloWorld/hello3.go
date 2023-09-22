package main

import "fmt"

func main() {
	fmt.Println("🤣🙄")
	fmt.Println(`
	Teste 123 
	      789
	      456
	    Total
	Soma tudo. 
	`)

	a := 44
	b := 43.0
	fmt.Printf("%v, %T \n", a, a)
	fmt.Printf("%v, %T \n", b, b)

	alex := 44
	fmt.Printf("Alex tem %b \n", alex)
	fmt.Printf("Alex tem %x \n", alex)
	fmt.Printf("Alex tem %X \n", alex)

	aa, bb, c, d, e, f := 20, 30, 40, 50, 60, 70
	fmt.Printf("%d \t %#b \t %x \n", aa, aa, aa)
	fmt.Printf("%d \t %#b \t %x \n", bb, bb, bb)
	fmt.Printf("%d \t %#b \t %x \n", c, c, c)
	fmt.Printf("%d \t %#b \t %x \n", d, d, d)
	fmt.Printf("%d \t %#b \t %x \n", e, e, e)
	fmt.Printf("%d \t %#b \t %x \n", f, f, f)
}
