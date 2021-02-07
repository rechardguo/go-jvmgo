package heap

import (
	"fmt"
	"jvmgo/ch06/classfile"
)

type Constant interface {
}

//运行时常量池
type ConstantPool struct {
	class  *Class
	consts []Constant
}

//把class文件中的常量池转换成运行时常量池
func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool {
	consts := make([]Constant, len(cfCp))
	rtCp := &ConstantPool{class, consts}
	for i, cfCpInfo := range cfCp {
		switch cfCpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			//强转
			intInfo := cfCpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Value()
		case *classfile.ConstantFloatInfo:
			floatInfo := cfCpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floatInfo.Value() // float32
		case *classfile.ConstantLongInfo:
			longInfo := cfCpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.Value() // int64
			i++
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cfCpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value() // float64
			i++
		case *classfile.ConstantStringInfo:
			stringInfo := cfCpInfo.(*classfile.ConstantStringInfo)
			consts[i] = stringInfo.String() // string
		case *classfile.ConstantClassInfo:
			classInfo := cfCpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(rtCp, classInfo)
		case *classfile.ConstantFieldrefInfo:
			fieldrefInfo := cfCpInfo.(*classfile.ConstantFieldrefInfo)
			consts[i] = newFieldRef(rtCp, fieldrefInfo)
		case *classfile.ConstantMethodrefInfo:
			methodrefInfo := cfCpInfo.(*classfile.ConstantMethodrefInfo)
			consts[i] = newMethodRef(rtCp, methodrefInfo)
		case *classfile.ConstantInterfaceMethodrefInfo:
			methodrefInfo := cfCpInfo.(*classfile.ConstantInterfaceMethodrefInfo)
			consts[i] = newInterfaceMethodRef(rtCp, methodrefInfo)
		default:
			// todo
		}

	}
	return rtCp
}

func (self *ConstantPool) GetConstant(index uint) Constant {
	if c := self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}
