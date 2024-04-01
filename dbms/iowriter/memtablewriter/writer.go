package memtablewriter

import (
	"sstable/dbms/state"
	"sstable/util"
)

type writerChanData struct {
	futureGroup   *util.FutureGroup[error]
	writeCommands []WriteCommand
}

func (memtableWriter *MemtableWriterJob) WriterExec() {
	for {
		writerData := <-memtableWriter.writerChan

		blockMemtableFlushing := func(dbState *state.DatabaseManagementState) {
			dbState.MemoryTable.LockFlushing()
		}

		state := memtableWriter.stateManagement.GetAtomicState(blockMemtableFlushing)

		var writeErr error = nil
		for _, writeCommand := range writerData.writeCommands {
			if err := state.MemoryTable.Write(writeCommand.Key, writeCommand.Value); err != nil {
				writeErr = err
			}
		}

		if writeErr == nil && state.MemoryTable.IsFull() {
			writeErr = memtableWriter.memtableManagement.SwitchAndFlushMemtable(state)
		}

		writerData.futureGroup.SetResult(writeErr)
	}
}
