package myziyue

import (
	"github.com/howeyc/fsnotify"
	"log"
	"os"
	"io/ioutil"
)

type monitor struct {
	watch *fsnotify.Watcher
}

func SyncFile(WatchPaths []string)  {
	M, err := NewMonitor()
	if err != nil {
		log.Println(err)
		return
	}
	M.Do()
	for i:=0;i<len(WatchPaths);i++{
		M.watch.Watch(WatchPaths[i])
		log.Println("开始监控：", WatchPaths[i])
	}
	select {}
}

func RmWatcher(watchPath string) {
	M, err := NewMonitor()
	if err != nil {
		log.Println(err)
		return
	}
	M.Do()
	M.watch.RemoveWatch(watchPath)
	log.Println("移除监控：", watchPath)
	return
}

func NewMonitor() (monitor, error) {
	Mon, err := fsnotify.NewWatcher()
	return monitor{Mon}, err
}
func (self monitor) Do() {
	go func() {
		for {
			select {
			case w := <-self.watch.Event:
				if w.IsCreate() {
					log.Println(w.Name, " 文件被创建.")
					//UploadFile(w.Name, w.Name)
					fileName := GetFilePath(w.Name)

					if fileIsDir(w.Name) {
						SyncFile([]string{w.Name})
						continue
					}
					UploadFile(w.Name, fileName)
					continue
				}

				if w.IsModify() {
					//log.Println(w.Name, " 文件被修改.")
					continue
				}

				if w.IsDelete() {
					//log.Println(w.Name, " 文件被删除.")
					//fileName := GetFilePath(w.Name)
					//log.Println(fileName + " is delete")

					//if fileIsDir(w.Name) {
					//	log.Println(fileName + " is dir")
					//	RmWatcher(w.Name)
					//	continue
					//}

					//DeleteFile(fileName)
					continue
				}

				if w.IsRename() {
					//w = <-self.watch.Event
					//log.Println(w)
					//self.watch.RemoveWatch(w.Name)
					//log.Println(w.Name, " 被重命名.")

					//fileName := GetFilePath(w.Name)
					//
					//if fileIsDir(w.Name) {
					//	continue
					//}

				}

			case err := <-self.watch.Error:
				log.Fatalln(err)
			}
		}
	}()
}


func fileIsDir(filename string) bool {
	fileHandle, err := os.Stat(filename)
	if err != nil {
		log.Println(err)
		return false
	}
	return fileHandle.IsDir()
}

func GetWatcherPaths(watchPath []string) []string {
	var dirNames []string
	for i:=0;i<len(watchPath);i++ {
		// 读取目录下的子目录
		dirList, err := ioutil.ReadDir(watchPath[i])
		if err != nil {
			log.Println(err)
			continue
		}

		// 监控子目录
		for _, v := range dirList {
			if v.IsDir() {
				dirName := watchPath[i] + v.Name() + string(os.PathSeparator)
				dirNames = append(dirNames, dirName)
				for _,path := range GetWatcherPaths([]string{dirName}){
					dirNames = append(dirNames, path)
				}
			}
		}
	}
	return dirNames
}

