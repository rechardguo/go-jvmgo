package rtda

type Stack struct {
	maxSize uint
	size    uint
	//_私有属性？
	_top *Frame
}

func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.stackOverflowError")
	}
	if self._top != nil {
		frame.lower = self._top
	}
	self._top = frame
	self.size++

}

func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size--
	return top
}

func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty")
	}
	return self._top
}
