package buildmodel

import "fmt"

type Person struct {
	//person detail
	name, address, pin string
	//job detail
	workAddress, company, position string
	salary                         int
}
type PersonBuilder struct {
	person *Person
}
type PersonJobBuilder struct {
	PersonBuilder
}
type PersonAddressBuilder struct {
	PersonBuilder
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{person: &Person{}}
}
func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}
func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}
func (a *PersonAddressBuilder) At(addr string) *PersonAddressBuilder {
	a.person.address = addr
	return a
}
func (a *PersonAddressBuilder) WithPostalCode(pin string) *PersonAddressBuilder {
	a.person.pin = pin
	return a
}
func (j *PersonJobBuilder) As(position string) *PersonJobBuilder {
	j.person.position = position
	return j
}
func (j *PersonJobBuilder) In(addr string) *PersonJobBuilder {
	j.person.workAddress = addr
	return j
}
func (j *PersonJobBuilder) For(company string) *PersonJobBuilder {
	j.person.company = company
	return j
}
func (j *PersonJobBuilder) WithSalary(salary int) *PersonJobBuilder {
	j.person.salary = salary
	return j
}
func (b *PersonBuilder) Build() *Person {
	return b.person
}

func OtherRun() {
	pb := NewPersonBuilder()
	pb.Lives().At("kok").WithPostalCode("sdfsd").Works().As("sdf").For("ibm").WithSalary(232423)
	person := pb.Build()
	fmt.Println(person)
}
