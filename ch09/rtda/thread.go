package rtda

import "jvmgo/ch09/rtda/heap"

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func newStack(stackLen int) *Stack {
	return &Stack{
		maxSize: 1000,
	}
}

//返回的是pc指针
func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return NewFrame(self, method)
}

func (self *Thread) TopFrame() *Frame {
	return self.CurrentFrame()
}

func (self *Thread) IsStackEmpty() bool {
	return self.stack.IsEmpty()
}
