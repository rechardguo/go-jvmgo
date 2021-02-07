package heap

import (
	"fmt"
	"jvmgo/ch06/classfile"
	"jvmgo/ch06/classpath"
)

/**
这里的代码这是演示了
记载-链接-初始化
*/
type ClassLoader struct {
	cp       *classpath.ClassPath
	classMap map[string]*Class
}

func NewClassLoader(cp *classpath.ClassPath) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}

func (self *ClassLoader) loadClass(name string) *Class {
	if class, exist := self.classMap[name]; exist {
		return class
	}
	return self.loadNonArrayClass(name)
}

//首先找到class文件并把数据读取到内存；
//然后解析class文件，生成虚拟机可以使用的类数据，并放入方法区；
//最后进行链接
func (self *ClassLoader) loadNonArrayClass(name string) *Class {

	//加载并解析成class文件
	data, entry := self.readClass(name)
	class := self.defineClass(data)
	//链接
	link(class)
	fmt.Printf("[loaded %s from %s]", name, entry)
	return class
}

func link(class *Class) {
	verify(class)
	prepare(class)
}

//准备阶段主要是给类变量分配空间并给予初始值
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = NewSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}

}

//初始化static
//如果静态变量属于基本类型或String类型，有final修饰符，
//且它的值在编译期已知，则该值存储在class文件常量池中。
//initStaticFinalVar（）函数从常量池中加载常量值，然后给静态变量
//赋值
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()
	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			goStr := cp.GetConstant(cpIndex).(string)
			jStr := JString(class.Loader(), goStr)
			vars.SetRef(slotId, jStr)
		}
	}

}

func calcStaticFieldSlotIds(class *Class) {
	slotIdCount := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotIdCount
			slotIdCount++
			if field.isLongOrDouble() {
				slotIdCount++
			}
		}
	}
	class.staticSlotCount = slotIdCount
}

//计算实例字段的个数，同时编号
func calcInstanceFieldSlotIds(class *Class) {
	slotIdCount := uint(0)
	if class.superClass != nil {
		slotIdCount = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotIdCount
			slotIdCount++
			if field.isLongOrDouble() {
				slotIdCount++
			}
		}
	}
	class.instanceSlotCount = slotIdCount
}

func verify(class *Class) {
	//todo
}

func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException:" + name)
	}
	return data, entry
}

func (self *ClassLoader) defineClass(data []byte) *Class {
	//讲class文件数据转成Class结构体
	class := parseClass(data)
	//对于每个class对象来说都有一个classloader
	class.loader = self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}

func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		// already loaded
		return class
	}

	return self.loadNonArrayClass(name)
}

func resolveInterfaces(class *Class) {
	interfacesCount := len(class.interfaceNames)
	if interfacesCount > 0 {
		class.interfaces = make([]*Class, interfacesCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.loadClass(interfaceName)
		}
	}
}

//加载类时候，发现有父类也会进行加载
//加载父类也是用当前类的加载器来加载的
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.loadClass(class.superClassName)
	}
}

func parseClass(data []byte) *Class {
	//解析出classfile
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	//通过classfile来得到一个Class对象
	return newClass(cf)
}
