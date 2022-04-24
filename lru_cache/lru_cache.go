package lru_cache

import (
	"errors"
	"sync"
)

// we use hashMap to access a key with O(1),
// and double linked list with this behavior: head -> least recently used key
// tail -> most recently used. When user access a key if key existed, we move that key to the tail of
// linked list, if cache is full, we delete the head of linked list

type (
	LRUCache interface {
		Get(key string) ([]byte, error)
		Set(key string, value []byte)
	}

	node struct {
		Val  []byte
		Key  string
		Prev *node
		Next *node
	}

	lRUCacheImpl struct {
		Cache         map[string]*node
		Capacity      int
		Size          int
		LRULinkedList *node
		head          *node
		tail          *node
		mutex         sync.RWMutex
	}
)

func NewLRUCache(size int) LRUCache {
	return &lRUCacheImpl{
		Size:  size,
		Cache: make(map[string]*node),
	}
}

func (l *lRUCacheImpl) removeHead() {
	l.mutex.Lock()
	delete(l.Cache, l.head.Key)
	l.head = l.head.Next
	if l.head != nil {
		l.head.Prev = nil
	}
	l.mutex.Unlock()
}

func (l *lRUCacheImpl) setTail(n *node) {
	l.mutex.Lock()
	if l.tail != nil {
		l.tail.Next = n
		n.Prev = l.tail
	}
	l.tail = n
	n.Next = nil
	if l.head == nil {
		l.head = n
	}
	l.mutex.Unlock()
}

func (l *lRUCacheImpl) moveToTail(n *node) {
	if n == l.tail {
		return
	}
	next := n.Next
	prev := n.Prev
	if prev != nil {
		prev.Next = next
	} else {
		// n is head
		l.head = next
		l.head.Prev = nil
	}
	n.Next = nil
	n.Prev = nil

	l.setTail(n)
}

func (l *lRUCacheImpl) pushToEnd(n *node) {
	l.mutex.RLock()
	if l.Capacity == l.Size {
		l.mutex.RUnlock()
		l.removeHead()
	} else {
		l.mutex.RUnlock()
		l.mutex.Lock()
		l.Capacity += 1
		l.mutex.Unlock()
	}
	l.setTail(n)
}

func (l *lRUCacheImpl) Get(key string) ([]byte, error) {
	if n, ok := l.Cache[key]; ok {
		l.moveToTail(n)
		return n.Val, nil
	}
	return nil, errors.New("not found")
}

func (l *lRUCacheImpl) Set(key string, value []byte) {
	n, ok := l.Cache[key]
	if ok {
		n.Val = value
		l.moveToTail(n)
	} else {
		n := &node{Val: value, Key: key}
		l.pushToEnd(n)
		l.Cache[key] = n
	}
}
