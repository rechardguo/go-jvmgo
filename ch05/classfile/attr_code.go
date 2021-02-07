package classfile

type CodeAtrribute struct {
	cp             ConstantPool
	maxStack       uint16                 //最大的栈深度
	maxLocals      uint16                 //本地变量表,单位是Slot是虚拟机为局部变量分配内存所使用的最小单位
	code           []byte                 //code才是真正用来存放字节码指令的，每一个code占用u1类型，也就是0～255
	exceptionTable []*ExceptionTableEntry //start_pc和end_pc划分了try｛｝，而catch_type代表了catch（exception）里面的那个参数exception，如果抓到异常就转到handler_pc处理
	attributes     []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (self *CodeAtrribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLen := reader.readUint32()
	self.code = reader.readBytes(codeLen)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

func (self *CodeAtrribute) MaxLocals() uint {
	return uint(self.maxLocals)
}

func (self *CodeAtrribute) MaxStack() uint {
	return uint(self.maxStack)
}

func (self *CodeAtrribute) Code() []byte {
	return self.code
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	expceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, expceptionTableLength)
	for i := range exceptionTable {
		entry := ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
		exceptionTable[i] = &entry
	}
	return exceptionTable
}
