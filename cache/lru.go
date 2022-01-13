// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2022-01-13 14:24
// version: 1.0.0
// desc   : LRU 缓存实现

package cache

import (
	"container/list"
	"fmt"
	"strings"
	"sync"
)

// LruCache 缓存类
type LruCache struct {
	lock     sync.Mutex
	size     int
	values   *list.List
	cacheMap map[interface{}]*list.Element
}

// NewLruCache 创建一个新的实例
func NewLruCache(size int) *LruCache {
	return &LruCache{
		size:     size,
		values:   list.New(),
		cacheMap: make(map[interface{}]*list.Element, size),
	}
}

// Put 添加元素
func (lc *LruCache) Put(key, value interface{}) {
	lc.lock.Lock()
	defer lc.lock.Unlock()
	// 是否满
	if lc.values.Len() == lc.size {
		back := lc.values.Back()
		lc.values.Remove(back)
		delete(lc.cacheMap, back)
	}
	front := lc.values.PushFront(value)
	lc.cacheMap[key] = front
}

// Get 获取某个元素
func (lc *LruCache) Get(key interface{}) (value interface{}, ok bool) {
	val, ok := lc.cacheMap[key]
	if ok {
		lc.values.MoveToFront(val)
		value = val.Value
	}
	return
}

// Remove 移除某个元素
func (lc *LruCache) Remove(key interface{}) (value interface{}, ok bool) {
	lc.lock.Lock()
	defer lc.lock.Unlock()

	val, ok := lc.cacheMap[key]
	if ok {
		lc.values.Remove(val)
		delete(lc.cacheMap, val)
		value = val.Value
	}
	return
}

// Size 当前有效缓存的数量
func (lc *LruCache) Size() int {
	return lc.values.Len()
}

// String 字符串输出
func (lc *LruCache) String() string {
	var sb strings.Builder
	for i := lc.values.Front(); i != nil; i = i.Next() {
		sb.WriteString(fmt.Sprintln(i.Value))
	}
	return sb.String()
}

// List 当前有效缓存数据
func (lc *LruCache) List() []interface{} {
	var data []interface{}
	for i := lc.values.Front(); i != nil; i = i.Next() {
		data = append(data, i.Value)
	}
	return data
}

// Clear 清空缓存
func (lc *LruCache) Clear() {
	lc.lock.Lock()
	defer lc.lock.Unlock()
	lc.values = list.New()
	lc.cacheMap = make(map[interface{}]*list.Element, lc.size)
}
