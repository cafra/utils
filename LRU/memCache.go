package LRU

import (
	"container/list"
	"sync"
)

type memCache struct {
	sync.Map              // O(1) 查找
	DataList   *list.List // O(1) LRU
	capability int        // 缓存大小
}

// list element 结构
type itemElement struct {
	data interface{} //存储数据
	key  string      // map key
}

var once sync.Once
var memCacher ICache

/**
 * 获取指定大小的缓存
 *
 * param: ...CacheOption ops
 * return: ICache
 */
func GetMemCache(ops ...CacheOption) ICache {
	once.Do(func() {
		options := defaultCacheOptions
		for _, op := range ops {
			op(&options)
		}
		memCacher = &memCache{
			DataList:   list.New(),
			capability: options.capability,
		}
	})
	return memCacher
}

func (m *memCache) Get(key string) (value interface{}, exist bool) {
	// map 查找，命中则更新node 到root
	nodePtr, exist := m.Load(key)
	if !exist {
		return
	}
	ptr := nodePtr.(*list.Element)
	m.DataList.MoveToFront(ptr)
	item := ptr.Value.(*itemElement)
	value = item.data
	return
}

func (m *memCache) Set(key string, value interface{}) {
	// map 存在，则移动node 到root
	_, exist := m.Get(key)
	if !exist {
		// 不存在，则new node 到root
		item := &itemElement{
			data: value,
			key:  key,
		}
		ptr := m.DataList.PushFront(item)
		m.Store(key, ptr)
	}
	// 检查：如果队列长度超过长度，则删除超出部分，删除map key
	for m.DataList.Len() > m.capability {
		element := m.DataList.Back()
		item := element.Value.(*itemElement)
		m.Delete(item.key)
		m.DataList.Remove(element)
	}
	return
}
func (m *memCache) Iterate(hundler func(key string, value interface{})) {
	for e := m.DataList.Front(); e != nil; e = e.Next() {
		item := e.Value.(*itemElement)
		hundler(item.key, item.data)
	}
}
