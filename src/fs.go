package src

import (
	"os"
)

type ShizaFS struct {
}

func NewFS() *ShizaFS {
	return &ShizaFS{}
}

func (fs *ShizaFS) GetFileContent(fileName string) string {
	dataBytes, err := os.ReadFile(fileName)
	if err != nil {
		LogFSError(fileName, true)
	}
	return string(dataBytes)
}

func (fs *ShizaFS) bulkReadFile(fileNames []string) []string {
	return []string{}
}

func (fs *ShizaFS) generateFullFileTree() string {
	return ""
}
