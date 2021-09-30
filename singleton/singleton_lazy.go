package singleton

import "sync"

//懒汉模式
type lazyPerson struct {
}

var (
	lazyInstance *lazyPerson
	once 		 = &sync.Once{}
)

func GetLazyInstance() *lazyPerson{
	if lazyInstance == nil {
		once.Do(func() {
			lazyInstance = &lazyPerson{}
		})
	}
	return lazyInstance
}
