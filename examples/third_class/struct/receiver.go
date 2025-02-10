package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// 结构体接收器，不会影响自身的状态
func (u User) ChangeName(newName string) {
	u.Name = newName
}

// 指针接收器，会影响自身的状态
func (u *User) ChangeAge(newAge int) {
	u.Age = newAge
}

func main() {
	u := User{
		Name: "小明",
		Age:  20,
	}
	u.ChangeName("xiaoming")
	u.ChangeAge(18)
	fmt.Printf("%#v\n", u)

	up := &User{
		Name: "小赵",
		Age:  20,
	}
	up.ChangeAge(18)
	up.ChangeName("xiaozhao")
	fmt.Printf("%#v\n", up)
}
