package cache

type ICache[K, V any] interface {
	Get(K) (V, error)
	Set(K, V) error
}

type Element[V any] struct {
	value     V
	timestamp int64
}

type Cache[K comparable, V any] struct {
	cacheList  []func(K) (V, error)
	localCache map[K]Element[V]
}

func (c *Cache[K, V]) Get(k K) (V, error) {

}

func (c *Cache[K, V]) Put(k K, v V) error {

}
