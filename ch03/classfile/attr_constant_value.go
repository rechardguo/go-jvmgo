package classfile

type ConstantValueAttribute struct {
	constantValueIndex uint16
}

//属于类的方法是小写
func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

//属性之类的是大写
func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
