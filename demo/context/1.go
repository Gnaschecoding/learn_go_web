package main

type A struct {
}

func (a A) f1()  {}
func (a *A) f2() {}

type B interface {
	f1()
	f2()
}

// 主函数
func main() {

	var b B
	//下面这一行会报错
	b = &A{}
	b.f1()
	b.f2()

	b = &A{}
	b.f1()
	b.f2()

}
