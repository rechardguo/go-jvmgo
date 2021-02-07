## 如何在idea里建立一个go的工程

以d:/dev_code为例子

1. 在d:/dev_code/go目录下建立src文件夹
2. d:/dev_code/go/src 建立工程,例如jvmgo
3. idea 导入  d:/dev_code/go/src/jvmgo
4. file -> setting 找到go 配置 goroot 和 gopath,在gopath的project path里加入d:/dev_code/go,
   如果在电脑的配置里配了gopath,在global path会显示出来，如果不需要可以在可以勾掉第一个选项，一般不需要

- goroot：类似jdk
- gopath：就是项目路径，不要使用公共的项目路径,**必须在gopath里建立一个src目录**



run debug configration里的字段不是很清楚是干什么用的,默认吧


## 各章节内容
- ch00 主要是各种实验，直接入手了golang,虽然有java的基础，C也学过一些，但是有些语法还是很陌生
  该章节主要是用来实验了。包括了如下内如：
  
  - 1.字符串的操作，strings工具类的使用
  ```go
    //字符串得替换,最后得-1表是替换所有，相当于java里得replaceAll
    strings.Replace(cmd.class,".","/",-1)
    //没有equals,相等得判断使用==
    m.Name()=="main"
   
    strings.Contains
  
  //字符串不能类似java的 x!=null
    var x string = nil //error
    
    if x == nil { //error
          x = "default"
     }
  
     if x == "" {
            x = "default"
        }
   ```
  - 2.数组的主要操作，例如数组的截取操作等
  ```go
    //创建数组1
    arr := []int{1, 4, 5, 2, 3}
    //创建数组2
    arr2:=make([]int,5)
    //遍历数组
    for _,v:=range arr{
     println(v)
    }
    //截取,2包含5不包含
    arr[2:5] 	
  ```
  - 3.interface的继承,如何实现继承?
      go的继承不需要显式的实现接口，而是方法和接口的一致即可
     
     interface如何判断是那个实现类呢？xxx.(type)
     
     参考 constant_pool.go 
     类似Java的 instanceOf 和类型强转 
  
  - 4.go语言没有继承的概念，但可以通过 结构体嵌套来模拟 
    ```go
      //定义了Animal的行为
      type Animal interface{
         Growl() string
         Beat() string 
      }  
      //定义了公共的Animal的行为
      type CommonAnimal struct{
      }
      func Growl(self *CommonAnimal)string{
         return "every animal can growl"
      } 
      func Beat(self *CommonAnimal)string{
         return "every animal can beat"
      } 
      //Dog,没有类似java的extends或implements
      type Dog struct{
         //这样就拥有了 CommonAnimal的行为，不用重复定义了
         CommonAnimal
      }
      //Pig
      type Pig struct{
        //这样就拥有了 CommonAnimal的行为，不用重复定义了
        CommonAnimal
        //定义属于Pig的属性
        name string
      }
      //覆盖了CommonAnimal的方法，类似 java 的@Override
      func (self *Pig)Growl()string{ 
        return "pig is growl..."
      }
    ```      
  - 5.数据转换
           
     |go|java|
     |---|---|
     |整数类型|
     ||uint(uint32)|
     |byte|int8|
     |short|int16|
     |int(int32)|int32|
     |long|int64|
     |char|uint16|
     |浮点数字类型|
     |float|float32|
     |double|float64| 
     |布尔类型|
     |boolean|bool|
     |引用类型|
     |类类型|*Object|
     
     uint32 转成为float32类型
     math.Float32frombits()    
     uint16 不能直接当成int来使用，可以使用int(xxx)来转成int
  - 6.panic-recover机制？  
    go得错误处理,recover()函数得处理
  - 7.struct 结构体
    ```go
     //直接xxx{}相当于是new出一个对象
    type User struct{...}    
    u:=User{}
    u2=&User{} //取得指针 
    ```  
  
  - 8.for 循环语法
  
    ```go
      for i:=1;i<cpCount;i++{
      ...
      }
      //或者使用range
     ``` 
   - 9.`.(type)` 必须在switch的语句里使用
   Use of .(type) outside type switch
   
   - 10.方法命名的问题，为什么有的大写有的小写?        
   > golang中根据首字母的大小写来确定可以访问的权限 
   > - 函数大写字母开头，其他包可以访问该函数
   > - 函数小写字母开头，同包可以访问，其他包不能访问 
   
   - 11.int int32 int64    
   > int数字， 就代表了几位
     go语言中的int的大小是和操作系统位数相关的，如果是32位操作系统，int类型的大小就是4字节。如果是64位操作系统，int类型的大小就是8个字节。
   
  - 12.怎么创建出int int32 int64 float32 float64 等等... 
   ```go
      var i int=10;
      j:=20
   ```
  - 13.map的使用
   >map类似Java里的HashMap
   ```go
     //key是int ,value是string 
     map[int]string{1:"rechard",2:"tom",3:"james"}
   ```  
  - 14.Go语言的包不能相互依赖
  
  - 15.数组
  ```go
  //建立数组方式1
  arr:=[]int{1,4,5,2,3} 
  //建立数组方式2
  arr2: make([]int,5)
  //取得数组的长度
  len(arr)
  //遍历数组1
  for k,v:=range arr{
  ...
  }
  //遍历数组2
  for i:=0;i<len(arr);i++{...}
  //数组的截取
  arr=arr[0:1]
  
  ```
  - 16.当一个变量有各种类型，而又没有一个公共的接口时候，可以定义一个空的interface
  ```go
    type Constant interface {
    }

  ```
  Go语言的interface{}类型很像C语言中的void*，该类型的变量可以容纳任何类型的值。
  
  
   - 17. go的递归应用学习，参考class_hierarchy.go
   
   - 18. go没有 while(true) ，用的是for
   ```go
    for{
     //某个条件满足  
     break
    }
   ```
   - 19 string 也可以当成数组,比如
   
   ```go
    str:="123";
    str[0] //取得第一个
   ```
   - 20 import for side effect
   
   > 如果没有任何包依赖lang包，它就不会被编译进可执行文件，
    上面的本地方法也就不会被注册。所以需要一个地方导入lang包，
    把它放在invokenative.go文件中。由于没有显示使用lang中的变量或
    函数，所以必须在包名前面加上下划线，否则无法通过编译
   
   ```go
 import _ "jvmgo/ch09/native/java/lang"
   ```
    
   - 21 import 
   下面这个fmt是Go语言的标准库，他其实是去GOROOT下去加载该模块
   ```go
     import(
         "fmt"
       )     
   ``` 
    
   Go的import还支持如下两种方式来加载自己写的模块   
   ```go
     //当前文件同一目录的model目录，但是不建议这种方式import
     import   "./model"
      // 绝对路径
     import   "shorturl/model"  //加载GOPATH/src/shorturl/model模块 
   ```  
    
   包导入的过程
   > 程序的初始化和执行都起始于main包。如果main包还导入了其它的包，那么就会在编译时将它们依次导入。有时一个包会被多个包同时导入，那么它只会被导入一次（例如很多包可能都会用到fmt包，但它只会被导入一次，因为没有必要导入多次）。当一个包被导入时，如果该包还导入了其它的包，那么会先将其它包导入进来，然后再对这些包中的包级常量和变量进行初始化，接着执行init函数（如果有的话），依次类推。等所有被导入的包都加载完毕了，就会开始对main包中的包级常量和变量进行初始化，然后执行main包中的init函数（如果存在的话），最后执行main函数
    
    
  
