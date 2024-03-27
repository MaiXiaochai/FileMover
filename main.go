package main

import (
	"file_mover/pkg/mover"
	"file_mover/pkg/utils"
	"log"
	"os"
	"time"
)

// 定义条件函数的类型
type conditionFunc func(file os.FileInfo) bool

// 实现一个具体的条件函数
var isTimeSafe conditionFunc = func(file os.FileInfo) bool {
	// 如果文件的修改日期比当前时间大两小时
	if file.IsDir() {
		return false
	}

	now := time.Now()
	twoHoursLater := now.Add(2 * time.Hour).Unix()
	return file.ModTime().Unix() <= twoHoursLater
}

func main() {

	cfg, err := utils.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	err = mover.MoveFiles(cfg.Path.SrcDir, cfg.Path.DestDir, isTimeSafe)
	if err != nil {
		log.Printf("Failed to move files: %v", err)
	}
	log.Println("Files moved successfully.")
}
