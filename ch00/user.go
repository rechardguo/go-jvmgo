package main

import (
	"strings"
)

type User struct {
	name string
	age  int
}

func newUser(name string, age int) *User {
	return &User{name, age}
}
func (u *User) String() string {
	//int 类型不能直接加
	//rechard=#
	//return u.name+"="+string(u.age)

	//info:=[]string{u.name,string(u.age)}
	//return strings.Join(info,"=");

	info := make([]string, 2)
	info[0] = u.name
	info[1] = string(u.age)
	//info=append(info,u.name)
	//info=append(info,string(u.age))
	//panic("panic") //相当于是抛出异常
	return strings.Join(info, "=")
}

func (self *User) setName(name string) {
	self.name = name
}
