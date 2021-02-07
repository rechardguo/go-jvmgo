package classfile

// 主要是deprecatedAttribute 和 syntheticAttribute 2个的定义

type DeprecatedAttribute struct {
	MarkerAttribute
}

// Synthetic合成的
type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct {
}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
}
