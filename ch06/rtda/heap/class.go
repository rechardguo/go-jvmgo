package heap

import (
	"jvmgo/ch06/classfile"
	"strings"
)

//和ClassFile的区别
//ClassFile里存的还是比较原始的，而这里则是从原始的材料里进一步提炼出来
type Class struct {
	accessFlags       uint16
	name              string //thisClassName
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	superClass        *Class
	interfaces        []*Class
	loader            *ClassLoader
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots //静态变量是和class相关的
}

//ClassFile结构体转换成Class结构体
func newClass(cf *classfile.ClassFile) *Class {
	cl := &Class{}
	cl.accessFlags = cf.AccessFlags()
	cl.name = cf.ClassName()
	cl.superClassName = cf.SuperClassName()
	cl.interfaceNames = cf.InterfaceNames()
	cl.constantPool = newConstantPool(cl, cf.ConstantPool())
	cl.fields = newFields(cl, cf.Fields())
	cl.methods = newMethods(cl, cf.Methods())
	return cl
}

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}

func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) getStaticMethod(name string, descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() &&
			method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return nil
}

func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() || self.getPackageName() == other.getPackageName()
}

func (self *Class) getPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""

}

func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}

//判断类是否是一个接口类型
func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}

func (self *Class) StaticVars() Slots {
	return self.staticVars
}

func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}

func (self *Class) Loader() *ClassLoader {
	return self.loader
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: NewSlots(class.instanceSlotCount),
	}
}
