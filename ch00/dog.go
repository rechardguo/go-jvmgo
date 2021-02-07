package main

type Dog struct {
	name string
	//组合方式，Dog就有了CommonAnimal的Growl,不要写成annimal CommonAnimal
	//不然就变成了成员变量了
	CommonAnimal
}

func NewDog(name string) *Dog {
	return &Dog{
		name: name,
	}
}

// 属于dog特有的行为
func (self *Dog) Name() string {
	return self.name
}

// 属于dog特有的行为
func (self *Dog) Eat() string {
	return "dog eat"
}
