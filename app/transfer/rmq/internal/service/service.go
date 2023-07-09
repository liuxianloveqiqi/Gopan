package service

import (
	"Gopan/app/transfer/rmq/internal/config"
	"Gopan/app/upload/model"
	"Gopan/common/init_db"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"log"
	"sync"
)

const (
	chanCount   = 10
	bufferCount = 1024
)

type Service struct {
	c       config.Config
	MysqlDb *gorm.DB

	waiter   sync.WaitGroup
	msgsChan []chan *model.NewUserFile
}

func NewService(c config.Config) *Service {
	mysqlDb := init_db.InitGorm(c.MysqlCluster.DataSource)
	mysqlDb.AutoMigrate(&model.NewUserFile{}, &model.File{})
	s := &Service{
		c:        c,
		msgsChan: make([]chan *model.NewUserFile, chanCount),
	}
	for i := 0; i < chanCount; i++ {
		ch := make(chan *model.NewUserFile, bufferCount)
		s.msgsChan[i] = ch
		s.waiter.Add(1)
		//go s.consume(ch)
		go s.consume(ch)
	}
	return s
}

func (s *Service) consume(ch chan *model.NewUserFile) {
	defer s.waiter.Done()

	for {
		m, ok := <-ch
		if !ok {
			log.Fatal("seckill rmq exit")
		}
		fmt.Printf("consume msg: %+v\n", m)
		var file model.File
		if err := s.MysqlDb.Model(&model.File{}).Where("file_sha1 = ?", m.FileSha1).First(&file).Error; err != nil {
			if err == nil {
				// 查到记录

			} else if errors.Is(err, gorm.ErrRecordNotFound) {
				// 发生其他错误
			}
		}

		file := model.File{
			FileSha1:   m.FileSha1,
			FileName:   m.FileName,
			FileSize:   m.FileSize,
			FileAddr:   m.FileAddr,
			Status:     m.Status,
			CreateTime: m.CreateTime,
			UpdateTime: m.UpdateTime,
			DeleteTime: m.DeleteTime,
		}

		if err := s.MysqlDb.Create(&m).Error; err != nil {
		}
		//if err := s.MysqlDb.Create(&userfile).Error; err != nil {
	}
}

func (s *Service) Consume(_ string, value string) error {
	logx.Infof("Consume value: %s\n", value)
	var data []*model.NewUserFile
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return err
	}
	for _, d := range data {
		s.msgsChan[d.UserId%chanCount] <- d
	}
	return nil
}
