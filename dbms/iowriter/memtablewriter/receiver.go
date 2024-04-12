package memtablewriter

import (
	"sstable/memtable"
	"sstable/types"
	"sstable/util/awaitable"
	"sstable/util/channelutil"
)

type receiverChanData struct {
	writeCommand types.KeyValueObject
	awaitable    *awaitable.Awaitable[error]
}

func (memtableWriter *MemtableWriterJob) ReceiverExec() {
	for {
		receiverDataBatch := channelutil.ReadBatch(memtableWriter.receiverChan, receiverChanSize)
		batchSize := len(receiverDataBatch)
		writerData := writerChanData{
			writeBytes:    make([][]byte, 0, batchSize),
			writeCommands: make([]types.KeyValueObject, 0, batchSize),
		}

		awaitables := make([]*awaitable.Awaitable[error], 0, batchSize)

		for i := range receiverDataBatch {

			awaitable := receiverDataBatch[i].awaitable
			writeCommand := receiverDataBatch[i].writeCommand

			if bytes, err := memtable.GetWriteByte(writeCommand); err == nil {
				awaitables = append(awaitables, awaitable)
				writerData.writeBytes = append(writerData.writeBytes, bytes)
				writerData.writeCommands = append(writerData.writeCommands, writeCommand)
			} else {
				receiverDataBatch[i].awaitable.SetResult(err)
			}
		}

		writerData.awaitableResultGroup = awaitable.NewAwaitableGroup(awaitables)
		memtableWriter.writerChan <- writerData
	}
}

func (memtableWriter *MemtableWriterJob) Write(writeCommand types.KeyValueObject) error {

	future := awaitable.NewAwaitable[error]()
	receiverData := receiverChanData{
		writeCommand: writeCommand,
		awaitable:    future,
	}

	memtableWriter.receiverChan <- receiverData

	return future.GetResult()
}
