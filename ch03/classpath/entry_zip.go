package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	//绝对路径
	absDir string
}

//相当于构造函数,为啥返回的是*ZipEntry，而不是Entry
func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absDir)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()

	for _, f := range r.File {
		if f.Name == className {
			//io.ReadCloser
			//Open returns a ReadCloser that provides access to the File's contents.
			//想当于返回了java中的流对象
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			//延迟关闭？
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)

}

func (self *ZipEntry) String() string {
	return self.absDir
}

//func newZipEntry(path string) Entry {
//}
