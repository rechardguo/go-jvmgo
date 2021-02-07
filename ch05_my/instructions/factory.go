package instructions

import (
	"fmt"
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/instructions/constants"
)

//单例变量
//有很大一部分指令是没有操作数的，没有必要每次都创建不同的实例?
var (
	nop         = &constants.NOP{}
	aconst_null = &ACONST_NULL{}
)

//根据opcode来得到具体的指令类
func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return nop
	case 0x01:
		return aconst_null

	default:
		//%x byte
		panic(fmt.Errorf("Unsupported opcode ：0x%x!", opcode))
	}
}
