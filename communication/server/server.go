package server

import (
	"fmt"
	"net"
	"sstable/dbms/core"
)

type DatabaseServer struct {
	host     string
	port     int
	dbms     *core.DatabaseManagementSystem
	Listener net.Listener
}

func (server *DatabaseServer) StartListen() {
	url := fmt.Sprintf("%v:%v", server.host, server.port)

	listener, err := net.Listen("tcp", url)
	if err != nil {
		panic(err)
	}

	defer listener.Close()
	fmt.Printf("Server is listening on port %v\n", server.port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		dbConnection := NewDatabaseConnection(conn, server.dbms)
		go dbConnection.startCommunication()
	}
}

func NewDatabaseServer(host string, port int, dbms *core.DatabaseManagementSystem) *DatabaseServer {
	return &DatabaseServer{
		host: host,
		port: port,
		dbms: dbms,
	}
}
