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

		keyValues := make([]util.KeyValueObject, 0)
		for _, writeCommand := range writerData.writeCommands {
			keyValues = append(keyValues, util.KeyValueObject{
				Key:   writeCommand.Key,
				Value: writeCommand.Value,
			})
		}

		if err := state.MemoryTable.WriteBatch(keyValues); err != nil {
			writeErr = err
		}

		if writeErr == nil && state.MemoryTable.IsFull() {
			writeErr = memtableWriter.memtableManagement.SwitchAndFlushMemtable(state)
		}

		writerData.futureGroup.SetResult(writeErr)
	}
}
