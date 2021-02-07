package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

//根据后缀名选出jar文件
func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] //去掉末尾的*
	compositeEntry := []Entry{}   //定义一个Entry的数组
	walkFN := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		//目录匹配，并且后缀名是.jar或.JAR则加入数组
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	//遍历baseDir目录，walkFN是一个回调
	filepath.Walk(baseDir, walkFN)
	return compositeEntry
}
