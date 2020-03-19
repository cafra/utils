package LRU

type ICache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
	Iterate(func(key string, value interface{})) // 迭代器
	Close()
}
