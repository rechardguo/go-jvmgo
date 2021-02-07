package classfile

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}
type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	//先读取出LineNumberTableEntry的个数
	entryLen := reader.readUint16()
	//建立entryLen个LineNumberTableEntry的数组
	entrys := make([]*LineNumberTableEntry, entryLen)
	for i := range entrys {
		entrys[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
	self.lineNumberTable = entrys
}
