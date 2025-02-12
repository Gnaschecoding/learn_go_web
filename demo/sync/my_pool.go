package sync

import (
	"sync"
	"unsafe"
)

type MyPool struct {
	p      sync.Pool
	maxctx int32
	ctx    int32
}

func (p *MyPool) Get() interface{} {
	return p.p.Get()
}

func (p *MyPool) Put(val any) {
	if unsafe.Sizeof(val) > 1024 {
		return
	}
	//下面这个不确定性，因为gc也会释放资源，对象在垃圾回收时可能会被清除，所以没有用，要根据具体情况来看
	//atomic.AddInt32(&p.ctx, 1)
	//if p.ctx >= p.maxctx {
	//	atomic.AddInt32(&p.ctx, -1)
	//	return
	//}

	p.p.Put(val)
}
