package main

import (
	"github.com/urfave/cli/v2"
	"go-file-buckup/config"
	"go-file-buckup/service"
	"go-file-buckup/util"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "backup-file-path",
				Usage:       "文件备份目录",
				Aliases:     []string{"backupFilePath"},
				Destination: &config.BackupFilePath,
				Required:    true,
				Value:       "~/liuqi/home",
			},
			&cli.StringFlag{
				Name:        "webdav-parent-file-path",
				Usage:       "webdav 父目录",
				Aliases:     []string{"webdavParentFilePath"},
				Destination: &config.WebdavParentFilePath,
				Required:    true,
				Value:       "n1-sync",
			},
			&cli.StringFlag{
				Name:        "webdav-url",
				Usage:       "webdav URL",
				Aliases:     []string{"webdavUrl"},
				Required:    true,
				Destination: &config.WebdavUrl,
				Value:       "",
			},
			&cli.StringFlag{
				Name:        "webdav-username",
				Usage:       "webdav 用户名",
				Aliases:     []string{"webdavUsername"},
				Required:    true,
				Destination: &config.WebdavUsername,
				Value:       "",
			},
			&cli.StringFlag{
				Name:        "webdav-password",
				Usage:       "webdav 密码",
				Aliases:     []string{"webdavPassword"},
				Required:    true,
				Destination: &config.WebdavPassword,
				Value:       "",
			},
		},
		Action: func(c *cli.Context) error {
			filePaths, _ := util.GetAllFilePaths(config.BackupFilePath)
			service.UploadAllFiles(filePaths)
			return nil
		},
		Name:  "file backup",
		Usage: "file backup",
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
