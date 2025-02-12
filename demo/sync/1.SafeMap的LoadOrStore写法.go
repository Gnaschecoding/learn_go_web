package sync

import "sync"

type SafeMap[K comparable, V any] struct {
	values map[K]V
	lock   sync.RWMutex
}

// 如果有值返回原有的值，没有值返回装载进来
//
// 如果同时进来两个 goroutine 1 ——>("key1",1)
// 如果进来两个 goroutine 2 ——>("key1",2)
func (s *SafeMap[K, V]) LoadOrStore(key K, newval V) (V, bool) {
	s.lock.RLock()
	oldval, ok := s.values[key]
	s.lock.RUnlock()
	if ok {
		return oldval, true
	}
	//到这里肯定有一个 goroutine先进来 一个后进来，
	//那么goroutine1 先写入1   ，然后那么goroutine2 再写入2，2会把1给覆盖掉。因此违背了有值返回旧值的原则
	s.lock.Lock()
	defer s.lock.Unlock()
	//因此 需要在这里还要判断一下是否 已经有值了
	oldval, ok = s.values[key]
	if ok {
		return oldval, true
	}

	s.values[key] = newval
	return newval, false

}

func (s *SafeMap[K, V]) LoadOrStoreV2(key K, newval V) (V, bool) {
	s.lock.RLock()
	oldval, ok := s.values[key]
	s.lock.RUnlock()
	if ok {
		return oldval, true
	}
	//到这里肯定有一个 goroutine先进来 一个后进来，
	//那么goroutine1 先写入1   ，然后那么goroutine2 再写入2，2会把1给覆盖掉。因此违背了有值返回旧值的原则
	s.lock.Lock()
	defer s.lock.Unlock()
	//因此 需要在这里还要判断一下是否 已经有值了
	//oldval, ok = s.values[key]
	//if ok {
	//	return oldval, true
	//}

	s.values[key] = newval
	return newval, false

}
