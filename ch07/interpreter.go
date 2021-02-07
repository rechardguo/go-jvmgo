package main

import (
	"fmt"
	"jvmgo/ch07/instructions"
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
	"jvmgo/ch07/rtda/heap"
)

//解释器
func interpret(method *heap.Method, logInst bool /*methodInfo *classfile.MemberInfo*/) {
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
	//defer catchErr(thread)
	loop(thread, logInst)
}

//计算pc、解码指令、执行指令
func loop(thread *rtda.Thread, logInst bool) {
	reader := &base.BytecodeReader{}
	thread.CurrentFrame()
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)
		//decode
		//pc相当于bytecode[]的index, 每个指令占一个字节
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadInt8() // opcode e.g 就是i_const0
		inst := instructions.NewInstruction(byte(opcode))
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		if logInst {
			logInstruction(frame, inst)
		}
		//excute
		//fmt.Printf("pc:%2d inst:%T %v \n",pc,inst,inst)
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}

}

func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)

}

func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		//fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		//fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		logFrames(thread)
		panic(r)
	}
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n",
			frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}
