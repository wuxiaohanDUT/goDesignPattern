package singleton

//饿汉模式
type person struct {
}

var instance *person

func init() {
	instance = &person{}
}

func GetInstance() *person{
	return instance
}
