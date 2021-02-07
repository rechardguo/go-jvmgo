package classfile

/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlag   uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	//为什么这里步定义成[]MemberInfo，[]MemberInfo和[]*MemberInfo有什么区别
	//[]*MemberInfo 是指针数组
	fields     []*MemberInfo
	methods    []*MemberInfo
	attributes []AttributeInfo
}

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}
func (self *ClassFile) MiniorVersion() uint16 {
	return self.minorVersion
}
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" //只有java.lang.Object超类
}
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

//获取所有的interface name
func (self *ClassFile) InterfaceNames() []string {
	//interfaces 存的是常量池里的index
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlag
}

//对class的文件结构的解析
//这里如果写成*ClassFile,相当于self就是this
func (self *ClassFile) read(reader *ClassReader) {
	//为啥步是reader.readAndCheckMagic(self)
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlag = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.classFormartError:magic!")
	}
}

func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError")
}

func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

//解析字节数组得到ClassFile结构体
func Parse(classDate []byte) (cf *ClassFile, err error) {
	//panic-recover机制,用来做异常catch处理,如果这个处理了，就不会详细打印了
	//defer func() {
	//	if r:= recover();r!=nil{
	//		var ok bool
	//		err,ok=r.(error)
	//
	//		if !ok{
	//			fmt.Errorf("%v",r)
	//		}
	//	}
	//}()

	cr := &ClassReader{classDate}
	cf = &ClassFile{}
	cf.read(cr)
	return
}
