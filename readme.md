### 使用webdav备份指定目录下所有文件

```
./go-file-backup -backupFilePath="~/n1-sync" -webdavParentFilePath="/n1-sync" -webdavUrl="https://dav.jianguoyun.com/dav/" -webdavUsername="username" -webdavPassword="password"
```

### 参数说明
- backupFilePath 备份目录
- webdavParentFilePath webdav目录
- webdavUrl webdav服务URL
- webdavUsername webdav服务用户名
- webdavPassword webdav服务密码