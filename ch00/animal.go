package main

//定义了Animal的行为的接口
type Animal interface {
	Beat() string  //撕咬
	Growl() string //咆哮
}

//Animal的公共实现类
type CommonAnimal struct {
}

func (self *CommonAnimal) Growl() string {
	return "every animal can growl"
}

//这个是CommonAnimal的方法
func (self *CommonAnimal) Beat() string {
	return "every animal can beat"
}

//如何实现继承，go的继承不需要显式的实现接口，而是方法一致即可
func newAnmial(name string) Animal {
	if name == "pig" {
		return newPig(name)
	} else if name == "dog" {
		return NewDog(name)
	}
	return nil
}
