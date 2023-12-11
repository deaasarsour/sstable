package dbms

import "encoding/json"

type DatabaseMetadata struct {
	MemtableFilename string `json:"memtable_filename"`
}

func (dbms *DatabaseManagementSystem) initializeMetadata() error {
	if rawMetadata, err := dbms.storage.ReadMetadataRaw(); err == nil {
		var metadata DatabaseMetadata
		json.Unmarshal(rawMetadata, &metadata)
		dbms.cachedMetadata.Store(&metadata)
		return nil
	} else {
		return err
	}
}

// TODO change in future to command pattern
func (dbms *DatabaseManagementSystem) writeMetadataUnsafe() error {
	return dbms.storage.WriteMetadata(dbms.cachedMetadata.Load())
}
