package base

import "jvmgo/ch05/rtda"

type Instruction interface {
	FetchOperands(reader *ByteCodeReader)
	Execute(frame *rtda.Frame)
}

//nop 表示什么都不做
type NoOperandsInstruction struct{}

func (self *NoOperandsInstruction) FetchOperands(reader *ByteCodeReader) {
	// nothing to do
}

//BranchInstruction表示跳转指令，Offset字段存放跳转偏移量
type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *ByteCodeReader) {
	self.Offset = int(reader.ReadInt16())
}

//存储和加载类指令需要根据索引存取局部变量表，索引由单字节操作数给出
type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *ByteCodeReader) {
	self.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

//有一些指令需要访问运行时常量池，常量池索引由两字节操作数给出
func (self *Index16Instruction) FetchOperands(reader *ByteCodeReader) {
	self.Index = uint(reader.ReadUint16())
}
