package main

import (
	"sstable/communication/server"
	"sstable/dbms/core"
	filesystemos "sstable/osfilesystem"
)

func main() {

	folder := filesystemos.NewOsDirectory("/Users/deaa/dbtest")

	dbms, err := core.NewDatabaseManagedSystem(folder)
	if err != nil {
		panic(err)
	}

	if err := dbms.Initialize(); err != nil {
		panic(err)
	}

	dbServer := server.NewDatabaseServer("127.0.0.1", 3009, dbms)
	dbServer.StartListen()
}
