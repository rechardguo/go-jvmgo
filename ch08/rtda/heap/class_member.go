package heap

import "jvmgo/ch08/classfile"

/**
字段和方法都属于类的成员，它们有一些相同的信息（访问标
志、名字、描述符）。为了避免重复代码，创建一个结构体存放这些
信息
*/

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (self *ClassMember) CopyMemberInfo(info *classfile.MemberInfo) {
	self.accessFlags = info.AccessFlags()
	self.name = info.Name()
	self.descriptor = info.Descriptor()
}

//判断方法或字段是否能被d访问到
func (self *ClassMember) isAccessibleTo(d *Class) bool {

	if self.IsPublic() {
		return true
	}

	c := self.class
	if self.Isprotected() {
		return d == c || d.IsSubClassOf(c) || c.getPackageName() == d.getPackageName()
	}
	if !self.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}
	return d == c
}

func (self *ClassMember) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}

func (self *ClassMember) Isprotected() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}

func (self *ClassMember) IsPrivate() bool {
	return 0 != self.accessFlags&ACC_PRIVATE
}
