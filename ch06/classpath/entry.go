package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

//对于传入的path,类型是string,进行解析
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.Contains(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") ||
		strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") ||
		strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
