package main

import . "fmt"

/**
演示go 语言的错误处理
*/
//类似 java class DividerErrorException extends Exception{...}
type DivideError struct {
	dividee int
	divider int
}

func (self *DivideError) Error() string {
	//这种写法和js里的``很像
	//js里的``可以写上$xx
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
    `
	return Sprintf(strFormat, self.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}

}
