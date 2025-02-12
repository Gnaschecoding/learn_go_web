package sync

import "sync"

type SafeArrayList[T any] struct {
	List []T
	lock sync.RWMutex
}

func (sl *SafeArrayList[T]) Get() T {
	sl.lock.RLock()
	defer sl.lock.RUnlock()
	if len(sl.List) == 0 {
		panic("len(sl.List) == 0")
	} else {
		return sl.List[len(sl.List)-1]
	}
}

func (sl *SafeArrayList[T]) Add(elem T) {
	sl.lock.Lock()
	defer sl.lock.Unlock()
	sl.List = append(sl.List, elem)
}
