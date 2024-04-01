package main

import (
	"fmt"
	"os"
	"sstable/dbms/core"
	databasereader "sstable/dbms/ioreader"
	filesystemos "sstable/osfilesystem"
	"sync"
	"time"
)

func bulkload(dbReader *databasereader.DatabaseReader) {
	var wg sync.WaitGroup
	write := func(i int) {
		wg.Add(1)
		key := fmt.Sprintf("key_%v", i)
		dbReader.Write(key, key)
		wg.Done()
	}
	const size = 1_000_000
	for i := 0; i < size; i++ {
		go write(i)
	}

	wg.Wait()
}

func main() {
	tmp, _ := os.MkdirTemp("", "")
	defer os.Remove(tmp)

	directory := filesystemos.NewOsDirectory(tmp)

	dbms, err := core.NewDatabaseManagedSystem(directory)
	if err != nil {
		panic(err)
	}

	if err := dbms.Initialize(); err != nil {
		panic(err)
	}

	startTime := time.Now()

	bulkload(dbms.DatabaseReader)

	endTime := time.Now()

	executionTime := endTime.Sub(startTime)

	fmt.Println("Execution time:", executionTime)
}
