package memtablewriter

import (
	"sstable/dbms/components/memtablemanagement"
	"sstable/dbms/statemanagement"
	"sstable/util"
)

const receiverChanSize = 1 << 10

type MemtableWriterJob struct {
	stateManagement    *statemanagement.DatabaseManagementStateManagement
	memtableManagement *memtablemanagement.MemtableManagement
	receiverChan       chan receiverChanData
	writerChan         chan writerChanData
}

type WriteCommand struct {
	Key   string
	Value any
}

func (memtableWriter *MemtableWriterJob) Initialize() {
	go util.RunInLoop(memtableWriter.ReceiverExec)
	go util.RunInLoop(memtableWriter.WriterExec)
}

func NewMemtableWriteJob(
	stateManagement *statemanagement.DatabaseManagementStateManagement,
	memtableManagement *memtablemanagement.MemtableManagement) *MemtableWriterJob {
	return &MemtableWriterJob{
		stateManagement:    stateManagement,
		memtableManagement: memtableManagement,
		receiverChan:       make(chan receiverChanData, receiverChanSize),
		writerChan:         make(chan writerChanData, 1),
	}
}
