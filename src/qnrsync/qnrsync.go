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

	myziyue.SyncFile(myziyue.GetWatcherPaths(strings.Split(WatchePath, ";")))
}
