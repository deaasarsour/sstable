package memtablewriter

import (
	"sstable/util"
)

type writerChanData struct {
	futureGroup   *util.FutureGroup[error]
	writeCommands []WriteCommand
}

func (memtableWriter *MemtableWriterJob) WriterExec() {
	for {
		writerData := <-memtableWriter.writerChan

		state := memtableWriter.stateManagement.GetState()

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
