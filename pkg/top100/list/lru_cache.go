package list

// NO.146
type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DListNode
	head, tail *DListNode
}

type DListNode struct {
	Key        int
	Val        int
	Prev, Next *DListNode
}

func Constructor(capacity int) LRUCache {
	lru := &LRUCache{
		capacity: capacity,
		size:     0,
		cache:    make(map[int]*DListNode),
	}
	return *lru
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		if v == this.head {
			return v.Val
		}

		if v == this.tail {
			//移动
			prev := this.tail.Prev
			this.tail.Prev.Next = nil //删除tail
			this.tail.Next = this.head
			this.head.Prev = this.tail
			this.cache[this.head.Key] = this.head
			this.tail.Prev = nil
			this.head = this.tail

			this.tail = prev
			return v.Val

		}

		v.Prev.Next = v.Next
		v.Next.Prev = v.Prev
		v.Next = this.head
		this.head.Prev = v
		this.cache[this.head.Key] = this.head
		v.Prev = nil
		this.head = v

		return v.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.cache[key] != nil {
		this.cache[key].Val = value
		this.Get(key)
		return
	}

	if this.size == 0 {
		this.head = &DListNode{
			Key: key,
			Val: value,
		}
		this.tail = this.head
		this.cache[key] = this.head
		this.size++
		return
	}

	//满了
	if this.size == this.capacity {
		delete(this.cache, this.tail.Key)
		if this.size == 1 {
			this.head, this.tail = nil, nil
			this.size--
			this.head = &DListNode{
				Key: key,
				Val: value,
			}
			this.tail = this.head
			this.cache[key] = this.head
			this.size++
		} else {
			this.tail = this.tail.Prev
			this.tail.Next = nil
			this.size--
			node := &DListNode{
				Key: key,
				Val: value,
			}
			this.cache[key] = node
			node.Next = this.head
			this.head.Prev = node
			this.head = node
			this.size++
		}
	}

	if this.size < this.capacity {
		node := &DListNode{
			Key: key,
			Val: value,
		}
		this.cache[key] = node
		node.Next = this.head
		this.head.Prev = node
		this.cache[this.head.Key] = this.head //更新head
		this.head = node
		this.size++
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