- ch01 java Hello.class 的执行过程，涉及了命令的解析

- ch02 对于一些classpath的解析，java的class文件查找机制
  
  ```go
   //组合模式的使用
  
  ```
- ch03 讲述对class文件结构的解析，主要是熟悉了class 文件结构
  todo:这里的解析过于麻烦了，等后期改成使用访问者模式来改造
  
- ch04 java的线程运行时数据区里栈，主要介绍了栈里的frame operandStack(操作数栈)
       和localvariableTable(本地变量表)

- ch05 Java虚拟机解释器工作：计算pc、指令解码、指令执行

  ```go
   do {
      atomically calculate pc and fetch opcode at pc;
      if (operands) fetch operands;
      execute the action for the opcode;
   } while (there is more to do);
 
   ```
   指令的操作，其实就是操作的操作数栈和本地变量表的之间的配合
   
   
   ```go
     //抽象一个指令的接口方法
     type Instruction interface {
     	FetchOperands(reader *BytecodeReader)
     	Execute(frame *rtda.Frame)
     }
   ```
   不同的具体指令类需要实现这个接口，并实现具体的逻辑
   
   
   
   
- ch06
   实现了方法区、运行时常量池、类和对象结构体、一个简单的类加载器，以及ldc和部分引用类指令
   基本的思路在
   classLoader.LoadClass
      
      - 加载：通过一个classloader操作: User.class->classfile->成为class    
      - 链接：verify,prepare(初始化并给static 变量赋值)  
      
   
