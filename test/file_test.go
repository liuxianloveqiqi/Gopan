package test

import (
	"log"
	"os"
	"path"
	"strconv"
	"testing"
)

func TestFile(t *testing.T) {
	// 获得文件句柄，用于存储分块内容
	filepath := "/Users/liuxian/GoProjects/project/Gopan/data/file/" + "ccccc" + "/" + strconv.FormatInt(2, 10)
	err := os.MkdirAll(path.Dir(filepath), 0744)
	if err != nil {
		log.Print(err)
	}

	fd, err := os.Create(filepath)
	if err != nil {
		log.Print(err)

	}
	defer fd.Close()
}
