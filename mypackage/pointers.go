package mypackage

import "fmt"

type Person struct {
	Name string
	Doc  int
	Age  int
}

func (p Person) Print() string {
	return fmt.Sprintf("%s - %d - %d", p.Name, p.Doc, p.Age)
}

func (p *Person) AddAge() {
	p.Age++
}
