package main

type Pig struct {
	name string
	CommonAnimal
}

func (self Pig) Name() string {
	return self.name
}

func (self Pig) Growl() string {
	return "Pig Growl"
}

func newPig(name string) *Pig {
	return &Pig{
		name: name,
	}
}
