package main

type Foo struct{}

func (f Foo) Print() {
}

func main() {
	var foo *Foo

	foo.Print()
}
