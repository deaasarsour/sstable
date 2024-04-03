package main

import (
	"fmt"
	"os"
	"sstable/dbms/core"
	"sstable/dbms/databaseio"
	filesystemos "sstable/osfilesystem"
	"sync"
	"time"
)

func bulkload(io *databaseio.DatabaseIO) {
	var wg sync.WaitGroup
	write := func(i int) {
		key := fmt.Sprintf("key_%v", i)
		io.Write(key, key)
		wg.Done()
	}
	const size = 1_000_000
	wg.Add(size)
	for i := 0; i < size; i++ {
		go write(i)
	}

	wg.Wait()
}

func main() {
	tmp, _ := os.MkdirTemp("", "")
	defer os.Remove(tmp)
	fmt.Printf("tmp folder : %v\n", tmp)
	directory := filesystemos.NewOsDirectory(tmp)

	dbms, err := core.NewDatabaseManagedSystem(directory)
	if err != nil {
		panic(err)
	}

	if err := dbms.Initialize(); err != nil {
		panic(err)
	}

	startTime := time.Now()

	bulkload(dbms.DatabaseIO)

	endTime := time.Now()

	executionTime := endTime.Sub(startTime)

	fmt.Println("Execution time:", executionTime)
}
