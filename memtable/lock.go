package memtable

func (memtable *MemoryTable) LockFlushing() {
	memtable.flushLock.Store(true)
}

func (memtable *MemoryTable) UnlockFlushing() {
	memtable.flushLock.Store(false)
}

func (memtable *MemoryTable) CanFlushing() bool {
	return !memtable.flushLock.Load()
}
