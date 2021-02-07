package heap

type Object struct {
	class  *Class
	fields Slots
}

func (slef *Object) Fields() Slots {
	return slef.fields
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}
