package native

import "jvmgo/ch09/rtda"

/**
把本地方法定义成一个函数，参数是Frame结构体指针，没有
返回值。这个frame参数就是本地方法的工作空间，也就是连接Java
虚拟机和Java类库的桥梁，
*/
type NativeMethod func(frame *rtda.Frame)

var registry = map[string]NativeMethod{}

func Register(className, methodName, methodDescriptor string, method NativeMethod) {
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}
	// 比如Object里的
	// private static native void registerNatives();
	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}
	return nil
}

func emptyNativeMethod(frame *rtda.Frame) {

}
