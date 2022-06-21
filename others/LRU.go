package others

// 自顶向下写：先默写模版，再实现双向链表
type LRUCache struct {
	HashMap    map[int]*node
	DoubleList *doubleList // 双向链表
	Cap        int         // 最大容量
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		HashMap:    make(map[int]*node),
		DoubleList: newDoubleList(),
		Cap:        capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.HashMap[key]; !ok {
		return -1
	}
	// 将该数据提升为最近使用的
	this.makeRecently(key)
	return this.HashMap[key].val
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.HashMap[key]; ok {
		//  删除旧数据
		this.deleteKey(key)
		// 新插入的数据为最近使用的数据
		this.addRecently(key, value)
		return
	}
	if this.size() == this.Cap {
		// 删除最久未使用的元素
		this.removeLeastRecently()
	}
	this.addRecently(key, value)
}

func (this *LRUCache) size() int {
	return len(this.HashMap)
}

// 将某个key提升为最近使用的
func (this *LRUCache) makeRecently(key int) {
	node := this.HashMap[key]
	// 先删除
	this.DoubleList.remove(node)
	// 再插入到队尾
	this.DoubleList.addLast(node)
}

// 添加最近使用的元素
func (this *LRUCache) addRecently(key, value int) {
	node := &node{key: key, val: value}
	this.DoubleList.addLast(node)
	this.HashMap[key] = node
}

// 删除某一个key
func (this *LRUCache) deleteKey(key int) {
	node := this.HashMap[key]
	delete(this.HashMap, key)
	this.DoubleList.remove(node)
}

// 删除最久未使用的元素
func (this *LRUCache) removeLeastRecently() {
	node := this.DoubleList.removeFirst()
	delete(this.HashMap, node.key)
}

// 实现双向链表
type node struct {
	prev     *node
	next     *node
	key, val int
}
type doubleList struct {
	// 头尾虚节点
	head, tail *node
	// 链表元素数
	size int
}

func newDoubleList() *doubleList {
	head := &node{prev: nil, next: nil, key: 0, val: 0}
	tail := &node{prev: nil, next: nil, key: 0, val: 0}
	head.next = tail
	tail.prev = head
	return &doubleList{
		head: head,
		tail: tail,
		size: 0,
	}
}

// 链表尾部添加节点x，时间复杂度为O(1)
func (dl *doubleList) addLast(x *node) {
	x.prev = dl.tail.prev
	x.next = dl.tail
	dl.tail.prev.next = x
	dl.tail.prev = x
	dl.size += 1
}

// 删除链表中的x节点（x一定存在）
// 由于是双链表且给的是目标node节点，时间复杂度为O(1)
func (dl *doubleList) remove(x *node) {
	x.prev.next = x.next
	x.next.prev = x.prev
	dl.size -= 1
}

// 删除链表中的第一个节点，并返回该节点，时间复杂度为O(1)
func (dl *doubleList) removeFirst() *node {
	if dl.head.next == dl.tail {
		return nil
	}
	first := dl.head.next
	dl.remove(first)
	return first
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
