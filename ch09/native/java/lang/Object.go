package lang

import (
	. "jvmgo/ch09/native"
	"jvmgo/ch09/rtda"
)

func init() {
	Register("java/lang/Object", "getClass", "()Ljava/lang/Class;", getClass)
}

func getClass(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}
