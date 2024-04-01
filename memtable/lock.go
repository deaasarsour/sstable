package memtable

func (memtable *MemoryTable) LockFlushing() {
	memtable.flushLock.Store(true)
}

func (memtable *MemoryTable) UnlockFlushing() {
	memtable.flushLock.Store(false)
}
