package daily

// LRUCache FIXME 有问题
type LRUCache struct {
	size, capacity int
	cache          map[int]*LRUNode
	head, tail     *LRUNode
}

type LRUNode struct {
	key, value int
	pre, next  *LRUNode
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		size:     0,
		capacity: capacity,
		cache:    make(map[int]*LRUNode),
	}
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}

	liftNode(node, this)
	return node.value
}

// 把 node 放到最前面
func liftNode(node *LRUNode, this *LRUCache) {
	// 不是最近的一个
	if node != this.head {

		// 处理 node 的前一个
		if node == this.tail {
			this.tail = node.next
		} else {
			node.pre.next = node.next
		}

		// 把 node 升到最近的一个
		node.next = nil
		node.pre = this.head
		this.head.next = node
		this.head = node
	}
}

func (this *LRUCache) Put(key int, value int) {
	// key 已经存在了，替换 value
	if node, ok := this.cache[key]; ok {
		node.value = value
		liftNode(node, this)
		return
	}

	// key 不存在
	node := &LRUNode{key: key, value: value}
	this.cache[key] = node

	// 未达容量
	if this.size != this.capacity {
		this.size++
		// 第一个
		if this.size == 1 {
			this.head = node
			this.tail = node
		} else {
			node.pre = this.head
			this.head.next = node
			this.head = node
		}
	} else {
		// 达到容量
		delete(this.cache, this.tail.key)
		if this.capacity == 1 {
			this.head = node
			this.tail = node
		} else {
			node.pre = this.head
			this.head.next = node
			this.head = node
			this.tail = this.tail.next
		}
	}
}
