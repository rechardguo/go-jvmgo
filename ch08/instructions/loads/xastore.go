package loads

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type AASTORE struct{ base.NoOperandsInstruction }
type BASTORE struct{ base.NoOperandsInstruction }
type CASTORE struct{ base.NoOperandsInstruction }
type DASTORE struct{ base.NoOperandsInstruction }
type FASTORE struct{ base.NoOperandsInstruction }
type IASTORE struct{ base.NoOperandsInstruction }
type LASTORE struct{ base.NoOperandsInstruction }
type SASTORE struct{ base.NoOperandsInstruction }

/**
<t>astore系列指令按索引给数组元素赋值


iastore指令的三个操作数分别是：要赋给数组元素的值、数组索引、数组引用，依次从操作数栈中弹出。
如果数组引用是null，则抛出NullPointerException。
如果数组索引小于0或者大于等于数组长度，则抛出ArrayIndexOutOfBoundsException异常。

如果一切正常，则按索引给数组元素赋值
*/

func (self *AASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopRef()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	if val.Class() != arrRef.Class() {
		panic("java.lang.ArrayStoreException")
	}

	checkNotNil(arrRef)
	refs := arrRef.Refs()
	checkIndex(len(refs), index)
	refs[index] = val
}

func (self *IASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	ints := arrRef.Ints()
	checkIndex(len(ints), index)
	ints[index] = int32(val)
}
