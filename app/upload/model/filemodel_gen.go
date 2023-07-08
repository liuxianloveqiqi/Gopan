// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	fileFieldNames          = builder.RawFieldNames(&File{})
	fileRows                = strings.Join(fileFieldNames, ",")
	fileRowsExpectAutoSet   = strings.Join(stringx.Remove(fileFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	fileRowsWithPlaceHolder = strings.Join(stringx.Remove(fileFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheFileIdPrefix       = "cache:file:id:"
	cacheFileFileSha1Prefix = "cache:file:fileSha1:"
)

type (
	fileModel interface {
		Insert(ctx context.Context, data *File) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*File, error)
		FindOneByFileSha1(ctx context.Context, fileSha1 string) (*File, error)
		Update(ctx context.Context, data *File) error
		Delete(ctx context.Context, id int64) error
	}

	defaultFileModel struct {
		sqlc.CachedConn
		table string
	}

	File struct {
		Id         int64        `db:"id"`
		FileSha1   string       `db:"file_sha1"` // 文件hash
		FileName   string       `db:"file_name"` // 文件名
		FileSize   int64        `db:"file_size"` // 文件大小
		FileAddr   string       `db:"file_addr"` // 文件存储位置
		Status     int64        `db:"status"`    // 状态(可用/禁用/已删除等状态)
		CreateTime time.Time    `db:"create_time"`
		UpdateTime time.Time    `db:"update_time"`
		DeleteTime sql.NullTime `db:"delete_time"`
	}
)

func newFileModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultFileModel {
	return &defaultFileModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`file`",
	}
}

func (m *defaultFileModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	fileFileSha1Key := fmt.Sprintf("%s%v", cacheFileFileSha1Prefix, data.FileSha1)
	fileIdKey := fmt.Sprintf("%s%v", cacheFileIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, fileFileSha1Key, fileIdKey)
	return err
}

func (m *defaultFileModel) FindOne(ctx context.Context, id int64) (*File, error) {
	fileIdKey := fmt.Sprintf("%s%v", cacheFileIdPrefix, id)
	var resp File
	err := m.QueryRowCtx(ctx, &resp, fileIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", fileRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFileModel) FindOneByFileSha1(ctx context.Context, fileSha1 string) (*File, error) {
	fileFileSha1Key := fmt.Sprintf("%s%v", cacheFileFileSha1Prefix, fileSha1)
	var resp File
	err := m.QueryRowIndexCtx(ctx, &resp, fileFileSha1Key, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `file_sha1` = ? limit 1", fileRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, fileSha1); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFileModel) Insert(ctx context.Context, data *File) (sql.Result, error) {
	fileFileSha1Key := fmt.Sprintf("%s%v", cacheFileFileSha1Prefix, data.FileSha1)
	fileIdKey := fmt.Sprintf("%s%v", cacheFileIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, fileRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.FileSha1, data.FileName, data.FileSize, data.FileAddr, data.Status, data.DeleteTime)
	}, fileFileSha1Key, fileIdKey)
	return ret, err
}

func (m *defaultFileModel) Update(ctx context.Context, newData *File) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	fileFileSha1Key := fmt.Sprintf("%s%v", cacheFileFileSha1Prefix, data.FileSha1)
	fileIdKey := fmt.Sprintf("%s%v", cacheFileIdPrefix, data.Id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, fileRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.FileSha1, newData.FileName, newData.FileSize, newData.FileAddr, newData.Status, newData.DeleteTime, newData.Id)
	}, fileFileSha1Key, fileIdKey)
	return err
}

func (m *defaultFileModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheFileIdPrefix, primary)
}

func (m *defaultFileModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", fileRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultFileModel) tableName() string {
	return m.table
}