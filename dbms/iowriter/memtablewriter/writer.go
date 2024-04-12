package memtablewriter

import (
	"sstable/memtable"
	"sstable/types"
	"sstable/util/awaitable"

	"github.com/samber/lo"
)

type writerChanData struct {
	awaitableResultGroup *awaitable.AwaitableGroup[error]
	writeCommands        []types.KeyValueObject
	writeBytes           [][]byte
}

func (memtableWriter *MemtableWriterJob) WriterExec() {
	for {
		writerData := <-memtableWriter.writerChan
		err := memtableWriter.writerExecLogic(writerData)

		writerData.awaitableResultGroup.SetResult(err)
	}
}

func (memtableWriter *MemtableWriterJob) writerExecLogic(writerData writerChanData) error {

	state := memtableWriter.stateManagement.GetState()

	currentWindow := 0
	for currentWindow != len(writerData.writeBytes) {
		nextWindow := getNextWriteWindowStart(currentWindow, state.MemoryTable, writerData)

		bytes := lo.Flatten(writerData.writeBytes[currentWindow:nextWindow])
		writeCommands := writerData.writeCommands[currentWindow:nextWindow]

		if err := state.MemoryTable.WriteBatchRaw(writeCommands, bytes); err != nil {
			return err
		}

		if state.MemoryTable.IsFull() {
			if err := memtableWriter.memtableManagement.SwitchAndFlushMemtable(state); err != nil {
				return err
			}

			state = memtableWriter.stateManagement.GetState()
		}

		currentWindow = nextWindow
	}

	return nil

}

func getNextWriteWindowStart(startIndex int, memtable *memtable.MemoryTable, writerData writerChanData) int {
	totalBytes := 0
	for i := startIndex; i < len(writerData.writeBytes); i++ {
		totalBytes += len(writerData.writeBytes[i])
		if memtable.WillBeFullAfterWrite(totalBytes) {
			return i + 1
		}
	}

	return len(writerData.writeBytes)
}
