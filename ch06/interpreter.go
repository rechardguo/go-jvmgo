package main

import (
	"fmt"
	"jvmgo/ch06/instructions"
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
	"jvmgo/ch06/rtda/heap"
)

//解释器
func interpret(method *heap.Method /*methodInfo *classfile.MemberInfo*/) {
	//codeAttr:=methodInfo.CodeAttribute()
	//maxLocals := codeAttr.MaxLocals()
	//maxStack := codeAttr.MaxStack()
	//bytecode := codeAttr.Code()
	//thread:=rtda.NewThread()
	//frame:=rtda.NewFrame(thread,maxLocals,maxStack)
	//thread.PushFrame(frame)
	//defer catchErr(frame)
	//loop(thread,bytecode)
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, method.Code())
}

//计算pc、解码指令、执行指令
func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		//一开始的pc是0
		pc := frame.NextPC()
		thread.SetPC(pc)
		//decode
		//pc相当于bytecode[]的index, 每个指令占一个字节
		reader.Reset(bytecode, pc)
		opcode := reader.ReadInt8() // opcode e.g 就是i_const0
		inst := instructions.NewInstruction(byte(opcode))
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		//excute
		fmt.Printf("pc:%2d inst:%T %v \n", pc, inst, inst)
		inst.Execute(frame)
	}

}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}
