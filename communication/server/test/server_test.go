package server_test

import (
	"fmt"
	"os/exec"
	"sstable/communication/server"
	"sstable/test/util/testdbms"
	"testing"

	"github.com/stretchr/testify/assert"
)

func runCommand(command string) string {

	cmd := exec.Command("bash", "-c", command)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}

	return string(output)
}

func runReadCommand(key string) string {
	command := fmt.Sprintf(`printf "%%b" "rd%v\x0a\x0a" | nc localhost 62151 | cut -c 3-`, key)
	return runCommand(command)
}
func runWriteCommand(key string, value string) string {
	command := fmt.Sprintf(`printf "%%b" "wd{\"key\":\"%v\",\"value\":\"%v\"}\\x0a\\x0a" | nc localhost 62151 | cut -c 3-`, key, value)
	return runCommand(command)
}

func closeListener(dbServer *server.DatabaseServer) {
	if dbServer.Listener != nil {
		dbServer.Listener.Close()
	}
}

func TestServer(t *testing.T) {

	//arrange
	dbms := testdbms.NewDummyDbms(nil)
	dbms.Initialize()
	dbServer := server.NewDatabaseServer("127.0.0.1", 62151, dbms)

	//act
	defer closeListener(dbServer)
	go dbServer.StartListen()

	emptyRead := runReadCommand("deea")
	runWriteCommand("deea", "1")
	runWriteCommand("deea", "2")
	runWriteCommand("player", "hero")
	deeaRead := runReadCommand("deea")
	playerRead := runReadCommand("player")

	//assert
	assert.Equal(t, "<nil>\n", emptyRead)
	assert.Equal(t, "2\n", deeaRead)
	assert.Equal(t, "hero\n", playerRead)

}
