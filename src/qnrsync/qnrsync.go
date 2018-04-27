package main

import (
	"myziyue"
	"strings"
	"qiniupkg.com/x/log.v7"
)


func main()  {
	WatchePath,err := myziyue.GetOption("WatcherPath", "watcher")

	if err != nil {
		log.Fatalf("请先设置监控目录", err)
		return
	}


	// 配置文件监控的子目录
	dirList := myziyue.GetWatcherPaths(strings.Split(WatchePath, ";"))
	// 配置文件监控目录
	for _,watcher := range strings.Split(WatchePath, ";") {
		dirList = append(dirList, watcher)
	}
	myziyue.SyncFile(dirList)
}
