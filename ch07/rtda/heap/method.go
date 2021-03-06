package heap

import "jvmgo/ch07/classfile"

type Method struct {
	ClassMember  //相当于继承了ClassMember
	maxStack     uint
	maxLocals    uint
	code         []byte
	argSlotCount uint
}

func (self *Method) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}

func (self *Method) MaxLocals() uint {
	return self.maxLocals
}

func (self *Method) MaxStack() uint {
	return self.maxStack
}

func (self *Method) Code() []byte {
	return self.code
}

func (self *Method) Class() *Class {
	return self.class
}

func (self *Method) Name() string {
	return self.name
}

func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
}

//从方面的描述符里可以确定方法需要几个Slot
func (self *Method) calcArgSlotCount() {
	parsedDescriptor := parseMethodDescriptior(self.descriptor)

	for _, paramType := range parsedDescriptor.parameterTypes {
		self.argSlotCount++
		//J float D dobule  要占2个slot
		if paramType == "J" || paramType == "D" {
			self.argSlotCount++
		}
	}
	//非static方法，第一个变量是this
	if !self.IsStatic() {
		self.argSlotCount++
	}

}

func (self *Method) IsProtected() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}

func (self *Method) Descriptor() string {
	return self.descriptor
}

func (self *Method) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}

func (self *Method) IsNative() bool {
	return 0 != self.accessFlags&ACC_NATIVE
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].CopyMemberInfo(cfMethod)
		copyAttributes(cfMethod, methods[i])
		methods[i].calcArgSlotCount()
	}
	return methods
}

func copyAttributes(cfMethod *classfile.MemberInfo, method *Method) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		method.maxStack = codeAttr.MaxStack()
		method.maxLocals = codeAttr.MaxLocals()
		method.code = codeAttr.Code()
	}
}
