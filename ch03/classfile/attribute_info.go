package classfile

//由于attribute也有很多种
//所以先定义一个interface
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

/**
  u2 attribute_count
  attribute_info attribute[attribute_count]
*/
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	//先读取出attribute的个数
	attributeCount := reader.readUint16()
	//建立attributeCount个AttributeInfo的数组
	attributes := make([]AttributeInfo, attributeCount)
	for i := range attributes {
		//依次读取组装到attributes里
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

/**
attribute_info{
  u2 attribute_name_index;
  u4 attribute_length;
  u1 info[attribute_length]
}
*/
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

//使用读出的信息来构建出一个AttributeInfo出来
func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		return &CodeAtrribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprected":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
