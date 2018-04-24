# qiniu-auto-sync
七牛云存储自动上传

### 依赖类库

```bash
$ go get -v github.com/howeyc/fsnotify
$ go get -v github.com/larspensjo/config
$ go get -v github.com/qiniu/api.v7
```

### 交叉编译

1. Mac 下编译 Linux 和 Windows 64位可执行程序

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

2. Linux 下编译 Mac 和 Windows 64位可执行程序

```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

3. Windows 下编译 Mac 和 Linux 64位可执行程序

```bash
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build main.go

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go
```