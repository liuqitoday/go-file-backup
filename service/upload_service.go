package service

import (
	"github.com/cheggaaa/pb/v3"
	"github.com/studio-b12/gowebdav"
	"go-file-buckup/config"
	"io/ioutil"
	"sync"
)

func UploadAllFiles(allFilePaths *[]string) {
	if len(*allFilePaths) <= 0 {
		return
	}
	// init webdav client
	root := config.WebdavUrl
	user := config.WebdavUsername
	password := config.WebdavPassword
	client := gowebdav.NewClient(root, user, password)
	// upload files
	group := sync.WaitGroup{}
	group.Add(len(*allFilePaths))
	processBar := pb.StartNew(len(*allFilePaths)).SetWidth(80)
	for _, filePath := range *allFilePaths {
		filePath := filePath
		go func() {
			uploadFile(filePath, client)
			group.Done()
			processBar.Increment()
		}()
	}
	group.Wait()
	processBar.Finish()
}

func uploadFile(filePath string, client *gowebdav.Client) error {
	bytes, _ := ioutil.ReadFile(filePath)
	err := client.Write(config.WebdavParentFilePath+filePath, bytes, 0644)
	return err
}
