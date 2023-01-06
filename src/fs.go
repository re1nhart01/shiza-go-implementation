package src

import (
	"errors"
	"fmt"
	"os"
)

type ShizaFS struct {
}

func NewFS() *ShizaFS {
	return &ShizaFS{}
}

const (
	PathName         = "shiza_out"
	SourceCodePath   = "sauce"
	SourceCodeName   = "shiza_out.bmx"
	AssemblyCodeName = "shiza_out.asm"
)

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

func (fs *ShizaFS) bundleCompiledFile(sourceCode, path, content string) error {
	currentPath, err := os.Getwd()
	if err != nil {
		return errors.New("after getting working directory caught an error!")
	}
	mainPath := fmt.Sprintf("%s\\%s", currentPath, PathName)
	fullPath := fmt.Sprintf("%s\\%s", mainPath, SourceCodePath)

	err = os.MkdirAll(fullPath, 0777)
	if err != nil {
		return errors.New("after creating file structure caught an error!")
	}
	if err = os.WriteFile(fmt.Sprintf("%s\\%s", fullPath, SourceCodeName), []byte(sourceCode), 0777); err != nil {
		return errors.New("after generating file with source code caught an error!")
	}

	if err = os.WriteFile(fmt.Sprintf("%s\\%s", mainPath, AssemblyCodeName), []byte(content), 0777); err != nil {
		return errors.New("after generating final assembly file caught an error!")
	}

	return nil
}

func (fs *ShizaFS) GenerateOut(sourceCode, path, content string) {
	err := fs.bundleCompiledFile(sourceCode, path, content)
	if err != nil {
		LogBundlingError(err.Error(), true)
	}
}