- ch07 方法调用
 方法调用指令需要n+1个操作数，其中第1个操作数是uint16索引，在字节码中紧跟在指令操作码的后面 
 
 什么是动态绑定？假设
 ```java
 public class Animal{
  public void growl(){...} 
}
  
 public class Dog extends Animal{
   public void growl(){...}
 }

 public class Pig extends Animal{
   public void growl(){...}
 }


 Animal dog=new Dog();
 ``` 
  动态绑定就是在编译时候是不知道dog是哪个具体类,它的符号引用还是Animal, 调用的方法依赖于隐式参数的实际类型
 
  静态绑定（static binding）：也叫前期绑定，在程序执行前，该方法就能够确定所在的类，此时由编译器或其它连接程序实现。
 
  接口方法符号和非接口方法符号?
  
  类D想通过方法符号引用访问类C的某个方法，先要解析符号引用得到类C
  如果C是接口，则是接口方法调用，如果C是类，则是非接口方法调用
  
  - invokestatic :指令用来调用静态方法。
  - invokespecial:指令用来调用无须动态绑定的实例方法，包括构造函数、私有方法和通过super关键字调用的超类方法
  - invokevirtual:用于调用对象的实例方法
  - invokeinterface:调用接口方法
  - invokedynamic:动态语言
  
  方法调用和参数传递:
  在定位到需要调用的方法之后，Java虚拟机要给这个方法创建一个新的帧并把它推入Java虚拟机栈顶，然后传递参数
 
 - ch08 数组和字符串的解析 
  
    - newarray指令用来创建数组
    需要两个操作数, 第一个操作数是一个uint8整数，在字节码中紧跟在指令操作码后面，表示要创建哪种类型的数组
    - anewarray指令用来创建引用类型数组
    - arraylength指令用于获取数组长度
    
    数组相关的指令<t>aload <t>astore <t> 可以是a,b,c,d,f,i,l,s
    
    - <t>aload 指令按索引取数组元素值,类似getfield
    - <t>astore 指令按索引给数组元素赋值
    - multianewarray 指令 
 
  数组的类名 [
   - int[]的类名是[I
   - int[][]的类名是[[I
   - Object[]的类名是 [Ljava/lang/Object；
   - String[][]的类名是[[java/lang/String；
 
 string 字符串的解析
 ---
 **第8章运行出现问题**
 
 
array  
objectref,helloword null


size =2-1-1

报nil错误

0 getstatic #2 <java/lang/System.out>
3 ldc #3 <hello world>
5 invokevirtual #4 <java/io/PrintStream.println>
8 return

invokevirtual 在执行的时候，在操作数栈里已经准备好了
invokevirtual所需要的数据
..., objectref, [arg1, [arg2 ...]] →
...

下面这句就是从操作数栈里取得objectref
ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() -1) 报nil
这句为空，也就是说

getstatic #2 <java/lang/System.out> 没有将数据放到栈里

getstatic 它取出类的某个静态变量值，然后推入栈顶

..., →

..., value

*references.GET_STATIC &{{2}}

取得System类的out静态变量

发现System类 staticVars都为空 说明初始化staticVars时候有问题

发现class_loader.go在initStaticFinalVar的时候，只是对基础的几个类型和String赋值，其他的类型没赋值

```go
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex() // //这里一直都是0，应该是有问题
	slotId := field.SlotId()

	if cpIndex > 0 { //为什么是大于0,0为什么不行
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			goStr := cp.GetConstant(cpIndex).(string)
			jStr := JString(class.Loader(), goStr)
			vars.SetRef(slotId, jStr)
		}
	}
}
```


 
 
## go 如何 debug
configuration里 Run kind选Deirectory,不要选file

## 注意点
1. java的写法 e.g.: 是byte[]  而 go是 []byte
2. panic-recover

## 错误
- 报undefined错误
有多个go文件在一起的时候，选file运行会报undefined错误

- 运行的时候要有main.go文件(入口文件),main.go里必须有 package main

- 同一个目录下可以不用import,不同的目录下的go文件要使用的话就需要import进来使用


## 参考

jvm的指令集合
https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-6.html#jvms-6.5# go-jvmgo
