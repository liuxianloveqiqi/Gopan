package model

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		FindUserBy(db *gorm.DB, field string, value interface{}) ([]User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c),
	}
}

// 查找
func (m *defaultUserModel) FindUserBy(db *gorm.DB, field string, value interface{}) ([]User, error) {
	var users []User
	if res := db.Where(field+" = ?", value).Find(&users); res.Error != nil {
		return nil, res.Error
	}
	fmt.Println(users)
	if len(users) == 0 {
		return nil, nil
	}
	return users, nil
}
