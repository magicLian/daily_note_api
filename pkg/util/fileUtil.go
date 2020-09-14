package util

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

var (
	PathNotFound = errors.New("path not found")
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreatePath(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func WriteFile(content []byte, filename string) error {
	return ioutil.WriteFile(filename, content, 0777)
}

func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func DeletePathAll(path string) error {
	return os.RemoveAll(path)
}


/**
CleanPath
clear the user tmp path for a single helm operation
*/
func CleanPath(path string) {
	log.Info("clean path = " + path)
	err := DeletePathAll(path)
	if err != nil {
		log.Error("clean path failed :" + err.Error())
	}
}

/**
CreatefileValuesAndConfig creates the file for POST,DELETE
    values.yaml : tmpPath["workPath"]/values.yaml
    config      : tmpPath["workPath"]/config
    vaules.yaml file and config file will be saved and they can be `covered` by new one
**/
func CreatefileValuesAndConfig(path, yaml, config string) error {
	if yaml != "" {
		err := CreateAndWritefile(path, "values.yaml", yaml)
		if err != nil {
			return errors.New("create file failed,reason : values.yaml create failed")
		}
	}
	err := CreateAndWritefile(path, "config", config)
	if err != nil {
		return errors.New("create file failed,reason : values.yaml create failed")
	}
	return nil
}

func CreateAndWritefile(path, name, content string) error {
	b, _ := PathExists(path)
	if !b {
		err := CreatePath(path)
		if err != nil {
			log.Error("create path failed")
		}
	}
	filename := path + "/" + name
	filev, err := os.Create(filename)
	if err != nil {
		log.Error("file create failed: " + filename)
		return errors.New("file create failed")
	}
	defer filev.Close()
	err = WriteFile([]byte(content), filename)
	if err != nil {
		log.Error("values file write failed file:" + filename + " content:" + content)
		return errors.New("values file write failed")
	}
	return nil
}
