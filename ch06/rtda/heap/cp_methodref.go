package heap

import "jvmgo/ch06/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func (self MethodRef) Descriptor() string {
	return self.descriptor
}

func (self MethodRef) Name() string {
	return self.name
}
func newMethodRef(cp *ConstantPool,
	refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}
