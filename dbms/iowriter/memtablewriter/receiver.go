package memtablewriter

import (
	"sstable/util"
)

type receiverChanData struct {
	writeCommand WriteCommand
	future       *util.Future[error]
}

func (memtableWriter *MemtableWriterJob) ReceiverExec() {

	for {
		receiverDataBatch := util.ReadBatch(memtableWriter.receiverChan, receiverChanSize)
		batchSize := len(receiverDataBatch)

		writerData := writerChanData{
			writeCommands: make([]WriteCommand, batchSize),
		}

		futures := make([]*util.Future[error], batchSize)

		for i := range receiverDataBatch {
			futures[i] = receiverDataBatch[i].future
			writerData.writeCommands[i] = receiverDataBatch[i].writeCommand
		}

		writerData.futureGroup = util.NewFutureGroup(futures)

		memtableWriter.writerChan <- writerData
	}
}

func (memtableWriter *MemtableWriterJob) Write(writeCommand WriteCommand) error {

	future := util.NewFuture[error]()
	receiverData := receiverChanData{
		writeCommand: writeCommand,
		future:       future,
	}

	memtableWriter.receiverChan <- receiverData

	return future.GetResult()
}
