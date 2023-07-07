package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FileModel = (*customFileModel)(nil)

type (
	// FileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFileModel.
	FileModel interface {
		fileModel
	}

	customFileModel struct {
		*defaultFileModel
	}
)

// NewFileModel returns a model for the database table.
func NewFileModel(conn sqlx.SqlConn, c cache.CacheConf) FileModel {
	return &customFileModel{
		defaultFileModel: newFileModel(conn, c),
	}
}

// 新增/更新文件元信息到mysql中
func (m *defaultFileModel) UpdateFileMeta(filehash string, filename string,
	filesize int64, fileaddr string) {

}
