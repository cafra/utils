package LRU

type Options struct {
	capability int // 缓存单元大小
}

type CacheOption func(*Options)

func WithCapability(capability int) CacheOption {
	return func(o *Options) {
		o.capability = capability
	}
}

var defaultCacheOptions = Options{
	capability: 1 << 10,
}
