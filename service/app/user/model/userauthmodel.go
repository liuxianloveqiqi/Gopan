package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

var _ UserAuthModel = (*customUserAuthModel)(nil)

type (
	// UserAuthModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserAuthModel.
	UserAuthModel interface {
		userAuthModel
		FindUserAuthBy(db *gorm.DB, field string, value interface{}) ([]UserAuth, error)
	}

	customUserAuthModel struct {
		*defaultUserAuthModel
	}
)

// NewUserAuthModel returns a model for the database table.
func NewUserAuthModel(conn sqlx.SqlConn, c cache.CacheConf) UserAuthModel {
	return &customUserAuthModel{
		defaultUserAuthModel: newUserAuthModel(conn, c),
	}
}

// 查找
func (m *defaultUserAuthModel) FindUserAuthBy(db *gorm.DB, field string, value interface{}) ([]UserAuth, error) {
	var user_auths []UserAuth
	if res := db.Where(field+" = ?", value).Find(&user_auths); res.Error != nil {
		return nil, res.Error
	}
	if len(user_auths) == 0 {
		return nil, nil
	}
	return user_auths, nil
}
