// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2022-01-12 11:44
// version: 1.0.0
// desc   : 协程局部存储，类型 ThreadLocal

package concurrent

import "sync"

// Local 局部存储类
type Local struct {
	store *storeMap
}

// storeMap 内部实现
type storeMap struct {
	sync.Mutex
	mp map[string]map[interface{}]interface{}
}

// NewLocal 创建一个实例
func NewLocal() *Local {
	return &Local{
		store: &storeMap{
			mp: make(map[string]map[interface{}]interface{}),
		},
	}
}

// Get 获取一个值
func (l *Local) Get(key interface{}) interface{} {
	return l.getMap()[key]
}

// Put 保存一个值
func (l *Local) Put(key interface{}, value interface{}) {
	l.getMap()[key] = value
}

// Remove 移除一个值
func (l *Local) Remove(key interface{}) {
	delete(l.getMap(), key)
}

// Clear 清除当前协程上的数据
func (l *Local) Clear() {
	l.store.Lock()
	defer l.store.Unlock()

	delete(l.store.mp, GoroutineID())
}

// getMap 获取当前协程上的全部存储数据
func (l *Local) getMap() map[interface{}]interface{} {
	l.store.Lock()
	defer l.store.Unlock()

	goID := GoroutineID()
	if m, _ := l.store.mp[goID]; m != nil {
		return m
	}

	m := make(map[interface{}]interface{})
	l.store.mp[goID] = m
	return m
}
