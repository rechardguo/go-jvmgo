package heap

import (
	"fmt"
	"jvmgo/ch09/classfile"
	"jvmgo/ch09/classpath"
)

/**
这里的代码这是演示了
记载-链接-初始化
*/
type ClassLoader struct {
	cp          *classpath.ClassPath
	verboseFlag bool
	classMap    map[string]*Class
}

func NewClassLoader(cp *classpath.ClassPath, verboseFlag bool) *ClassLoader {
	loader := &ClassLoader{
		cp:          cp,
		verboseFlag: verboseFlag,
		classMap:    make(map[string]*Class),
	}
	loader.loadBasicClasses()
	loader.loadPrimitiveClasses()
	return loader
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
	if self.verboseFlag {
		fmt.Printf("[loaded %s from %s]", name, entry)
	}
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
	var class *Class
	if name[0] == '[' {
		class = self.loadArrayClass(name)
	} else {
		class = self.loadNonArrayClass(name)
	}

	if jlClassClass, ok := self.classMap["java/lang/Class"]; ok {
		class.jClass = jlClassClass.NewObject()
		class.jClass.extra = class
	}
	return class
}

//数组的class是构造出来的
//非数组的class是从class file读出来的
func (self *ClassLoader) loadArrayClass(name string) *Class {
	class := &Class{
		accessFlags: ACC_PUBLIC,
		name:        name,
		loader:      self,
		initStarted: true,
		superClass:  self.loadClass("java/lang/Object"),
		interfaces: []*Class{
			self.loadClass("java/lang/Cloneable"),
			self.loadClass("java/io/Serializable"),
		},
	}
	self.classMap[name] = class
	return class
}

func (self *ClassLoader) loadBasicClasses() {
	jlClassClass := self.LoadClass("java/lang/Class")
	for _, class := range self.classMap {
		if class.jClass == nil {
			class.jClass = jlClassClass.NewObject()
			class.jClass.extra = class
		}
	}
}

// primitiveType是void、int、float等
func (self *ClassLoader) loadPrimitiveClasses() {
	for primitiveType, _ := range primitiveTypes {
		self.loadPrimitiveClass(primitiveType)
	}
}

func (self *ClassLoader) loadPrimitiveClass(className string) {
	class := &Class{
		accessFlags: ACC_PUBLIC,
		name:        className,
		loader:      self,
		initStarted: true,
	}
	class.jClass = self.classMap["java/lang/Class"].NewObject()
	class.jClass.extra = class
	self.classMap[className] = class
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
