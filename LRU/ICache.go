package LRU

type ICache interface {
	Get(key string) (interface{}, bool) //todo 提供公工具方法，针对数据类型
	Set(key string, value interface{})
	Iterate(func(key string, value interface{})) // 迭代器
}
