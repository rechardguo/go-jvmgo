package classpath

import (
	"os"
	"path/filepath"
)

type ClassPath struct {
	bootClassPath Entry
	extClassPath  Entry
	userClassPath Entry
}

func (self *ClassPath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClassPath = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClassPath = newWildcardEntry(jreExtPath)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		//将路径连在一起
		return filepath.Join(jh, "jre")
	}
	panic(" Can not find jre path")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// -cpa/b;a.jar;b
func (self *ClassPath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClassPath = newEntry(cpOption)
}

//-Xjre解析启动类和扩展类
func Parse(jreOption string, cpOption string) *ClassPath {
	cp := &ClassPath{}
	//parseBootAndExtClasspath里的 (self *ClassPath),其中的self就是指向了cp
	//由于cp是指针，所以类型是*ClassPath
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

//读取class文件,双亲加载原则
func (self *ClassPath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"

	data, entry, err := self.bootClassPath.readClass(className)
	if err == nil {
		return data, entry, err
	}
	//上面的写法也可以写成下面的，使用;连在一起更加的紧凑
	if data, entry, err := self.extClassPath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClassPath.readClass(className)
}

func (self ClassPath) String() string {
	return self.userClassPath.String()
}
