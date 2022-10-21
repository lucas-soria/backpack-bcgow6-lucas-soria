package store

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Type string

const (
	FileType  Type = "file"
	MongoType Type = "mongo"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type fileStore struct {
	filepath string
}

func NewStore(storeType Type, path string) Store {
	switch storeType {
	case FileType:
		return &fileStore{filepath: path}
	default:
		log.Fatalf("Only '%s' Store is implemented", FileType)
	}
	return nil
}

func (f *fileStore) Read(data interface{}) (err error) {
	textBytes, err := os.ReadFile(f.filepath)
	if err != nil {
		log.Fatalf("Error reading file:\n%+v", err)
	}
	return json.Unmarshal(textBytes, &data)
}

func (f *fileStore) Write(data interface{}) (err error) {
	fileData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return fmt.Errorf("Error writing file:\n%+v", err)
	}
	return os.WriteFile(f.filepath, fileData, 0777)
}
