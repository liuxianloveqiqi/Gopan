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
	chanCount   = 10   // 通道数量
	bufferCount = 1024 // 每个通道的缓冲区大小
)

type Service struct {
	c        config.Config // 配置信息
	MysqlDb  *gorm.DB      // MySQL 数据库连接对象
	Log      logx.LogConf
	waiter   sync.WaitGroup            // 用于等待所有消费者 goroutine 完成的等待组
	msgsChan []chan *model.NewUserFile // 消息通道切片，每个元素是一个通道，用于存放消息
}

// NewService 创建一个新的 Service 实例
func NewService(c config.Config) *Service {
	// 初始化 MySQL 数据库连接
	mysqlDb := init_db.InitGorm(c.MysqlCluster.DataSource)

	// 创建 user_file 和 file 表
	mysqlDb.AutoMigrate(&model.NewUserFile{}, &model.File{})

	// 创建 Service 实例
	s := &Service{
		c:        c,
		msgsChan: make([]chan *model.NewUserFile, chanCount),
		MysqlDb:  mysqlDb,
	}

	// 创建 chanCount 个消费者 goroutine
	for i := 0; i < chanCount; i++ {
		ch := make(chan *model.NewUserFile, bufferCount)
		s.msgsChan[i] = ch
		s.waiter.Add(1)
		go s.consume(ch)
	}

	return s
}

// consume 是消费者 goroutine 的函数，负责处理从通道中接收的消息
func (s *Service) consume(ch chan *model.NewUserFile) {
	defer s.waiter.Done()

	for {
		message, ok := <-ch
		if !ok {
			log.Fatal("接受消息失败")
		}
		m := *message
		fmt.Printf("消费消息: %+v\n", m)

		// 创建 File 和 UserFile 对象，用于写入数据库
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

		// 写入 file 表
		if err := s.MysqlDb.Create(&file).Error; err != nil {
			logx.Error(err)
		}

		// 写入 userfile 表
		if err := s.MysqlDb.Create(&userfile).Error; err != nil {
			logx.Error(err)
		}
	}
}

// Consume 是消费者的方法，用于处理消息
func (s *Service) Consume(_ string, value string) error {
	logx.Infof("消费消息: %s\n", value)

	// 将 JSON 数据解析为 []*model.NewUserFile 对象
	var data []*model.NewUserFile
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return err
	}

	// 将解析后的消息根据 UserId 分发到不同的通道
	for _, d := range data {
		s.msgsChan[d.UserId%chanCount] <- d
	}

	return nil
}
