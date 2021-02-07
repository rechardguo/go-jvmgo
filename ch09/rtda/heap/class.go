package heap

import (
	"jvmgo/ch09/classfile"
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
	initStarted       bool  //表示类是否已经初始化
	jClass            *Object
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

//func (self *Class) IsSuperClassOf(d *Class) bool {
//
//}

func (self *Class) GetPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

func (self *Class) Name() string {
	return self.name
}

func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}

func (self *Class) SuperClass() *Class {
	return self.superClass
}

func (self *Class) InitStarted() bool {
	return self.initStarted
}
func (self *Class) StartInit() {
	self.initStarted = true
}

func (self *Class) GetClinitMethod() *Method {
	return self.getStaticMethod("<clinit>", "()V")
}

func (self *Class) Loader() *ClassLoader {
	return self.loader
}

func (self *Class) ArrayClass() *Class {
	arrClassName := getArrayClassName(self.name)
	return self.loader.LoadClass(arrClassName)
}

func (self *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := self; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic && field.name == name && field.descriptor == descriptor {
				return field
			}
		}
	}
	return nil
}

func (self *Class) isJlObject() bool {
	return self.name == "java/lang/Object"
}
func (self *Class) isJlCloneable() bool {
	return self.name == "java/lang/Cloneable"
}
func (self *Class) isJioSerializable() bool {
	return self.name == "java/io/Serializable"
}

func (self *Class) JClass() *Object {
	return self.jClass
}

func (self *Class) JavaName() string {
	return strings.Replace(self.name, "/", ".", -1)
}

func (self *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[self.name]
	return ok
}
