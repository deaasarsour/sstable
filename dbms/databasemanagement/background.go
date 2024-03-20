package databasemanagement

func (databaseManagement *DatabaseManagement) InitializeBackgroundJobs() {
	go databaseManagement.fulledMemtableJob()
}

func (databaseManagement *DatabaseManagement) fulledMemtableJob() {
	for {

	}
}
