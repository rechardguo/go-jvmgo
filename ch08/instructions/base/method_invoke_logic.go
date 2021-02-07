package base

import (
	"fmt"
	"jvmgo/ch08/rtda"
	"jvmgo/ch08/rtda/heap"
)

/**
  invokerFrame:调用者的栈
*/
func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	//创建新的帧
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)
	argSlotSlot := int(method.ArgSlotCount())
	if argSlotSlot > 0 {
		for i := argSlotSlot - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}

	//hack!
	//而Java类库中的很多类都要注册本地方法，比如Object类就有一个registerNatives（）
	//这里先跳过
	if method.IsNative() {
		if method.Name() == "registerNatives" {
			thread.PopFrame()
		} else {
			panic(fmt.Sprintf("native method: %v.%v%v\n",
				method.Class().Name(), method.Name(), method.Descriptor()))
		}
	}

}
