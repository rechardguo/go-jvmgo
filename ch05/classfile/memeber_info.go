package classfile

//jvm对字段和方法表的结构定义
/**
field_info{
u2 access_flags;
u2 name_index;
u2 description_index;
u2 attributes_count;
attribute_info arrtibutes[attributes_count];
}
*/
//对照上面的结构定义出MemberInfo
type MemberInfo struct {
	//保存的是指针，为什么不是*ConstantPool？
	cp               ConstantPool
	accessFlags      uint16
	nameIndex        uint16
	descriptionIndex uint16
	attributes       []AttributeInfo
}

//jvm的 class文件结构定义
/**
u2          fields_count
fieldinfo   field[fields_count]

u2          methods_count
methodinfo  method[methods_count]
*/
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:               cp,
		accessFlags:      reader.readUint16(),
		nameIndex:        reader.readUint16(),
		descriptionIndex: reader.readUint16(),
		attributes:       readAttributes(reader, cp),
	}
}

func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptionIndex)
}

func (self *MemberInfo) CodeAttribute() *CodeAtrribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *CodeAtrribute:
			return attrInfo.(*CodeAtrribute) //感觉像是java的强制类型转换
		}
	}
	return nil
}
