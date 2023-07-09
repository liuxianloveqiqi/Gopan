package service

import (
	"Gopan/app/transfer/rmq/internal/config"
	"Gopan/app/upload/model"
	"Gopan/common/init_db"
	"encoding/json"
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
		MysqlDb:  mysqlDb,
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
		message, ok := <-ch
		if !ok {
			log.Fatal("seckill rmq exit")
		}
		m := *message
		fmt.Printf("consume msg: %+v\n", m)
		file := model.File{
			FileSha1:   m.FileSha1,
			FileName:   m.FileName,
			FileSize:   m.FileSize,
			FileAddr:   m.FileAddr,
			Status:     m.Status,
			CreateTime: m.CreateTime,
			UpdateTime: m.UpdateTime,
		}
		userfile := model.UserFile{
			UserId:     m.UserId,
			FileSha1:   m.FileSha1,
			FileName:   m.FileName,
			FileSize:   m.FileSize,
			Status:     m.Status,
			CreateTime: m.CreateTime,
			UpdateTime: m.UpdateTime,
		}

		if err := s.MysqlDb.Create(&file).Error; err != nil {
			logx.Error(err)
		}
		if err := s.MysqlDb.Create(&userfile).Error; err != nil {
			logx.Error(err)

		}
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
