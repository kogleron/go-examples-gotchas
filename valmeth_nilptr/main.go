package main

type Foo interface {
	Print()
}

type FooImpl struct {
}

func (i FooImpl) Print() {
}

func main() {

	var foo Foo
	var fooImplRef *FooImpl

	foo = &FooImpl{}
	foo.Print()

	foo = fooImplRef
	foo.Print()
}
