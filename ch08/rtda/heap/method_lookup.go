package heap

//先攻当前类开始找,如果找不到就去接口里找
//类只能继承一个
func LookupMethodInClass(class *Class, name string, descriptor string) *Method {
	for c := class; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}
	return nil
}

//接口可以多继承，所以这里从接口方法里找到方法的话，需要实现递归查找
func lookupMethodInInterfaces(interfaces []*Class, name string, descriptor string) *Method {
	/*for _,c:=range interfaces{
		if m:=LookupMethodInClass(c,name,descriptor);m!=nil{
			return m
		}
	}*/

	for _, iface := range interfaces {
		for _, method := range iface.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
		method := lookupMethodInInterfaces(iface.interfaces, name, descriptor)
		if method != nil {
			return method
		}
	}

	return nil
}
