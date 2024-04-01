package main

import (
	"fmt"
	"os"
	"sstable/dbms/core"
	"sstable/dbms/databasemanagement"
	filesystemos "sstable/osfilesystem"
	"sync"
	"time"
)

func bulkload(db *databasemanagement.DatabaseManagement) {
	var wg sync.WaitGroup
	write := func(i int) {
		wg.Add(1)
		key := fmt.Sprintf("key_%v", i)
		db.Write(key, key)
		wg.Done()
	}
	const size = 500_000
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

	bulkload(dbms.DatabaseManagement)

	endTime := time.Now()

	executionTime := endTime.Sub(startTime)

	fmt.Println("Execution time:", executionTime)
}
