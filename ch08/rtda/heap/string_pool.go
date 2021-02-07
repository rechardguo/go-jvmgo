package heap

import "unicode/utf16"

/**
java code:

String s1="hello world";

存在常量池里

const pool:

const #1=class  #2;
const #2=Asciz    xx/xxx/xx/TestClass

*/

var internedStrings = map[string]*Object{}

func JString(loader *ClassLoader, goStr string) *Object {
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}
	chars := stringToUtf16(goStr)
	jChars := &Object{loader.LoadClass("[C"), chars}
	jStr := loader.LoadClass("java/lang/String").NewObject()
	jStr.SetRefVar("value", "[C", jChars)
	internedStrings[goStr] = jStr
	return jStr
}

// java.lang.String -> go string
func GoString(jStr *Object) string {
	charArr := jStr.GetRefVar("value", "[C")
	return utf16ToString(charArr.Chars())
}

//Go语言字符串在内存中是UTF8编码的，先把它强制转成
//UTF32，然后调用utf16包的Encode（）函数编码成UTF16
func stringToUtf16(s string) interface{} {
	runes := []rune(s)
	return utf16.Encode(runes)
}

func utf16ToString(s []uint16) string {
	runes := utf16.Decode(s) // utf8
	return string(runes)
}
