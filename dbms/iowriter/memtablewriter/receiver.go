package memtablewriter

import (
	"sstable/memtable"
	"sstable/util"
)

type receiverChanData struct {
	writeCommand util.KeyValueObject
	future       *util.Future[error]
}

func (memtableWriter *MemtableWriterJob) ReceiverExec() {
	cnt := 0
	for {
		receiverDataBatch := util.ReadBatch(memtableWriter.receiverChan, receiverChanSize)
		batchSize := len(receiverDataBatch)
		cnt += batchSize
		writerData := writerChanData{
			writeBytes:    make([][]byte, 0, batchSize),
			writeCommands: make([]util.KeyValueObject, 0, batchSize),
		}

		futures := make([]*util.Future[error], 0, batchSize)

		for i := range receiverDataBatch {

			future := receiverDataBatch[i].future
			writeCommand := receiverDataBatch[i].writeCommand

			if bytes, err := memtable.GetWriteByte(writeCommand); err == nil {
				futures = append(futures, future)
				writerData.writeBytes = append(writerData.writeBytes, bytes)
				writerData.writeCommands = append(writerData.writeCommands, writeCommand)
			} else {
				receiverDataBatch[i].future.SetResult(err)
			}
		}

		writerData.futureGroup = util.NewFutureGroup(futures)
		memtableWriter.writerChan <- writerData
	}
}

func (memtableWriter *MemtableWriterJob) Write(writeCommand util.KeyValueObject) error {

	future := util.NewFuture[error]()
	receiverData := receiverChanData{
		writeCommand: writeCommand,
		future:       future,
	}

	memtableWriter.receiverChan <- receiverData

	return future.GetResult()
}
