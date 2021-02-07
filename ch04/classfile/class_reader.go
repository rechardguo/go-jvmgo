package classfile

import "encoding/binary"

//就是byte[]类型的包装
type ClassReader struct {
	data []byte
}

//u1 无符号1个字节
//读取出一个字节
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

//u2 无符号2个字节，一个字word
func (self *ClassReader) readUint16() uint16 {
	/*val:=self.data[0:2]
	self.data=self.data[2:]
	return val*/
	//大端模式
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

//u4
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

//?u8
func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

//读取uint16表，表的大小由开头的uint16数据指出
func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	s := make([]uint16, n)
	// s是个数组
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

//读出即个字节，length uint32
//uint32 等价于 java 的int类型
func (self *ClassReader) readBytes(n uint32) []byte {
	//读取从0到n-1子字段
	//这里没做长度判断？ 如果超了呢？
	//超了就会抛出 panic错误
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes

}
