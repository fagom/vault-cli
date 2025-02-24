package utils

import (
	"log"
	"os"
	"path"
	"runtime"
)

func GetOs() string {
	return runtime.GOOS
}

func GetHomeDir() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Unable to get user home directory. Exiting")
		return "", err
	}

	return dir, nil
}

func GetWorkingDir() (string, error) {
	home, err := GetHomeDir()
	if err != nil {
		return "", err
	}

	workDirPath := path.Join(home, "vault")

	os.MkdirAll(workDirPath, os.ModePerm)

	return workDirPath, nil
}
