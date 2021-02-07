package classfile

/*
字段符号的引用
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
普通（非接口）方法符号的引用
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
接口方法符号引用
CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberrefInfo) name() string {
	return self.cp.getClassName(self.classIndex)
}

func (self *ConstantMemberrefInfo) nameAndTypeDescription() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

//go语言没有继承的概念，但可以通过 结构体嵌套来模拟

type ConstantFieldRefInfo struct {
	ConstantMemberrefInfo
}
type ConstantMethodRefInfo struct {
	ConstantMemberrefInfo
}
type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}
