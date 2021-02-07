package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	//panic("implement me")
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}

//返回的是结构体指针
//类似class newDirEntry{},本身加上new是构造函数
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		//?
		panic(err)
	}
	return &DirEntry{absDir}
}
