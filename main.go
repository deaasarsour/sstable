package main

import (
	"sstable/dbms/core"
	filesystemos "sstable/osfilesystem"
	"sstable/server"
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

	dbServer := server.NewDatabaseServer("0.0.0.0", 3009, dbms)
	dbServer.StartListen()
}
