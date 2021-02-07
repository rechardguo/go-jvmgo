package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type ISHL struct{ base.NoOperandsInstruction }  // int左位移
type ISHR struct{ base.NoOperandsInstruction }  // int算术右位移
type IUSHR struct{ base.NoOperandsInstruction } // int逻辑右位移
type LSHL struct{ base.NoOperandsInstruction }  // long左位移
type LSHR struct{ base.NoOperandsInstruction }  // long算术右位移
type LUSHR struct{ base.NoOperandsInstruction } // long逻辑右位移

func (self *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

//逻辑右移指令
func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}
