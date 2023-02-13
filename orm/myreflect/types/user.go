package types

import "fmt"

type User struct {
	Name string
	age  int
}

func NewUser(name string, age int) User {
	return User{
		Name: name,
		age:  age,
	}
}

func NewUserPr(name string, age int) *User {
	return &User{
		Name: name,
		age:  age,
	}
}

// 这里用了*User的话， 方法是无法直接被反射的
//
//	func GetAge(u User) int {
//		return u.age
//	}
func (u User) GetAge() int {
	return u.age
}

func (u *User) ChangeName(name string) {
	u.Name = name
}

func (u User) private() {
	fmt.Println("...private")
}
