package dbms

func (dbms *DatabaseManagementSystem) Read(key string) any {
	memtable := dbms.memoryTable.Load()
	if result := memtable.Read(key); result != nil {
		return result
	} else {
		return nil
	}
}

func (dbms *DatabaseManagementSystem) Write(key string, value any) error {
	memtable := dbms.memoryTable.Load()
	return memtable.Write(key, value)
}
