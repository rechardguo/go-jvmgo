package classpath

import (
	"errors"
	"strings"
)

//type name name的类型
type CompositeEntry []Entry

//  (self CompositeEntry) 指明了readClass是CompositeEntry 方法
func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, from, err := entry.readClass(className)
		if err == nil { //表示正确
			return data, from, err
		}
	}
	return nil, nil, errors.New("class not found :" + className)
}

func (self CompositeEntry) String() string {
	//相当于new string[n]
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}

func newCompositeEntry(pathList string) CompositeEntry {
	//定义了一个Entry的数组
	//java: Entry[] compositeEntry=new Entry[];
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		//var entry=newEntry(path)
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}
