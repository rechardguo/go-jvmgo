package heap

import "jvmgo/ch06/classfile"

//相当于Java的继承 Class Field extends ClassMemeber
type Field struct {
	ClassMember
	constValueIndex uint
	slotId          uint
}

func (self Field) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}

func (self Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}

func (self Field) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *Field) IsVolatile() bool {
	return 0 != self.accessFlags&ACC_VOLATILE
}
func (self *Field) IsTransient() bool {
	return 0 != self.accessFlags&ACC_TRANSIENT
}
func (self *Field) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

func (self Field) Descriptor() string {
	return self.descriptor
}

func (self Field) ConstValueIndex() uint {
	return self.constValueIndex
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].CopyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}

func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (self *Field) SlotId() uint {
	return self.slotId
}

func (self *Field) isAccessibleTo(d *Class) bool {
	if self.IsPublic() {
		return true
	}
	c := self.class
	if self.Isprotected() {
		return d == c || d.isSubClassOf(c) || c.getPackageName() == d.getPackageName()
	}
	if !self.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}
	return d == c
}

func (self *Field) Class() *Class {
	return self.class
}
