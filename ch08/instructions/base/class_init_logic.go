package base

import (
	"jvmgo/ch08/rtda"
	"jvmgo/ch08/rtda/heap"
)

func InitClass(thread *rtda.Thread, class *heap.Class) {
	//把类的initStarted状态置成true以免进入死循环
	class.StartInit()
	//初始化,这样的话如果多个线程同时到这里怎么办？
	scheduleClinit(thread, class)
	initSuperClass(thread, class)
}

func initSuperClass(thread *rtda.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && !superClass.InitStarted() {
			InitClass(thread, superClass)
		}
	}
}

//
func scheduleClinit(thread *rtda.Thread, class *heap.Class) {
	clinit := class.GetClinitMethod()
	if clinit != nil {
		//初始化方法new出一个栈帧
		newFrame := thread.NewFrame(clinit)
		//将栈帧压入到栈里
		thread.PushFrame(newFrame)
	}
}
