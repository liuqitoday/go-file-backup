package util

import "io/ioutil"

var allFilePath []string

func GetAllFilePaths(path string) (*[]string, error) {
	dir, err := ioutil.ReadDir(path)
	for _, info := range dir {
		if info.IsDir() {
			GetAllFilePaths(path + "/" + info.Name())
		} else {
			allFilePath = append(allFilePath, path+"/"+info.Name())
		}
	}
	return &allFilePath, err
}
