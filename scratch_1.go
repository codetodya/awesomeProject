package main

import "fmt"

type A struct {
	Name string
}

func main() {
	a1 := A{}
	a1.Name = "I am a1"

	a2 := A{}
	a2.Name = "I am a2"

	a3 := &A{}
	a3.Name = "I am a3"

	var as []A

	as = append(as, a1)
	as = append(as, a2)
	as = append(as, *a3)

	fmt.Printf("length of as: %d\n", len(as))
	fmt.Printf("as[0].Name: %s\n", as[0].Name)
	fmt.Printf("as[1].Name: %s\n", as[1].Name)
	fmt.Printf("as[2].Name: %s\n", as[2].Name)

	letChangeTheNames(as)

	fmt.Printf("length of as: %d\n", len(as))
	fmt.Printf("as[0].Name: %s\n", as[0].Name)
	fmt.Printf("as[1].Name: %s\n", as[1].Name)
	fmt.Printf("as[2].Name: %s\n", as[2].Name)

	areTheNamesReallyChanging(as)

	fmt.Printf("length of as: %d\n", len(as))
	fmt.Printf("as[0].Name: %s\n", as[0].Name)
	fmt.Printf("as[1].Name: %s\n", as[1].Name)
	fmt.Printf("as[2].Name: %s\n", as[2].Name)
}

func letChangeTheNames(bs []A) {
	for i, a := range bs {
		a.Name = fmt.Sprintf("a%d got a new Name, Hurray!!", i+1)
	}
}


func areTheNamesReallyChanging(bs []A) {
	for i, _ := range bs {
		bs[i].Name = fmt.Sprintf("a%d getting a new Name?", i+1)
	}
}
