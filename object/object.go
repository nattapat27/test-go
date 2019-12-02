package object

type person struct {
	Name string
	Age int
}

func NewWithValue(name string, age int) *person{
	obj := person{Name:name, Age:age}
	return &obj
}
