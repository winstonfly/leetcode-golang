package list

import "testing"

func TestLruCache(t *testing.T) {
	t.Run("lruCache", func(t *testing.T) {
		lru := Constructor(2)
		lru.Put(1, 1)
		lru.Put(2, 2)
		v := lru.Get(1)
		t.Log(v)
		lru.Put(3, 3)
		v = lru.Get(2)
		t.Log(v)
		lru.Put(4, 4)
		v = lru.Get(1)
		t.Log(v)
		v = lru.Get(3)
		t.Log(v)
		v = lru.Get(4)
		t.Log(v)
	})

	t.Run("lruCache with capacity 1", func(t *testing.T) {
		lru := Constructor(1)
		lru.Put(2, 1)
		v := lru.Get(2)
		t.Log(v)
		lru.Put(3, 2)
		v = lru.Get(2)
		t.Log(v)
		v = lru.Get(3)
		t.Log(v)
	})

	t.Run("lruCache with capacity 0", func(t *testing.T) {
		lru := Constructor(2)
		v := lru.Get(2)
		t.Log(v)
		lru.Put(2, 6)
		v = lru.Get(1)
		t.Log(v)
		lru.Put(1, 5)
		lru.Put(1, 2)
		v = lru.Get(1)
		t.Log(v)
		v = lru.Get(2)
		t.Log(v)

	})

	t.Run("lruCache with capacity 3", func(t *testing.T) {
		lru := Constructor(3)
		lru.Put(1, 1)
		lru.Put(2, 2)
		lru.Put(3, 3)
		lru.Put(4, 4)
		v := lru.Get(4)
		t.Log(v)
		v = lru.Get(2)
		t.Log(v)
		v = lru.Get(2)
		t.Log(v)
		v = lru.Get(1)
		t.Log(v)
		lru.Put(5, 5)
		v = lru.Get(1)
		t.Log(v)
		v = lru.Get(2)
		t.Log(v)
		v = lru.Get(3)
		t.Log(v)
		v = lru.Get(4)
		t.Log(v)
		v = lru.Get(5)

	})
}
