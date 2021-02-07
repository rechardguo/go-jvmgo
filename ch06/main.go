package main

import (
	"fmt"
	"jvmgo/ch06/classfile"
	"jvmgo/ch06/classpath"
	"jvmgo/ch06/rtda/heap"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Printf("version 1.0")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.xjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp)
	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	//method:=getMainMethod(classfile.Methods())
	method := mainClass.GetMainMethod()
	if method != nil {
		interpret(method)
	} else {
		panic("main method not found")
	}
}

func getMainMethod(methods []*classfile.MemberInfo) *classfile.MemberInfo {
	for _, m := range methods {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}

/*func loadClass(className string, cp *classpath.ClassPath) *classfile.ClassFile {
	classData,_,err:=cp.ReadClass(className)
	if err!=nil{
		panic(err)
	}
	cf,err:=classfile.Parse(classData)
	if err!=nil{
		panic(err)
	}
	return cf
}*/
func printClassInfo(cf *classfile.ClassFile) {
	//%v	相应值的默认格式。
	//%t	true 或 false  布尔占位符
	//整数占位符
	/**
	%b	二进制表示	Printf("%b", 5)	101
	%c	相应Unicode码点所表示的字符	Printf("%c", 0x4E2D)	中
	%d	十进制表示	Printf("%d", 0x12)	18
	%o	八进制表示	Printf("%d", 10)	12
	%q	单引号围绕的字符字面值，由Go语法安全地转义	Printf("%q", 0x4E2D)	'中'
	%x	十六进制表示，字母形式为小写 a-f	Printf("%x", 13)	d
	%X	十六进制表示，字母形式为大写 A-F	Printf("%x", 13)	D
	%U	Unicode格式：U+1234，等同于 "U+%04X"	Printf("%U", 0x4E2D)	U+4E2D
	*/
	//%s	输出字符串表示（string类型或[]byte)
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MiniorVersion())
	fmt.Printf("constant counts: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("   %s\n", f.Name())
	}
	fmt.Printf("methodss count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("   %s\n", m.Name())
	}

}
